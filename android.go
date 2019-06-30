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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//panic(err)
		response["success"] = false
		response["message"] = "Can not read POST body"
		return
	}
	log.Println(string(body))

	var receivedPacket = types.TtnMapperAndroidV2Message{}
	if err := json.Unmarshal(body, &receivedPacket); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print(err.Error())
		return
	}

	var resultPacket = types.TtnMapperUplinkMessage{}
	CopyAndroidV2ToTtnMapper(receivedPacket, &resultPacket)

	if err := CheckData(resultPacket); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print(err.Error())
		return
	}

	SanitizeData(&resultPacket)

	publishChannel <- resultPacket

	response["success"] = true
	response["message"] = "New packet accepted into queue"
	response["packet"] = resultPacket

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
		log.Print(err.Error())
		return
	}

	var resultPacket = types.TtnMapperUplinkMessage{}
	CopyAndroidV3ToTtnMapper(receivedPacket, &resultPacket)

	if err := CheckData(resultPacket); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print(err.Error())
		return
	}

	SanitizeData(&resultPacket)

	publishChannel <- resultPacket

	response["success"] = true
	response["message"] = "New packet accepted into queue"
	response["packet"] = resultPacket
}

func CopyAndroidV3ToTtnMapper(source types.TtnMapperAndroidMessage, destination *types.TtnMapperUplinkMessage) {
	destination.TtnMLatitude = source.PhoneLat
	destination.TtnMLongitude = source.PhoneLon
	destination.TtnMAltitude = source.PhoneAlt
	destination.TtnMAccuracy = source.PhoneLocAccuracy
	destination.TtnMProvider = source.PhoneLocProvider
	destination.TtnMUserId = source.Iid
	destination.TtnMUserAgent = source.UserAgent
	destination.TtnMExperiment = source.Experiment
}

func CopyAndroidV2ToTtnMapper(source types.TtnMapperAndroidV2Message, destination *types.TtnMapperUplinkMessage) {

	destination.TtnMLatitude = source.Latitude
	destination.TtnMLongitude = source.Longitude
	destination.TtnMAccuracy = source.Accuracy
	destination.TtnMAltitude = source.Altitude

	destination.TtnMUserId = source.Iid
	destination.TtnMProvider = source.Provider
	destination.TtnMUserAgent = source.UserAgent
	destination.TtnMExperiment = source.Experiment

	destination.AppID = source.AppId
	destination.DevID = source.DevId

	destination.Metadata.Time = source.Time
	destination.Metadata.DataRate = source.Datarate
	destination.Metadata.Frequency = source.Freq

	gateway := types.GatewayMetadata{}
	gateway.GtwID = source.GwId
	gateway.RSSI = source.Rssi
	gateway.SNR = source.Snr
	gateway.Time = source.Time

	var gateways []types.GatewayMetadata
	gateways = append(gateways, gateway)
	destination.Metadata.Gateways = gateways
}
