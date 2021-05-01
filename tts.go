package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	_ "time"
	"ttnmapper-ingress-api/ttsV3/models"
	"ttnmapper-ingress-api/types"
)

func TtsRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/v3/uplink-message", PostTtsV3Uplink)
	router.Post("/v3/join-accept", PostTtsV3JoinAccept)
	router.Post("/v3/location-solved", PostTtsV3LocationSolved)
	router.Get("/v3", GetTtsV3)

	return router
}

/*
TTS V3 Webhook
Header TTNMAPPERORG-USER contains the email address of the user for identification of the source of the data
Header TTNMAPPERORG-EXPERIMENT indicates if mapping is done to an experiment, and the experiment name
*/
func PostTtsV3Uplink(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	// Read data
	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	email := r.Header.Get("TTNMAPPERORG-USER")
	log.Print("["+i+"] User: ", email)
	if err := validateEmail(email); err != nil {
		w.WriteHeader(http.StatusForbidden)
		response["success"] = false
		response["message"] = err.Error()
		log.Print("[" + i + "] " + err.Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print("[" + i + "] " + err.Error())
		return
	}

	var packetIn models.V3ApplicationUp
	if err := json.Unmarshal(body, &packetIn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print("[" + i + "] " + err.Error())
		return
	}

	if packetIn.UplinkMessage == nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "uplink_message not set"
		log.Print("[" + i + "] uplink_message not set")
		return
	}

	if packetIn.UplinkMessage.DecodedPayload == nil {
		response["success"] = false
		response["message"] = "payload_fields not set"
		log.Print("[" + i + "] payload_fields not set")
		//return // Do not return, as we can still use the metadata to update gateway last seen and contribute channels and signal stats
	}

	var packetOut types.TtnMapperUplinkMessage
	packetOut.NetworkType = types.NS_TTS_V3

	// Use X-TTS-DOMAIN header to determine tenant and cluster
	ttsDomain := r.Header.Get("X-TTS-DOMAIN")
	if ttsDomain == "" {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Originating network server header not set"
		log.Print("[" + i + "] X-TTS-DOMAIN header not set")
		return
	}

	/*
		Determine the network
		X-Tts-Domain suffix .cloud.thethings.network is The Things Stack Community Edition (The Things Network)
		X-Tts-Domain suffix .cloud.thethings.industries is The Things Stack Cloud (The Things Industries)
		home_network_net_id 000013 and home_network_tenant_id ttn is The Things Network
		home_network_net_id 000013 with any other home_network_tenant_id can be anything: The Things Stack Cloud, The Things Stack Enterprise, The Things Stack Open Source, even ChirpStack using Passive Roaming and using address space of TTN

		Packet Broker will combine tenant ID and cluster ID in the NSID (tenant-id@cluster-id) when it gets LoRaWAN Backend Interfaces 1.1 support.
		TODO: Follow what happens on https://github.com/TheThingsNetwork/lorawan-stack/issues/4076
	*/
	if strings.HasSuffix(ttsDomain, ".cloud.thethings.network") {
		ttsDomain = "ttn@000013"
	}
	if strings.HasSuffix(ttsDomain, ".cloud.thethings.industries") {
		tenantId := strings.Split(ttsDomain, ".")[0]
		ttsDomain = tenantId + "@000013"
	}

	packetOut.NetworkAddress = ttsDomain                                           // ttn@000013
	packetOut.NetworkId = packetOut.NetworkType + "://" + packetOut.NetworkAddress // NS_TTS_V3://ttn@000013

	// TODO Try to use the location from the metadata first. This is likely the location set on the console.
	// TODO Need to update the ttsV3 models to include this data
	//if packetIn.UplinkMessage. != 0 && packetIn.Metadata.Longitude != 0 {
	//	packetOut.AccuracySource = packetIn.Metadata.Source
	//	packetOut.Latitude = float64(packetIn.Metadata.Latitude)
	//	packetOut.Longitude = float64(packetIn.Metadata.Longitude)
	//	packetOut.Altitude = float64(packetIn.Metadata.Altitude)
	//}

	// If payload fields are available, try getting coordinates from there
	if packetIn.UplinkMessage.DecodedPayload != nil {
		if err := ParsePayloadFields(packetIn.UplinkMessage.FPort, packetIn.UplinkMessage.DecodedPayload, &packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print("[" + i + "] " + err.Error())
			//return // also accept invalid payload_fields, as the metadata is still useful
		}
	}

	packetOut.UserAgent = r.Header.Get("USER-AGENT")
	packetOut.UserId = email

	// For V3 assume the experiment is passed via header so that we do not need a custom model
	experiment := r.Header.Get("TTNMAPPERORG-EXPERIMENT")
	packetOut.Experiment = experiment // Default header is empty

	// TODO move the sanity check to where we insert the data into the db, as invalid data is still used to update gateway last seen
	if packetOut.Experiment == "" {
		if err := CheckData(packetOut); err != nil {
			//response["success"] = false
			//response["message"] = err.Error()
			log.Print("["+i+"] Data invalid: ", err.Error())
			//return
		}

		SanitizeData(&packetOut)
	}

	// Add metadata fields
	CopyTtnV3Fields(packetIn, &packetOut)

	log.Print("["+i+"] Network: ", packetOut.NetworkType, "://", packetOut.NetworkAddress)
	log.Print("["+i+"] Device: ", packetOut.AppID, " - ", packetOut.DevID)

	// Push this new packet into the stack
	//publishChannel <- packetOut
	log.Println("[" + i + "] " + prettyPrint(packetOut))

	// TODO check if habhub header is set and true
	// TODO check if aprs header is set and true

	w.WriteHeader(http.StatusAccepted)
	response["success"] = true
	response["message"] = "New packet accepted into queue"
}

func GetTtsV3(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["success"] = true
	response["message"] = "GET test success"
	render.JSON(w, r, response)
}

func PostTtsV3JoinAccept(w http.ResponseWriter, r *http.Request) {

	// Read data
	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print(err.Error())
		return
	}

	log.Println("Join Accept: ", string(body))
}

func PostTtsV3LocationSolved(w http.ResponseWriter, r *http.Request) {

	// Read data
	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print(err.Error())
		return
	}

	log.Println("Location Solved: ", string(body))

}
