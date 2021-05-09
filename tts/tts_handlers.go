package tts

import (
	"bytes"
	"github.com/go-chi/render"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"
)

/*
TTS V3 Webhook
Header TTNMAPPERORG-USER contains the email address of the user for identification of the source of the data
Header TTNMAPPERORG-EXPERIMENT indicates if mapping is done to an experiment, and the experiment name
*/
func (handlerContext *Context) PostV3Uplink(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	// Read data
	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	email := r.Header.Get("TTNMAPPERORG-USER")
	log.Print("["+i+"] User: ", email)
	if err := utils.ValidateEmail(email); err != nil {
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
	var packetIn ttnpb.ApplicationUp

	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true, // we don't want to fail on unknown fields
		}
		if err := unmarshaler.Unmarshal(bytes.NewReader(body), &packetIn); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response["success"] = false
			response["message"] = "Can not parse json body"
			log.Print("[" + i + "] " + err.Error())
			return
		}
	} else if contentType == "application/protobuf" || contentType == "application/x-protobuf" || contentType == "application/octet-stream" { // TTS uses application/octet-stream
		if err := proto.Unmarshal(body, &packetIn); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response["success"] = false
			response["message"] = "Can not parse protobuf body"
			log.Print("[" + i + "] " + err.Error())
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Content-Type header not set"
		return
	}

	if packetIn.GetUplinkMessage() == nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "uplink_message not set"
		log.Print("[" + i + "] uplink_message not set")
		return
	}

	if packetIn.GetUplinkMessage().DecodedPayload == nil {
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

	// 1. Try to use the location from the metadata first. This is likely the location set on the console.
	if packetIn.GetUplinkMessage().Locations["user"] != nil {
		packetOut.Latitude = packetIn.GetUplinkMessage().Locations["user"].Latitude
		packetOut.Longitude = packetIn.GetUplinkMessage().Locations["user"].Longitude
		packetOut.Altitude = float64(packetIn.GetUplinkMessage().Locations["user"].Altitude)
		packetOut.AccuracyMeters = float64(packetIn.GetUplinkMessage().Locations["user"].Accuracy)
		packetOut.AccuracySource = packetIn.GetUplinkMessage().Locations["user"].Source.String()
	}

	// 2. If the packetIn contains a solved location, rather use that - TODO this is likely sent to the /location-solved endpoint
	if packetIn.GetLocationSolved() != nil {
		packetOut.Latitude = packetIn.GetLocationSolved().Latitude
		packetOut.Longitude = packetIn.GetLocationSolved().Longitude
		packetOut.Altitude = float64(packetIn.GetLocationSolved().Altitude)
		packetOut.AccuracyMeters = float64(packetIn.GetLocationSolved().Accuracy)
		packetOut.AccuracySource = packetIn.GetLocationSolved().Source.String()
	}

	// 3. If payload fields are available, try getting coordinates from there
	if packetIn.GetUplinkMessage().DecodedPayload != nil {
		// DecodedPayload is &Struct{Fields:map[string]*Value{},XXX_unrecognized:[],}.
		// Convert to a more standard map[string]interface{}
		parsedPayload := make(map[string]interface{})
		for k, v := range packetIn.GetUplinkMessage().DecodedPayload.Fields {
			parsedPayload[k] = v
		}

		if err := utils.ParsePayloadFields(int64(packetIn.GetUplinkMessage().FPort), parsedPayload, &packetOut); err != nil {
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
		if err := utils.CheckData(packetOut); err != nil {
			//response["success"] = false
			//response["message"] = err.Error()
			log.Print("["+i+"] Data invalid: ", err.Error())
			//return
		}

		utils.SanitizeData(&packetOut)
	}

	// Add metadata fields
	CopyV3Fields(packetIn, &packetOut)

	log.Print("["+i+"] Network: ", packetOut.NetworkType, "://", packetOut.NetworkAddress)
	log.Print("["+i+"] Device: ", packetOut.AppID, " - ", packetOut.DevID)

	// Push this new packet into the stack
	handlerContext.PublishChannel <- packetOut

	w.WriteHeader(http.StatusAccepted)
	response["success"] = true
	response["message"] = "New packet accepted into queue"
}

func (handlerContext *Context) GetV3(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["success"] = true
	response["message"] = "GET test success"
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}

func (handlerContext *Context) PostV3JoinAccept(w http.ResponseWriter, r *http.Request) {

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

	// TODO implement this endpoint
	w.WriteHeader(http.StatusNotImplemented)
}

func (handlerContext *Context) PostV3LocationSolved(w http.ResponseWriter, r *http.Request) {

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

	// TODO implement this endpoint
	w.WriteHeader(http.StatusNotImplemented)

}
