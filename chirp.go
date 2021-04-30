package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	_ "time"
	"ttnmapper-ingress-api/types"

	chirpstack "github.com/brocaar/chirpstack-api/go/v3/as/integration"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/golang/protobuf/jsonpb"
)

func ChirpRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/v3/events", PostChirpV3Event)
	return router
}

/*
Chirpstack V3 (Event) Webhook
*/
func PostChirpV3Event(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print("[" + i + "] " + err.Error())
		return
	}

	var packetOut types.TtnMapperUplinkMessage
	packetOut.NetworkType = types.NS_CHIRP
	event := r.URL.Query().Get("event")

	if event != "up" { // We handle ONLY Uplink Events
		response["success"] = false
		response["message"], _ = fmt.Printf("Handler for event %s is not implemented", event)
		log.Print("[" + i + "] " + err.Error())
		return
	}

	unmarshaler := &jsonpb.Unmarshaler{
		AllowUnknownFields: true, // we don't want to fail on unknown fields
	}

	var packetIn chirpstack.UplinkEvent
	if err := unmarshaler.Unmarshal(bytes.NewReader(body), &packetIn); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print("[" + i + "] " + err.Error())
		return
	}

	CopyChirpV3Fields(packetIn, &packetOut)
	packetOut.NetworkAddress = r.RemoteAddr //Might not work if behind a load-balancer

	if err := ParseChirpV3Payload(packetIn, &packetOut); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print("[" + i + "] " + err.Error())
		return
	}

	log.Print("["+i+"] Network: ", packetOut.NetworkType, "://", packetOut.NetworkAddress)
	log.Print("["+i+"] Device: ", packetOut.AppID, " - ", packetOut.DevID)

	// Push this new packet into the stack
	publishChannel <- packetOut

	w.WriteHeader(http.StatusAccepted)
	response["success"] = true
	response["message"] = "New packet accepted into queue"
}
