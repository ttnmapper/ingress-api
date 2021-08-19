package helium

import (
	"encoding/json"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"
)

func (handlerContext *Context) PostHelium(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	response := make(map[string]interface{})
	defer render.JSON(w, r, response) // TODO: this defer makes a copy of the response object, so it is always empty

	email := r.Header.Get("TTNMAPPERORG-USER")
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

	var packetIn UplinkMessage
	if err := json.Unmarshal(body, &packetIn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print("[" + i + "] " + err.Error())
		return
	}
	log.Println(utils.PrettyPrint(packetIn))

	var packetOut types.TtnMapperUplinkMessage

	packetOut.UserAgent = r.Header.Get("USER-AGENT")
	packetOut.UserId = email
	packetOut.Experiment = r.Header.Get("TTNMAPPERORG-EXPERIMENT")

	// 2. If payload fields are available, try getting coordinates from there
	if packetIn.Decoded.Payload != nil {
		if err := utils.ParsePayloadFields(int64(packetIn.Port), packetIn.Decoded.Payload, &packetOut); err != nil {
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
	CopyHeliumFields(packetIn, &packetOut)

	log.Print("["+i+"] Network: ", packetOut.NetworkType, "://", packetOut.NetworkAddress)
	log.Print("["+i+"] Device: ", packetOut.AppID, " - ", packetOut.DevID)

	handlerContext.PublishChannel <- packetOut

	w.WriteHeader(http.StatusAccepted)
	response["success"] = true
	response["message"] = "New packet accepted into queue"
}

func (handlerContext *Context) GetHelium(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["message"] = "GET test success"
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}
