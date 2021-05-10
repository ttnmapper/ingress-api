package ttn

import (
	"encoding/json"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"ttnmapper-ingress-api/ttn/ttn_types"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"
)

/*
TTN V2 HTTP integration
Authorization header contains email address
There is an extra Experiment field added to the model
*/
func (handlerContext *Context) PostTtnV2(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	response := make(map[string]interface{})
	defer render.JSON(w, r, response) // TODO: this defer makes a copy of the response object, so it is always empty

	email := r.Header.Get("Authorization")
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

	var packetIn ttn_types.UplinkMessage
	if err := json.Unmarshal(body, &packetIn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print("[" + i + "] " + err.Error())
		return
	}

	var packetOut types.TtnMapperUplinkMessage

	// For ttnv2 we use the ip address of the originating stack to id the network
	packetOut.NetworkType = types.NS_TTN_V2
	packetOut.NetworkAddress = r.RemoteAddr
	packetOut.NetworkId = packetOut.NetworkType + "://" + packetOut.NetworkAddress

	packetOut.UserAgent = r.Header.Get("USER-AGENT")
	packetOut.UserId = email
	packetOut.Experiment = packetIn.Experiment

	// 1. Try to use the location from the metadata first. This is likely the location set on the console.
	if packetIn.Metadata.Latitude != 0 && packetIn.Metadata.Longitude != 0 {
		packetOut.AccuracySource = packetIn.Metadata.Source
		packetOut.Latitude = float64(packetIn.Metadata.Latitude)
		packetOut.Longitude = float64(packetIn.Metadata.Longitude)
		packetOut.Altitude = float64(packetIn.Metadata.Altitude)
	}

	// 2. If payload fields are available, try getting coordinates from there
	if packetIn.PayloadFields != nil {
		if err := utils.ParsePayloadFields(int64(packetIn.FPort), packetIn.PayloadFields, &packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print("[" + i + "] " + err.Error())
			//return // also accept invalid payload_fields, as the metadata is still useful
		}
	}

	// Ignore data validity checks for experiments
	if packetOut.Experiment == "" {
		if err := utils.CheckData(packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print("["+i+"] Data invalid: ", err.Error())
			return
		}
		utils.SanitizeData(&packetOut)
	}

	// Copy metadata fields
	CopyTtnV2Fields(packetIn, &packetOut)

	log.Print("["+i+"] Network: ", packetOut.NetworkType, "://", packetOut.NetworkAddress)
	log.Print("["+i+"] Device: ", packetOut.AppID, " - ", packetOut.DevID)

	handlerContext.PublishChannel <- packetOut

	w.WriteHeader(http.StatusAccepted)
	response["success"] = true
	response["message"] = "New packet accepted into queue"
}

func (handlerContext *Context) GetTtnV2(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["message"] = "GET test success"
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}
