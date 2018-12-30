package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"net/http"
	"ttnmapper-ingress-api/types"
)

func AndroidRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/v2", PostAndroidV2)
	router.Post("/v3", PostAndroidV3)

	return router
}

func PostAndroidV2(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	response["success"] = false
	response["message"] = "Not implemented"

}

func PostAndroidV3(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//panic(err)
		response["success"] = false
		response["message"] = "Can not read POST body"
		return
	}
	log.Println(string(body))

	var receivedPacket types.TtnMapperAndroidMessage
	if err := json.Unmarshal(body, &receivedPacket); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		return
	}

	var resultPacket = types.TtnMapperUplinkMessage{}
	CopyAndroidToTtnMapper(receivedPacket, &resultPacket)

	if err := CheckData(resultPacket); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		return
	}

	SanitizeData(&resultPacket)

	publish_channel <- resultPacket

	response["success"] = true
	response["message"] = "New packet accepted into queue"
	response["packet"] = resultPacket
}

func CopyAndroidToTtnMapper(source types.TtnMapperAndroidMessage, destination *types.TtnMapperUplinkMessage) {
	destination.TtnMLatitude = source.PhoneLat
	destination.TtnMLongitude = source.PhoneLon
	destination.TtnMAltitude = source.PhoneAlt
	destination.TtnMAccuracy = source.PhoneLocAccuracy
	destination.TtnMProvider = source.PhoneLocprovider
	destination.TtnMUserId = source.Iid
	destination.TtnMUserAgent = source.UserAgent
	destination.TtnMExperiment = source.Experiment
}
