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
	"ttnmapper-ingress-api/ttnV3/models"
	"ttnmapper-ingress-api/types"
)

func TtnRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/v2", PostTtnV2)
	router.Get("/v2", GetTtnV2)
	router.Post("/v3", PostTtnV3)
	router.Get("/v3", GetTtnV3)

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

	// Try to use the location from the metadata first. This is likely the location set on the console.
	if packetIn.Metadata.Latitude != 0 && packetIn.Metadata.Longitude != 0 {
		packetOut.AccuracySource = packetIn.Metadata.Source
		packetOut.Latitude = float64(packetIn.Metadata.Latitude)
		packetOut.Longitude = float64(packetIn.Metadata.Longitude)
		packetOut.Altitude = float64(packetIn.Metadata.Altitude)
	}

	if packetIn.PayloadFields != nil {
		if err := ParsePayloadFields(int64(packetIn.FPort), packetIn.PayloadFields, &packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print(err.Error())
			return
		}
	}

	packetOut.UserAgent = "ttn-v2-integration"
	packetOut.UserId = email
	packetOut.Experiment = packetIn.Experiment

	// If we got the location from the metadata, store under an experiment
	if packetOut.AccuracySource == packetIn.Metadata.Source {
		packetOut.Experiment = "registry-location" // TODO remove after verified to work

		// Also check and sanitize registry data
		if err := CheckData(packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print(err.Error())
			return
		}

		SanitizeData(&packetOut)
	}

	if packetOut.Experiment == "" {
		if err := CheckData(packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print(err.Error())
			return
		}

		SanitizeData(&packetOut)
	}

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

/*
TTN and TTI V3 stacks Webhook
*/
func PostTtnV3(w http.ResponseWriter, r *http.Request) {

	// Read data
	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	email := r.Header.Get("Authorization")
	log.Print(email)
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

	var packetIn models.V3ApplicationUp
	if err := json.Unmarshal(body, &packetIn); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print(err.Error())
		return
	}

	if packetIn.UplinkMessage.DecodedPayload == nil {
		response["success"] = false
		response["message"] = "payload_fields not set"
		log.Print("payload_fields not set")
		return
	}

	// Parse payload fields and check validity
	var packetOut types.TtnMapperUplinkMessage
	if err := ParsePayloadFields(packetIn.UplinkMessage.FPort, packetIn.UplinkMessage.DecodedPayload, &packetOut); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print(err.Error())
		return
	}

	packetOut.UserAgent = "ttn-v3-integration"
	packetOut.UserId = email

	// For V3 assume the experiment is passed via header so that we do not need a custom model
	experiment := r.Header.Get("Experiment")
	if experiment != "" {
		packetOut.Experiment = experiment
	}

	if packetOut.Experiment == "" {
		if err := CheckData(packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print(err.Error())
			return
		}

		SanitizeData(&packetOut)
	}

	// Add metadata fields
	CopyTtnV3Fields(packetIn, &packetOut)

	publishChannel <- packetOut

	response["success"] = true
	response["message"] = "New packet accepted into queue"
}

func GetTtnV3(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["success"] = true
	response["message"] = "GET test success"
	render.JSON(w, r, response)
}
