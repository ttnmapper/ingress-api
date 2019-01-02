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

func TtnRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/v2", PostTtnV2)
	router.Get("/v2", GetTtnV2)
	router.Post("/v3", PostTtnV3)

	return router
}

func PostTtnV2(w http.ResponseWriter, r *http.Request) {

	/*
		TODO
		Authorization header should contain a valid email address
	*/

	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	email := r.Header.Get("Authorization")
	log.Print(email)
	if err := validateEmail(email); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//panic(err)
		response["success"] = false
		response["message"] = "Can not read POST body"
		return
	}
	//log.Println(string(body))

	//packet := make(map[string]interface{})
	var packet types.TtnMapperUplinkMessage
	if err := json.Unmarshal(body, &packet); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		return
	}

	if err := ParsePayloadFields(&packet); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		return
	}

	if err := CheckData(packet); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		return
	}

	SanitizeData(&packet)

	packet.TtnMUserAgent = "ttn-v2-integration"
	//TODO packet.TtnMUserId = //email address

	publish_channel <- packet

	response["success"] = true
	response["message"] = "New packet accepted into queue"
	response["packet"] = packet

}

func GetTtnV2(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["message"] = "GET test success"
	render.JSON(w, r, response)
}

func PostTtnV3(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["success"] = false
	response["message"] = "Not implemented"
	render.JSON(w, r, response)
}

func GetTtnV3(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["success"] = true
	response["message"] = "GET test success"
	render.JSON(w, r, response)
}
