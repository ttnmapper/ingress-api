package ttn

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
		`{"app_id":"jpm_crickets","dev_id":"cricket_001","hardware_serial":"00E150E95369A0B0","port":1,"counter":29932,"payload_raw":"AIj60loC4SUALuY=","payload_fields":{"gps_0":{"altitude":120.06,"latitude":-33.9366,"longitude":18.8709}},"metadata":{"time":"2020-04-05T13:26:44.519096325Z","frequency":868.1,"modulation":"LORA","data_rate":"SF7BW125","coding_rate":"4/5","gateways":[{"gtw_id":"eui-7276ff00080e0176","timestamp":3590545203,"time":"2020-04-05T13:26:44Z","channel":0,"rssi":-64,"snr":10,"rf_chain":0},{"gtw_id":"eui-b827ebfffed88375","gtw_trusted":true,"timestamp":510723987,"time":"2020-04-05T13:26:44Z","channel":0,"rssi":-60,"snr":9.75,"rf_chain":1,"latitude":-33.93597,"longitude":18.870806},{"gtw_id":"eui-647fdafffe007a1a","timestamp":587147324,"time":"","channel":5,"rssi":-54,"snr":10.2,"rf_chain":0},{"gtw_id":"eui-60c5a8fffe71a964","timestamp":2492163779,"time":"","channel":0,"rssi":-43,"snr":9,"rf_chain":0}]},"downlink_url":"https://integrations.thethingsnetwork.org/ttn-eu/api/v2/down/jpm_crickets/test-post-to-ttnmapper?key=ttn-account-v2.G7VmPspZFw8TGRJFk7bcSFM4hBCJrm00P-JQuOcG1TQ"}`,
		`{"app_id":"homebug-qalcosonic-w1","dev_id":"05332403","hardware_serial":"0007090000515DB3","port":100,"counter":30,"payload_raw":"CsajXxABAAAAIOyiXwEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=","metadata":{"time":"2020-11-05T09:31:50.742526241Z","frequency":868.3,"modulation":"LORA","data_rate":"SF7BW125","airtime":112896000,"coding_rate":"4/5","gateways":[{"gtw_id":"eui-647fdafffe007a1a","timestamp":2079146396,"time":"","channel":6,"rssi":-47,"snr":8,"rf_chain":0},{"gtw_id":"eui-60c5a8fffe761551","timestamp":28410708,"time":"","channel":6,"rssi":-68,"snr":9.8,"rf_chain":0},{"gtw_id":"eui-58a0cbfffe80049a","timestamp":1526860595,"time":"2020-11-05T09:31:50.487390995Z","channel":0,"rssi":-61,"snr":10,"rf_chain":0},{"gtw_id":"eui-647fdafffe007a1f","timestamp":2815033635,"time":"","channel":1,"rssi":-85,"snr":10,"rf_chain":0},{"gtw_id":"eui-60c5a8fffe71a964","timestamp":48949428,"time":"","channel":6,"rssi":-87,"snr":10.3,"rf_chain":0,"latitude":34,"longitude":19,"altitude":1000}],"latitude":-33.93664,"longitude":18.870983,"location_source":"registry"}}`,
	}

	for _, postbody := range postbodies {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("POST", "", strings.NewReader(postbody))
		if err != nil {
			t.Fatal(err)
		}

		// Set request headers
		req.Header.Set("Host", "private.ttnmapper.org")
		req.Header.Set("User-Agent", "http-ttn/2.6.0")
		req.Header.Set("Authorization", "test@jpmeijers.com")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Experiment", "test-experiment")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(context.PostTtnV2)

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
