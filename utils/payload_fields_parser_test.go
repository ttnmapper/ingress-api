package utils

import (
	"log"
	"testing"
	"ttnmapper-ingress-api/types"
)

func TestParsePayloadFields(t *testing.T) {

	payloadFields := make(map[string]interface{})
	payloadFields["latitude"] = 1.0
	payloadFields["longitude"] = 2.0
	payloadFields["altitude"] = 0.0

	packetOut := types.TtnMapperUplinkMessage{}

	_ = ParsePayloadFields(1, payloadFields, &packetOut)

	SanitizeData(&packetOut)

	log.Println(PrettyPrint(packetOut))
}
