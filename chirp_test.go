package main

import (
	"os"
	"testing"
	"ttnmapper-ingress-api/types"

	chirpstack "github.com/brocaar/chirpstack-api/go/v3/as/integration"
	"github.com/golang/protobuf/jsonpb"
)

func TestChirpV3(t *testing.T) {
	// Open our jsonFile
	body, err := os.Open("tests/chirp_v3_output.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		t.Error(err)
		return
	}

	defer body.Close()
	unmarshaler := &jsonpb.Unmarshaler{
		AllowUnknownFields: true, // we don't want to fail on unknown fields
	}

	var packetOut types.TtnMapperUplinkMessage
	packetOut.NetworkType = types.NS_CHIRP

	var packetIn chirpstack.UplinkEvent
	if err := unmarshaler.Unmarshal(body, &packetIn); err != nil {
		t.Error(err)
		return
	}
	// Gateway
	// "location": {
	//     "latitude": 52.445313166666665,
	//     "longitude": 10.8140175,
	//     "altitude": 100,
	//     "source": "UNKNOWN",
	//     "accuracy": 0
	//   }
	CopyChirpV3Fields(packetIn, &packetOut)
	gw := packetOut.Gateways[0]

	gatewayCheck := gw.Latitude == 52.445313166666665 && gw.Longitude == 10.8140175 &&
		gw.Altitude == 100 && gw.LocationSource == "UNKNOWN" && gw.LocationAccuracy == 0
	if !gatewayCheck {
		t.Error("Gateway was not parsed correctly")
		return
	}

	// {\"altitude\":0,\"battery\":3977,\"downlink\":3,\"latitude\":52.42758333333333,\"longitude\":10.7915,\"rssi\":-54,\"sats\":5,
	if err := ParseChirpV3Payload(packetIn, &packetOut); err != nil {
		t.Error(err)
		return
	}

	payloadCheck := packetOut.Altitude == 0 && packetOut.Latitude == 52.42758333333333 &&
		packetOut.Longitude == 10.7915 && packetOut.Satellites == 5
	if !payloadCheck {
		t.Error("Payload was not parsed correctly")
		return
	}

}
