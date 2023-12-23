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
		err = handlerContext.up(r, body)
	case "join":
		err = handlerContext.join(r, body)
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
}

func (handlerContext *Context) up(r *http.Request, body []byte) error {
	var up integration.UplinkEvent
	if err := handlerContext.unmarshal(r, body, &up); err != nil {
		return err
	}
	log.Printf("Uplink received from %s with payload: %s", up.GetDeviceInfo().DevEui, hex.EncodeToString(up.Data))
	for _, gateway := range up.RxInfo {
		log.Print(utils.PrettyPrint(gateway))
	}
	return nil
}

func (handlerContext *Context) join(r *http.Request, body []byte) error {
	var join integration.JoinEvent
	if err := handlerContext.unmarshal(r, body, &join); err != nil {
		return err
	}
	log.Printf("Device %s joined with DevAddr %s", join.GetDeviceInfo().DevEui, join.DevAddr)
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
