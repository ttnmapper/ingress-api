package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"net/http"
	_ "time"
	"ttnmapper-ingress-api/ttnV2"
	"ttnmapper-ingress-api/types"
)

func TtnRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/v2", PostTtnV2)
	router.Get("/v2", GetTtnV2)

	return router
}

/*
TTN V2 HTTP integration
Authorization header contains email address
There is an extra Experiment field added to the model
*/
func PostTtnV2(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	email := r.Header.Get("Authorization")
	if err := validateEmail(email); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print(err.Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print(err.Error())
		return
	}

	var packetIn ttnV2.UplinkMessage
	if err := json.Unmarshal(body, &packetIn); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print(err.Error())
		return
	}

	log.Println(email, " ", packetIn.AppID, " ", packetIn.DevID)

	var packetOut types.TtnMapperUplinkMessage
	packetOut.NetworkType = types.NS_TTN_V2
	packetOut.NetworkAddress = r.RemoteAddr

	packetOut.UserAgent = "ttn-v2-integration"
	packetOut.UserId = email
	packetOut.Experiment = packetIn.Experiment

	// Try to use the location from the metadata first. This is likely the location set on the console.
	if packetIn.Metadata.Latitude != 0 && packetIn.Metadata.Longitude != 0 {
		packetOut.AccuracySource = packetIn.Metadata.Source
		packetOut.Latitude = float64(packetIn.Metadata.Latitude)
		packetOut.Longitude = float64(packetIn.Metadata.Longitude)
		packetOut.Altitude = float64(packetIn.Metadata.Altitude)
	}

	// If payload fields are available, try getting coordinates from there
	if packetIn.PayloadFields != nil {
		if err := ParsePayloadFields(int64(packetIn.FPort), packetIn.PayloadFields, &packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print(err.Error())
			return
		}
	}

	// Ignore data validity checks for experiments
	if packetOut.Experiment == "" {
		if err := CheckData(packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print(err.Error())
			return
		}
		SanitizeData(&packetOut)
	}

	// Copy metadata fields
	CopyTtnV2Fields(packetIn, &packetOut)

	publishChannel <- packetOut

	response["success"] = true
	response["message"] = "New packet accepted into queue"
}

func GetTtnV2(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["message"] = "GET test success"
	render.JSON(w, r, response)
}
