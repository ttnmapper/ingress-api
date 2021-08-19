package helium

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
		`{"app_eui":"70B3D57ED0019B7F","decoded":{"payload":{"accuracy":3,"altitude":-6,"ant":2,"latitude":52.2118033,"longitude":5.9836761},"status":"success"},"dev_eui":"1D4A7D00003F9D6B","devaddr":"A9020048","downlink_url":"https://console.helium.com/api/v1/down/fefc30b0-8b6c-4100-882d-59fa70a853e4/OgdYHinBMIczRSD6icmaRnsxZiUYvEH7/40b8e98b-2ff1-493d-864d-c3ee794dcb24","fcnt":93,"hotspots":[{"channel":2,"frequency":868.5,"hold_time":549,"id":"11JisG5QnEzkLAtDGbbdMbgA3wBUEmwpneTAp7FKDHcRKPPEHFa","lat":52.222011315157566,"long":5.993294065143459,"name":"lone-glass-wasp","reported_at":1629294297659,"rssi":-88.0,"snr":7.800000190734863,"spreading":"SF7BW125","status":"success"}],"id":"40b8e98b-2ff1-493d-864d-c3ee794dcb24","metadata":{"adr_allowed":false,"cf_list_enabled":false,"labels":[{"id":"5f2d2530-0fbd-493d-86d3-0a55c8d53bcf","name":"mapper","organization_id":"dd06ec79-9cb3-42a8-967f-d8c7a9225c50"},{"id":"69608d29-b773-4d1a-aba9-c1e079a96876","name":"cargo","organization_id":"dd06ec79-9cb3-42a8-967f-d8c7a9225c50"},{"id":"f8ccecef-cc8f-43ef-864b-79ca587ddc25","name":"glamos","organization_id":"dd06ec79-9cb3-42a8-967f-d8c7a9225c50"}],"multi_buy":9999,"organization_id":"dd06ec79-9cb3-42a8-967f-d8c7a9225c50"},"name":"RFSee Walker","payload":"ykG8hEFL//oC","payload_size":9,"port":1,"reported_at":1629294297659,"uuid":"594a0090-4d80-40f9-b462-7db719445f80"}`,
	}

	for _, postbody := range postbodies {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("POST", "", strings.NewReader(postbody))
		if err != nil {
			t.Fatal(err)
		}

		// Set request headers
		req.Header.Set("Host", "private.ttnmapper.org")
		req.Header.Set("User-Agent", "hackney/1.15.2")
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set("TTNMAPPERORG-USER", "test@ttnmapper.org")
		req.Header.Set("TTNMAPPERORG-EXPERIMENT", "test-experiment")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(context.PostHelium)

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
