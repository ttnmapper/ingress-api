package chirpstack

import (
	"bytes"
	chirpstack "github.com/brocaar/chirpstack-api/go/v3/as/integration"
	"github.com/go-chi/render"
	"github.com/golang/protobuf/jsonpb"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"ttnmapper-ingress-api/types"
)

/*
Chirpstack V3 (Event) Webhook
No authorization is done as it is assumed that a ChirpStack instance is a private network
*/
func (handlerContext *Context) PostChirpV3Event(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print("[" + i + "] " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var packetOut types.TtnMapperUplinkMessage
	packetOut.NetworkType = types.NS_CHIRP
	// TODO ChirpStack should also provide us some unique identifier, along with the NetID, then we can do UUID@NetID to provide as a unique networkid
	packetOut.NetworkAddress = r.RemoteAddr //Might not work if behind a load-balancer
	packetOut.NetworkId = packetOut.NetworkType + "://" + packetOut.NetworkAddress

	event := r.URL.Query().Get("event")

	if event == "" {
		response["success"] = false
		response["message"] = "event parameter not specified"
		log.Print("[" + i + "] event parameter not specified")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// We handle ONLY Uplink Events
	if event != "up" {
		response["success"] = false
		response["message"] = "Handler for event \"" + event + "\" is not implemented"
		log.Print("[" + i + "] Handler for event \"" + event + "\" is not implemented")
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	unmarshaler := &jsonpb.Unmarshaler{
		AllowUnknownFields: true, // we don't want to fail on unknown fields
	}

	log.Printf("%+v", string(body))
	var packetIn chirpstack.UplinkEvent
	if err := unmarshaler.Unmarshal(bytes.NewReader(body), &packetIn); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print("[" + i + "] " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	CopyChirpV3Fields(packetIn, &packetOut)

	if err := ParseChirpV3Payload(packetIn, &packetOut); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print("[" + i + "] " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If the experiment header is set, mark this packetOut as experiment
	experiment := r.Header.Get("TTNMAPPERORG-EXPERIMENT")
	packetOut.Experiment = experiment // Default header is empty
	// If the user header is set, add it to packetOut. For ChirpStack we do not authenticate users.
	user := r.Header.Get("TTNMAPPERORG-USER")
	packetOut.UserId = user // Default header is empty

	log.Print("["+i+"] Network: ", packetOut.NetworkType, "://", packetOut.NetworkAddress)
	log.Print("["+i+"] Device: ", packetOut.AppID, " - ", packetOut.DevID)

	// Push this new packet into the stack
	handlerContext.PublishChannel <- packetOut

	w.WriteHeader(http.StatusAccepted)
	response["success"] = true
	response["message"] = "New packet accepted into queue"
}
