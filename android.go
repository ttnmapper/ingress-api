package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ulule/deepcopier"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"ttnmapper-ingress-api/ttnV2"
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

	if resultPacket.Experiment == "" {
		if err := CheckData(resultPacket); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print(err.Error())
			return
		}

		SanitizeData(&resultPacket)
	}

	publishChannel <- resultPacket

	response["success"] = true
	response["message"] = "New packet accepted into queue"
	//response["packet"] = resultPacket

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
	//response["packet"] = resultPacket
}

func CopyAndroidV3ToTtnMapper(source types.TtnMapperAndroidMessage, destination *types.TtnMapperUplinkMessage) {

	destination.Latitude = source.PhoneLat
	destination.Longitude = source.PhoneLon
	destination.Altitude = source.PhoneAlt
	destination.AccuracyMeters = source.PhoneLocAccuracy
	destination.AccuracySource = source.PhoneLocProvider
	destination.UserId = source.Iid
	destination.UserAgent = source.UserAgent
	destination.Experiment = source.Experiment

	packetIn := ttnV2.UplinkMessage{}

	// Copy the matching fields in TtnMapperAndroidMessage to ttnV2.UplinkMessage
	deepcopier.Copy(source).To(&packetIn)
	// Copy and format the ttnV2.UplinkMessage fields we are interested in into destination
	CopyTtnV2Fields(packetIn, destination)
}

func CopyAndroidV2ToTtnMapper(source types.TtnMapperAndroidV2Message, destination *types.TtnMapperUplinkMessage) {

	destination.Latitude = source.Latitude
	destination.Longitude = source.Longitude
	destination.AccuracyMeters = source.Accuracy
	destination.Altitude = source.Altitude

	destination.UserId = source.Iid
	destination.AccuracySource = source.Provider
	destination.UserAgent = source.UserAgent
	destination.Experiment = source.Experiment

	destination.AppID = source.AppId
	destination.DevID = source.DevId

	destination.Time = time.Time(source.Time).UnixNano()

	// Assume LORA
	destination.Modulation = "LORA"
	drParts := strings.Split(source.Datarate, "BW")
	bandwidth, _ := strconv.Atoi(drParts[1])
	destination.Bandwidth = uint64(bandwidth * 1000)
	sf, _ := strconv.Atoi(strings.TrimPrefix(drParts[0], "SF"))
	destination.SpreadingFactor = uint8(sf)

	destination.Frequency = uint64(source.Freq)

	// Append gateway
	gatewayOut := types.TtnMapperGateway{}
	gatewayOut.GatewayId = source.GwId

	// If the id is eui-deadbeef, strip the prefix, capitalize and use as EUI
	if strings.HasPrefix(source.GwId, "eui-") && len(source.GwId) == 20 {
		eui := strings.TrimPrefix(source.GwId, "eui-")
		strings.ToUpper(eui)
		gatewayOut.GatewayEui = eui
	}

	gatewayOut.Time = time.Time(source.Time).UnixNano()
	gatewayOut.Rssi = source.Rssi
	gatewayOut.Snr = source.Snr

	destination.Gateways = append(destination.Gateways, gatewayOut)
}
