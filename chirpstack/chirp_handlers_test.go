package chirpstack

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"
)

func TestHandlerJson(t *testing.T) {
	var publishChannel = make(chan types.TtnMapperUplinkMessage, 1)
	context := &Context{PublishChannel: publishChannel}

	postbodies := []string{
		`{"applicationID":"42","applicationName":"Adeunis_Fieldtester","deviceName":"Adeunis-Fieldtester-1","devEUI":"ABiyAAAAImU=","rxInfo":[{"gatewayID":"AIAAAKAARjM=","time":"2021-04-20T12:41:05.733901733Z","timeSinceGPSEpoch":"1302957684.734172002s","rssi":-103,"loRaSNR":7.75,"channel":1,"rfChain":1,"board":0,"antenna":0,"location":{"latitude":52.445313166666665,"longitude":10.8140175,"altitude":100,"source":"UNKNOWN","accuracy":0},"fineTimestampType":"NONE","context":"y4EMIw==","uplinkID":"BFRd6U9+R8W7pHc0ufA3BQ==","crcStatus":"CRC_OK"},{"gatewayID":"JOEk//7xFn0=","time":"2021-04-20T12:41:06.748077Z","timeSinceGPSEpoch":null,"rssi":-64,"loRaSNR":13.2,"channel":1,"rfChain":1,"board":0,"antenna":0,"location":{"latitude":52.42808398810916,"longitude":10.7915997505188,"altitude":0,"source":"UNKNOWN","accuracy":0},"fineTimestampType":"NONE","context":"CiIbgg==","uplinkID":"kpjPB3V3RgufxvhuG+B1iQ==","crcStatus":"NO_CRC"},{"gatewayID":"/MI9//4LkjA=","time":null,"timeSinceGPSEpoch":null,"rssi":-60,"loRaSNR":10,"channel":1,"rfChain":1,"board":0,"antenna":0,"location":{"latitude":52.428698929952496,"longitude":10.79209327697754,"altitude":0,"source":"UNKNOWN","accuracy":0},"fineTimestampType":"NONE","context":"8cRDyw==","uplinkID":"axbEx7GDQX+ztG/mpRJlBQ==","crcStatus":"CRC_OK"},{"gatewayID":"AIAAAKAAQZo=","time":null,"timeSinceGPSEpoch":null,"rssi":-50,"loRaSNR":10.25,"channel":1,"rfChain":1,"board":0,"antenna":0,"location":{"latitude":52.4280970720675,"longitude":10.791621208190918,"altitude":0,"source":"UNKNOWN","accuracy":0},"fineTimestampType":"NONE","context":"xGRPWw==","uplinkID":"dnNpkaURR/OYSwE/GKB88g==","crcStatus":"CRC_OK"},{"gatewayID":"AIAAAKAARjI=","time":"2021-04-20T12:41:05.733894733Z","timeSinceGPSEpoch":"1302957684.733703858s","rssi":-107,"loRaSNR":8,"channel":1,"rfChain":1,"board":0,"antenna":0,"location":{"latitude":52.42793283333334,"longitude":10.791740833333334,"altitude":85,"source":"UNKNOWN","accuracy":0},"fineTimestampType":"NONE","context":"maePww==","uplinkID":"JLsuAsZoSFOZWELy2M42ng==","crcStatus":"CRC_OK"}],"txInfo":{"frequency":868300000,"modulation":"LORA","loRaModulationInfo":{"bandwidth":125,"spreadingFactor":7,"codeRate":"4/5","polarizationInversion":false}},"adr":true,"dr":5,"fCnt":5,"fPort":1,"data":"vyVSJWVQAQR0kBUEAw+JNgY=","objectJSON":"{\"altitude\":0,\"battery\":3977,\"downlink\":3,\"latitude\":52.42758333333333,\"longitude\":10.7915,\"rssi\":-54,\"sats\":5,\"snr\":6,\"temperature\":37,\"trigger\":\"pushbutton\",\"uplink\":4}","tags":{"productive":"false"},"confirmedUplink":false,"devAddr":"AVoyvw=="}`,
	}

	for _, postbody := range postbodies {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("POST", "?event=up", strings.NewReader(postbody))
		if err != nil {
			t.Fatal(err)
		}
		req.RemoteAddr = "localhost"

		// Set request headers
		req.Header.Set("Host", "private.chirpstack.org")
		req.Header.Set("User-Agent", "http-ttn/2.6.0")
		req.Header.Set("Ttnmapperorg-user", "test@jpmeijers.com")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Ttnmapperorg-Experiment", "test-experiment")
		req.Header.Set("Ttnmapperorg-Network", "test-network")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(context.PostChirpV3Event)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		log.Println(rr.Body.String())

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusAccepted {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusAccepted)
		}

		// Check the response body is what we expect.
		expected := `{"message":"New packet accepted into queue","success":true}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}

		// Check if a packet was written to the queue
		select {
		case packetOut, ok := <-publishChannel:
			if ok {
				log.Println(utils.PrettyPrint(packetOut))
			} else {
				t.Error("Channel closed!")
			}
		default:
			t.Error("No value ready, moving on.")
		}
	}
}
