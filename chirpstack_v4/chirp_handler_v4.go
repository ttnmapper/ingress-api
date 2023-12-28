package chirpstack_v4

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"

	"github.com/chirpstack/chirpstack/api/go/v4/integration"
)

func (handlerContext *Context) PostChirpV4Event(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print("[" + i + "] " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := r.URL.Query().Get("event")

	switch event {
	case "up":
		err = handlerContext.up(w, r, body)
	case "join":
		err = handlerContext.join(w, r, body)
	case "status":
		err = handlerContext.status(w, r, body)
	default:
		response["success"] = false
		response["message"] = fmt.Sprintf("handler for event %s is not implemented", event)
		log.Printf("[%s] handler for event %s is not implemented", i, event)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Printf("[%s] handling event '%s' returned error: %s", i, event, err.Error())
		response["success"] = false
		response["message"] = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	response["success"] = true
	response["message"] = "packet accepted"
}

func (handlerContext *Context) up(w http.ResponseWriter, r *http.Request, body []byte) error {
	var packetIn integration.UplinkEvent
	if err := handlerContext.unmarshal(r, body, &packetIn); err != nil {
		return err
	}

	log.Printf("Device %s with EUI %s - up payload %s", packetIn.DeviceInfo.DeviceName, packetIn.GetDeviceInfo().DevEui, hex.EncodeToString(packetIn.Data))

	var packetOut types.TtnMapperUplinkMessage

	// TODO ChirpStack should also provide us some unique identifier, along with the NetID, then we can do UUID@NetID to provide as a unique networkid
	networkAddress := r.Header.Get("TTNMAPPERORG-NETWORK")
	if err := utils.ValidateChirpNetworkAddress(networkAddress); err != nil {
		return errors.New("header TTNMAPPERORG-NETWORK is empty")
	}
	packetOut.NetworkId = types.NS_CHIRP + "://" + networkAddress // NS_CHIRP://my.network.name

	if err := CopyChirpV4Fields(packetIn, &packetOut); err != nil {
		return err
	}

	if err := ParseChirpV4Payload(packetIn, &packetOut); err != nil {
		return err
	}

	// If the experiment header is set, mark this packetOut as experiment
	experiment := r.Header.Get("TTNMAPPERORG-EXPERIMENT")
	packetOut.Experiment = experiment // Default header is empty
	// If the user header is set, add it to packetOut. For ChirpStack we do not authenticate users.
	user := r.Header.Get("TTNMAPPERORG-USER")
	packetOut.UserId = user // Default header is empty

	log.Print("Network: ", packetOut.NetworkId)
	log.Print("Device: ", packetOut.AppID, " - ", packetOut.DevID)

	// Push this new packet into the stack
	handlerContext.PublishChannel <- packetOut
	//log.Println(utils.PrettyPrint(packetOut))

	return nil
}

func (handlerContext *Context) join(w http.ResponseWriter, r *http.Request, body []byte) error {
	var join integration.JoinEvent
	if err := handlerContext.unmarshal(r, body, &join); err != nil {
		return err
	}
	log.Printf("Device %s with EUI %s - join", join.DeviceInfo.DeviceName, join.GetDeviceInfo().DevEui)
	// Not using this yet
	return nil
}

func (handlerContext *Context) status(w http.ResponseWriter, r *http.Request, body []byte) error {
	var status integration.StatusEvent
	if err := handlerContext.unmarshal(r, body, &status); err != nil {
		return err
	}
	log.Printf("Device %s with EUI %s - status", status.DeviceInfo.DeviceName, status.GetDeviceInfo().DevEui)
	// Nothing we can do with this
	return nil
}

func (handlerContext *Context) unmarshal(r *http.Request, body []byte, v proto.Message) error {
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		// JSON
		return protojson.UnmarshalOptions{
			DiscardUnknown: true,
			AllowPartial:   true,
		}.Unmarshal(body, v)
	} else if contentType == "application/protobuf" || contentType == "application/x-protobuf" || contentType == "application/octet-stream" { // Chirp uses application/octet-stream
		// Protobuf
		return proto.Unmarshal(body, v)
	} else {
		return errors.New("Content-Type header not set or not supported")
	}

}
