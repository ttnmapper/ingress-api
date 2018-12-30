package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"net/http"
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

	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//panic(err)
		response["success"] = false
		response["message"] = "Can not parse json body"
		return
	}

	response["success"] = true
	response["message"] = "Created entry"

}

func PostAndroidV3(w http.ResponseWriter, r *http.Request) {

}
