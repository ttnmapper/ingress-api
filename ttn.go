package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"net/http"
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
		Fix SCG frequency

	*/

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

	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		return
	}

	log.Println(data)

	num, success := data["num"].(float64)
	if success != true {
		log.Printf("num doesn't exist")
	} else {
		log.Println(num)
	}

	response["success"] = true
	response["message"] = "Created entry"

}

func GetTtnV2(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["message"] = "GET test success"
	publish_channel <- response
	render.JSON(w, r, response)
}

func PostTtnV3(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Not implemented"
	render.JSON(w, r, response)
}

func GetTtnV3(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "GET test success"
	render.JSON(w, r, response)
}
