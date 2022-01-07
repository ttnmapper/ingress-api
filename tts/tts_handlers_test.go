package tts

import (
	"bytes"
	b64 "encoding/base64"
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
		//`{"end_device_ids":{"device_id":"cricket-002","application_ids":{"application_id":"jpm-crickets"},"dev_addr":"260BEAC2"},"correlation_ids":["as:up:01F55DCJ0AAJPSX90MWGKX0A41","gs:conn:01F4S0KGVPX2QJQFH6K8KE47FR","gs:up:host:01F4S0KGZTYC78NWW5Q86GK1S7","gs:uplink:01F55DCHS8EA3P0HBMKPV9QP0V","ns:uplink:01F55DCHSKX3D9CYNKRV6RSG3W","rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01F55DCHSJRM33M74T3WG5HPYF","rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01F55DCJ0A688GR1D7QZWH545S"],"received_at":"2021-05-08T07:17:07.723112244Z","uplink_message":{"f_cnt":23171,"frm_payload":"","decoded_payload":{},"rx_metadata":[{"gateway_ids":{"gateway_id":"eui-000080029c09dd87","eui":"000080029C09DD87"},"timestamp":1408363515,"rssi":-27,"channel_rssi":-27,"snr":10.5,"location":{"latitude":-33.93667538260562,"longitude":18.871081173419956,"source":"SOURCE_REGISTRY"},"uplink_token":"CiIKIAoUZXVpLTAwMDA4MDAyOWMwOWRkODcSCAAAgAKcCd2HEPvXx58FGgwI8/XYhAYQx8nF7AEg+JiHyP6FXw==","channel_index":3},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01F55DCHTCMTM0SSJ1RSN3BBJ9","forwarder_net_id":"000013","forwarder_tenant_id":"ttn","forwarder_cluster_id":"ttn-v2-eu-3","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1","hops":[{"received_at":"2021-05-08T07:17:07.532727295Z","sender_address":"40.113.68.198","receiver_name":"router-dataplane-57d9d9bddd-dsrjj","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.533400574Z","sender_name":"router-dataplane-57d9d9bddd-dsrjj","sender_address":"forwarder_uplink","receiver_name":"router-5b5dc54cf7-psxlt","receiver_agent":"pbrouter/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.534011884Z","sender_name":"router-5b5dc54cf7-psxlt","sender_address":"deliver.000013_ttn_ttn-eu1.uplink","receiver_name":"router-dataplane-57d9d9bddd-f7h6k","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"}]},"rssi":-25,"channel_rssi":-25,"snr":9.5,"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU2sxaVJXeEVaRmMxV0dSRWFFZE9Sa0p4VlVaa1NVbHBkMmxrUjBadVNXcHZhVlV3ZHpSU2JsSkpWRzA1V2sweVNUTldiR2hDVWxVeFVsVlZTbEphZVVvNUxtc3RlRTh0WkhwNlJtRkdUbTV0VlZCWVZXOWtNMUV1YTBSaVZHMUJXbXhVZVROWFdtbEpVaTVoYldwVVpVZHdVVTFKWVZWT1RsSnRTR3BKWW5seFJrcFpZMUI2WDB4dVdsOUlUalJpWVcxR1psTmxRbTV3TTAxYU56a3hXblk1ZUdFMVV6QlFVbEJXYzBKbmExWk5UV1psVmxGRFRWSnRSM0JvWkdFMloxZEROMkZtWlZSbk9FVkdkbEUzWVRSelZrbzROMXBEWDJKeGIwbFJjbTFZTm1WU2JGOHdaaTFrYUZwU2RVbzBlRVZTZG1kRVUwbE1PRGxmT0dGQlNUa3lhVGg1YzJSeFpGOXdMWEUwUkU5aVh6TTRUM3BITG5ZeVpXcHFSM0Z5U2kwMUxWaHJlVTVTWkZwblIyYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19"},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01F55DCHTK8M44A02PSWXSTNF4","forwarder_net_id":"000013","forwarder_tenant_id":"ttn","forwarder_cluster_id":"ttn-v2-eu-3","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1","hops":[{"received_at":"2021-05-08T07:17:07.539820178Z","sender_address":"40.113.68.198","receiver_name":"router-dataplane-57d9d9bddd-xjszp","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.540236935Z","sender_name":"router-dataplane-57d9d9bddd-xjszp","sender_address":"forwarder_uplink","receiver_name":"router-5b5dc54cf7-mwf8m","receiver_agent":"pbrouter/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.540851394Z","sender_name":"router-5b5dc54cf7-mwf8m","sender_address":"deliver.000013_ttn_ttn-eu1.uplink","receiver_name":"router-dataplane-57d9d9bddd-f7h6k","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"}]},"rssi":-28,"channel_rssi":-28,"snr":8.5,"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU2tWa2JWbzFZVVJSTTFSdGN6Tk5NMG95V2pJMWMwbHBkMmxrUjBadVNXcHZhVTlJVms5a01rcEhWMVpXWVdGclZubGpWemx5VGxWb1dWTXpSbFJSVTBvNUxtaEVZelJGUlhaNGMyUXpaVEZuYjBvM2VYRXRjRUV1YXpkYWJ6Um5UamRRVjNsdFREWkxkUzVRVTBzMk9XdGpaMlpyY0habGRuUjZRa3hJWTBkcVJIZEhWSFY1YW5NdGRWWmlRbFozYlRFMGJuSkVkR2xZTVdOYVNHUkpaa05WVUUxdmNEVkxNSFZ5T0RsdlZHUTRkMWhMY1VsWFIzaFNNemxvWlhaVGJXbFJNbWhPUWs0d1F6VlhNRWwyYnkwNU5WbzBXbGcxYjJOSWVrNVVNVmRVTTNwUFNFMWxlREF4WnprMk5GUnJUMmN4YkdaQlZVNVJTWEZ0YVRsVmNFbzBiemhFZDBSM2QwODBhM3BOTTNGUmN6WlVaVWhvTGsxNmExRnJaMlUwYXpoYVRtWllUWHBIVTBkQ1puYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19"},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01F55DCHVF2ES39GECE99FPC8M","forwarder_net_id":"000013","forwarder_tenant_id":"ttn","forwarder_cluster_id":"ttn-v2-eu-1","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1","hops":[{"received_at":"2021-05-08T07:17:07.567259072Z","sender_address":"52.169.73.251","receiver_name":"router-dataplane-57d9d9bddd-f7h6k","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.568763883Z","sender_name":"router-dataplane-57d9d9bddd-f7h6k","sender_address":"forwarder_uplink","receiver_name":"router-5b5dc54cf7-psxlt","receiver_agent":"pbrouter/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.570193353Z","sender_name":"router-5b5dc54cf7-psxlt","sender_address":"deliver.000013_ttn_ttn-eu1.uplink","receiver_name":"router-dataplane-57d9d9bddd-dsrjj","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"}]},"time":"2021-05-08T07:17:06.010555982Z","rssi":-33,"channel_rssi":-33,"snr":9.5,"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU2pWV2JXeHJVekl4ZW1WcVJteGtWM0ExV2tWd2JrbHBkMmxrUjBadVNXcHZhV0p1YUhWaldHaHJVbXMxVWs5RVZUVmFha1pPVkd4S1MyVnVTWGhhZVVvNUxraHNablpYU1MxcVNuVjBNVzFqT0VkUlNVWmhURUV1ZURaV1FqVjVPVGN0Ykdob1pXOHhXUzVyV0doSFpYazBXbTR4Um5aaVZXbzVaMDV4Tmw5dFNFVTJSMVp2T0ZsVVZWTkhRVVYwWm1aWVFYSllZbEJKZFVGeFYyaEhha3BVTTBGdVJtSm5SblptVEZoSExVSnhhMkpOYUZGbFZFVTBlbFpYWkRCSVZHSnpWSE5oU3pKYWR6bFBOelV3VTJoWFluVTVTa3B0VjAxUmNGQk1XR052VFhGVVJsRmhlV3A0TURGSk1uTlRRMUZHY2tGS2NsZE9lblpmTjI1aldYbzJaRTlwTm00MWJrTTNPRk56ZVZwRk5UbFJSakZETGpRNWNVUTFVRmRPU1d3NWFHRmFlbU5oYjNWS1ptYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0xIn19"},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01F55DCHVTK41WZ054G87YDX87","forwarder_net_id":"000013","forwarder_tenant_id":"ttn","forwarder_cluster_id":"ttn-v2-eu-4","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1","hops":[{"received_at":"2021-05-08T07:17:07.578226151Z","sender_address":"52.169.150.138","receiver_name":"router-dataplane-57d9d9bddd-xjszp","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.582886038Z","sender_name":"router-dataplane-57d9d9bddd-xjszp","sender_address":"forwarder_uplink","receiver_name":"router-5b5dc54cf7-xh822","receiver_agent":"pbrouter/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.585417009Z","sender_name":"router-5b5dc54cf7-xh822","sender_address":"deliver.000013_ttn_ttn-eu1.uplink","receiver_name":"router-dataplane-57d9d9bddd-f7h6k","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"}]},"rssi":-67,"channel_rssi":-67,"snr":10,"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU1RWalJHeDJaVlUxUjJFeVNtOWhSVTVYVGxWYVFrbHBkMmxrUjBadVNXcHZhV0V6UVRCU2JHUkpaV3BDU1dGWFpFVlNNMDB6V2pOc05XVkZUbTlWVTBvNUxraG9iRFpaYmpobWRqQm9jakpuVEc5T09HSTVaVkV1VjFCVFJuUktUa1JCUW1Gc05UWlRXaTV4TTJSeE9VOUlhRlU0WTI4MVJtRnJUazEwYkdKMGQyRlpRVWhKY0cxVGMySmFUVFpRYkRSb2JtVXpiVmxrUkhCNE9YTkRTaTFCWkMxTlZXSmpaell6V0VoWVFuUnRWM0JuYm1KZmFXNURRbUpYUnpGNFNYVnJiM0JxUWtkWVQyUndMUzFuVm5CNVoyWkZNbmhIY1dWS1dIRXdaMnBSTldNd2MxWnVUbGd0WjJsRWIyVnRSRjlDYkcxaU1XUjNNR0o2Y1ZsWE4ybEZRMVoyVUhBNWNESnFjVVprYm5SV2QyUmZNVFZWTG5aWVdISjNjbEU0ZEdKeVQwTnllbmw1TTA5SlZrRT0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS00In19"}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":7}},"data_rate_index":5,"coding_rate":"4/5","frequency":"867100000","timestamp":1408363515},"received_at":"2021-05-08T07:17:07.507008642Z","consumed_airtime":"0.041216s","locations":{"user":{"latitude":-33.93623477040523,"longitude":18.871655166149143,"source":"SOURCE_REGISTRY"}}}}`,
		//`{"end_device_ids":{"device_id":"cricket-001","application_ids":{"application_id":"jpm-crickets"},"dev_addr":"26011CE4"},"correlation_ids":["as:up:01E175D2K6EHZH7GGH9TWRCVBN","gs:conn:01E16YPNYG4HEXHYJ7VFYKH2EW","gs:uplink:01E175D2AYR39QT12BY0ESMPP7","ns:uplink:01E175D2AZPJF4RDZH7A5EP2BS","rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01E175D2AYJYFSCZ6NMXKJ2QWQ"],"received_at":"2020-02-16T14:10:59.302096081Z","uplink_message":{"f_port":1,"f_cnt":527,"frm_payload":"AIj60lkC4SQAMY8=","decoded_payload":{"gps_0":{"altitude":126.87000274658203,"latitude":-33.93669891357422,"longitude":18.870800018310547}},"rx_metadata":[{"gateway_ids":{"gateway_id":"pisupply-shield","eui":"B827EBFFFED88375"},"timestamp":2732493451,"rssi":-72,"channel_rssi":-72,"snr":9.8,"uplink_token":"Ch0KGwoPcGlzdXBwbHktc2hpZWxkEgi4J+v//tiDdRCLlfqWCg=="}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":7}},"data_rate_index":5,"coding_rate":"4/5","frequency":"868100000","timestamp":2732493451},"received_at":"2020-02-16T14:10:59.039048589Z"}}`,
		//`{"end_device_ids":{"device_id":"cricket-002","application_ids":{"application_id":"jpm-crickets"},"dev_addr":"260BEAC2"},"correlation_ids":["as:up:01FD2A6GEN2F1VY4GEEVCJTM62","gs:conn:01FD1ZZQ1T6CEAMJ29PVCAKKF8","gs:up:host:01FD1ZZQ226GN31YFKSCN3SN15","gs:uplink:01FD2A6G7ZJHR911MNB5G44EYF","ns:uplink:01FD2A6G81E8TH80T8GJ9WXASM","rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01FD2A6G80R72JW4DY3QHN4N4K","rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01FD2A6GEM3Q5TWEJ8VPKQY14F"],"received_at":"2021-08-14T12:29:15.095206575Z","uplink_message":{"f_cnt":33540,"frm_payload":"","decoded_payload":{},"rx_metadata":[{"gateway_ids":{"gateway_id":"eui-000080029c09dd87","eui":"000080029C09DD87"},"timestamp":4266137228,"rssi":-33,"channel_rssi":-33,"snr":8.5,"location":{"latitude":-33.93667538260562,"longitude":18.871081173419956,"source":"SOURCE_REGISTRY"},"uplink_token":"CiIKIAoUZXVpLTAwMDA4MDAyOWMwOWRkODcSCAAAgAKcCd2HEIytoPIPGgwImu7eiAYQw9akowMg4KXgzJT2Ag==","channel_index":7},{"gateway_ids":{"gateway_id":"eui-58a0cbfffe80049a","eui":"58A0CBFFFE80049A"},"time":"2021-08-14T12:29:14.725533962Z","timestamp":2108938836,"rssi":-41,"channel_rssi":-41,"snr":8.75,"uplink_token":"CiIKIAoUZXVpLTU4YTBjYmZmZmU4MDA0OWESCFigy//+gASaENS0z+0HGgwImu7eiAYQq8H1qQMgoLCztLC3Ag=="},{"gateway_ids":{"gateway_id":"eui-60c5a8fffe71a964","eui":"60C5A8FFFE71A964"},"timestamp":3259758379,"rssi":-74,"channel_rssi":-74,"snr":8.8,"uplink_token":"CiIKIAoUZXVpLTYwYzVhOGZmZmU3MWE5NjQSCGDFqP/+calkEKvur5IMGgwImu7eiAYQ+va7sAMg+P/1xe+/NA==","channel_index":4},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01FD2A6GA8GDGP8GR539JHB45H","forwarder_net_id":"000013","forwarder_tenant_id":"ttnv2","forwarder_cluster_id":"ttn-v2-eu-3","forwarder_gateway_eui":"647FDAFFFE007A1F","forwarder_gateway_id":"eui-647fdafffe007a1f","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1"},"rssi":-37,"channel_rssi":-37,"snr":8.2,"location":{"latitude":-33.93626794,"longitude":18.87168703},"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU1RWTlZrcG1WVmR3U2xReFJuSmtWVlpoVTBkR1ZVbHBkMmxrUjBadVNXcHZhV050YkhWbFJXeHhUMWhzWm1KSFRubFJibFpSVkVka2JFOUZSalpSVTBvNUxrSkhha3hVTURkaE1GQnVUVEpCVTBWUlluTjRWVUV1WW5NdFZXbFRTMjV2WjBaa01VbEdaQzVSUXpGNVdtdFpTV3BrU1hKbk1TMXFUalpUWkhsRFRsTTRNMjVSYWpFNVMxaHBRemhHTFVOVmJGaE9SbGM0TUdWWU1tWklZVVZoUlZCUFRURk5PSEJTZUVjdGJWQkRNMkpUTFhaUU9EaE9ZamwxVFhkUllsVlVkM0IwUTFaNmFrMXdSR05TVDBaNFZGVjROazVoTmt3MlVISktObUZ1YVdGUFZraGtWVXBrUW5Oc2RHOTZjV05qU21rNVprVlBZemMxWDBaUlVHeExRV3hYUVZwTFZrUnNNMWxSVmpkME5VOTZaeTVJY21wdVFrVTVkRlJTYzJKTGFEUTRZWFJaY2taUiIsImEiOnsiZm5pZCI6IjAwMDAxMyIsImZ0aWQiOiJ0dG52MiIsImZjaWQiOiJ0dG4tdjItZXUtMyJ9fQ=="}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":7}},"data_rate_index":5,"coding_rate":"4/5","frequency":"867900000","timestamp":4266137228},"received_at":"2021-08-14T12:29:14.881054809Z","consumed_airtime":"0.041216s","network_ids":{"net_id":"000013","tenant_id":"ttn","cluster_id":"ttn-eu1"}}}`,
	}

	for _, postbody := range postbodies {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("POST", "", strings.NewReader(postbody))
		if err != nil {
			t.Fatal(err)
		}

		// Set request headers
		req.Header.Set("TTNMAPPERORG-USER", "test@ttnmapper.org")
		req.Header.Set("TTNMAPPERORG-EXPERIMENT", "test-experiment")
		req.Header.Set("COntent-Type", "application/json")
		req.Header.Set("X-Tts-Domain", "test.cloud.thethings.network")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(context.PostV3Uplink)

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
			t.Error("No value read, moving on.")
		}
	}
}

func TestHandlerProtobuf(t *testing.T) {
	var publishChannel = make(chan types.TtnMapperUplinkMessage, 1)
	context := &Context{PublishChannel: publishChannel}

	// Create a request to pass to our handler.
	data, _ := b64.StdEncoding.DecodeString("CiMKC2NyaWNrZXQtMDAyEg4KDGpwbS1jcmlja2V0czIEJgvqwhIgYXM6dXA6MDFGNTZGRThKMkdOTjI5Vlc1V0gyVlNNNFISImdzOmNvbm46MDFGNFMwS0dWUFgyUUpRRkg2SzhLRTQ3RlISJWdzOnVwOmhvc3Q6MDFGNFMwS0daVFlDNzhOV1c1UTg2R0sxUzcSJGdzOnVwbGluazowMUY1NkZFOEJBNEpUSDBQQVFHSE5LUDZGUBIkbnM6dXBsaW5rOjAxRjU2RkU4QkJWWU5TRzJUVlEyRFhSUkpOEkBycGM6L3R0bi5sb3Jhd2FuLnYzLkdzTnMvSGFuZGxlVXBsaW5rOjAxRjU2RkU4QkJSOFQzMzFSNDU3RkJaRVQyEkBycGM6L3R0bi5sb3Jhd2FuLnYzLk5zQXMvSGFuZGxlVXBsaW5rOjAxRjU2RkU4SjE5VEJFV0tYSlNEQllNUDFLYgsI74zbhAYQwdG4URqwIhiAuQEqADKSAQogChRldWktMDAwMDgwMDI5YzA5ZGQ4NxIIAACAApwJ3Ycgu7KZogpFAACowU0AAKjBXZqZ2UBqFAk1oZv65PdAwBEBAAAt/94yQCgDekAKIgogChRldWktMDAwMDgwMDI5YzA5ZGQ4NxIIAACAApwJ3YcQu7KZogoaDAjujNuEBhCwwZPHAyD4nLKWm5VniAECMo4ICg4KDHBhY2tldGJyb2tlckUAANjBTQAA2MFdzcwcQXqcBHsiZyI6IlpYbEthR0pIWTJsUGFVcENUVlJKTkZJd1RrNVRNV05wVEVOS2JHSnRUV2xQYVVwQ1RWUkpORkl3VGs1SmFYZHBZVmhaYVU5cFNuaFVhM0JZVkZSYVRrMHdNVXROYkZwelRrUk5NVWxwZDJsa1IwWnVTV3B2YVU5RVVtNU5SV3hQVEZZNVlXUkhUbTVQUjBwelRWUldVVkpWZUhwUlUwbzVMbVJKY0hCbVJrRllUVVJVUjFaMWIwVlBNSGh2YlhjdVpGaDFXR3BWVWpWc1YxRjVlRXRKYmk1bk9IUmpORGxwYjI5b1NtbFZVakZCTkZWMllrZFZjekUxVjI1WmVWaEdSRXBKVTNCSVFrODVlWHBpWjNwTlFscEpSbmREUVhsRlFuTnRVbVJWWjFsNVZFSnVjWEJDYTNCYVJIRjZXRlZrUTNVMGRYcE1NM05yV1ZOYVpGVTJYMWRVUkRaQmFWbHVORzFoZDFoaGQwNW1ZVlpDUmpOdlMzSnNSMVZ3UTBJM1FrOVRhRGRGZFdob2QxbHVORXd6ZFhwRGRYaFZSVkJtY0VKek9HZE1ORE5OYWxGRVlVazBVbmRrVWtoVUxtaElPRUZSY0U1dFNXc3lTemwxWVVWRVFqWTJWM2M9IiwiYSI6eyJmbmlkIjoiMDAwMDEzIiwiZnRpZCI6InR0biIsImZjaWQiOiJ0dG4tdjItZXUtMyJ9fZIBzAMKGjAxRjU2RkU4Q1cxS0JDN1IxUVpFUUFWMEZEEgMAABMaA3R0biILdHRuLXYyLWV1LTMqAwAAEzIDdHRuOmgKCwjvjNuEBhDAnPUBGg00MC4xMTMuNjguMTk4IiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQteGpzenAqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqBAQoLCO+M24QGEOmFlAISIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC14anN6cBoQZm9yd2FyZGVyX3VwbGluayIXcm91dGVyLTViNWRjNTRjZjctcHN4bHQqJHBicm91dGVyLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqVAQoLCO+M24QGEIL51AISF3JvdXRlci01YjVkYzU0Y2Y3LXBzeGx0GiFkZWxpdmVyLjAwMDAxM190dG5fdHRuLWV1MS51cGxpbmsiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0Qgd0dG4tZXUxMo8ICg4KDHBhY2tldGJyb2tlckUAAIrCTQAAisJdmpnZQHqcBHsiZyI6IlpYbEthR0pIWTJsUGFVcENUVlJKTkZJd1RrNVRNV05wVEVOS2JHSnRUV2xQYVVwQ1RWUkpORkl3VGs1SmFYZHBZVmhaYVU5cFNuUlRXR2gwVjFjMVNrMHdaSEJpV0ZFd1pGUnJORWxwZDJsa1IwWnVTV3B2YVdGWVRrcGxWMmhTWXpOQ1dtUXhVbGhoTWpGUFdtMDFVV0p0TlRWYWVVbzVMa2RTTTFoS1EyTlZTelpzTm14SmJrNWZiMWx5ZEhjdWVHWTVkVnBtTVZWblZ5MTNNMjl3Umk1SFoyNU9OR055UVRVM1RIVlpYMjVMWVRWcFVHd3hSMjVOZFdwVE5GWndTVE5uTUdsNWNqSk9TRzA0TlVKNFEyTXhPSGs0VVdjdGFrVnVPWFpsUXpsVU1GZFZRVTVTUVhwbFFsUllkbE5wYWtGbWVWWmFUR2xNU1dkVllqaHBTa3h3ZGxwd1ZGVXdSVVpKTUZONU1WZEZTMHRUTVcwNGR6VjJXRU5ZV1ZwWWVYVkphV3Q2Ykc1ak0zWkJha0UyYmxKT2VXSlpaVFpJYm5CSmNYSk1OMHBpYjBNeWFHZFJTVmMxUldObkxrWk5RWGwyY0ZOMVZtcDFVWGc0WjFacVVrWkdNRkU9IiwiYSI6eyJmbmlkIjoiMDAwMDEzIiwiZnRpZCI6InR0biIsImZjaWQiOiJ0dG4tdjItZXUtNCJ9fZIBzQMKGjAxRjU2RkU4Q1o2WUU3U1Q0SkJKV0I0TVdEEgMAABMaA3R0biILdHRuLXYyLWV1LTQqAwAAEzIDdHRuOmkKCwjvjNuEBhDxoLgDGg41Mi4xNjkuMTUwLjEzOCIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6gQEKCwjvjNuEBhCpgr4EEiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoaEGZvcndhcmRlcl91cGxpbmsiF3JvdXRlci01YjVkYzU0Y2Y3LXBzeGx0KiRwYnJvdXRlci8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6lQEKCwjvjNuEBhCxjoUFEhdyb3V0ZXItNWI1ZGM1NGNmNy1wc3hsdBohZGVsaXZlci4wMDAwMTNfdHRuX3R0bi1ldTEudXBsaW5rIiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZjdoNmsqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NEIHdHRuLWV1MTKGCAoOCgxwYWNrZXRicm9rZXJFAAAUwk0AABTCXQAAAEF6lAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTbkZOVjJSRlRraFpkMk5HV2xOaGVsVXdWMGhrY0VscGQybGtSMFp1U1dwdmFXUnBNVk5TYmtKWVkydG9WazFHWkZGTlNFb3lZVzFPZVdWR1pFbGFlVW81TG10S1dIQkxPVzlSZUVzMGMwbExUekZmZFhGRlpuY3ViRFI2TmxkNlkwWTVOSFl3UWxCNFZDNVVja1prUm5WMk9IUjFNVkIzUTFKdGIzTjBObmczVW5saFJYQjVZVkozVlZoc1lsRnFNSEU1WVRodmQxRTBiSGszTFdsRU1WUnJXamhLUTFJd1ZYUm5TM3A2WXpKVFoyMVlSbGw0VkdaUk4wMU9OMHRGY3pnMVNrdHVUbXRNU0hSQ09WcFBTV2h1VVRSWlVtUkNSVnBVU1VsVGNreDVRVWRPZVU1eWFXZEJRMmhmVUMxQmQxRkdUMjFNTVcxdVdFWnRhRXhUVlRoVlRqVnhURW81VFdGcmNXRXpYM040T0M1QmNHbHlURFo1TlhWSldHWkdaVGhHYm5oYWJIbFIiLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19kgHMAwoaMDFGNTZGRThEOTMzRFRQNUg5WDk2TTdFN0sSAwAAExoDdHRuIgt0dG4tdjItZXUtMyoDAAATMgN0dG46aAoLCO+M24QGEMu9kAgaDTQwLjExMy42OC4xOTgiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoEBCgsI74zbhAYQ/fz0DBIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqGhBmb3J3YXJkZXJfdXBsaW5rIhdyb3V0ZXItNWI1ZGM1NGNmNy14aDgyMiokcGJyb3V0ZXIvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OpUBCgsI74zbhAYQ4bHHEBIXcm91dGVyLTViNWRjNTRjZjcteGg4MjIaIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTEylAgKDgoMcGFja2V0YnJva2VyGgwI7YzbhAYQ9e+6rgFFAAD4wU0AAPjBXQAA4EB6lAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTVFJhYlU1Q1lXc3dlR0ZHY0ZGamVscFlZMGRvTTBscGQybGtSMFp1U1dwdmFWTjZXVFJqUjJ3elRucE9hRkZYTVVoaVJUbFdUMVZ6TlZGcWFHMVJVMG81TGxGRVNISlhXbDk2T1Rad1VqSk9NbkZ2WkRaak5tY3VOVmcxZDE5R1ZFaFNTMVZrV1d0VWNpNWlVa0pMZHpWRE1GSnplWE5TVURGdVJUSldiMlZNZW5KT1JXeEdVMjkyY1ZCTWJURlFTWFJNWTJJMmFtOXZVRlZJTjFWNlQyVldNRk5zVlZsaFl6RndVRTkzYm1oNFJXTmlhMDFRUTBkV1dESmZUamRtU1ZKTmFqWTVjSGhKWjA1WVdqTnpOMnhwYUY5a1pVRmpTa3RsYWxKMU1sWkNXbVJCT0hGd1RGQk5hbkY2Y0RkWFRYcENXa3RYY2xSZk4xaFlTekUwVlRKNVNqWnFkVFIwVG5OVlpIVk1PRlZwWXk1elRISm5iQzFJUlRWc1RFRlpOblF6UVhoc2NFTlIiLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0xIn19kgHMAwoaMDFGNTZGRThEOVZQNk04SldIN0JERzVCSkQSAwAAExoDdHRuIgt0dG4tdjItZXUtMSoDAAATMgN0dG46aAoLCO+M24QGELGqqggaDTUyLjE2OS43My4yNTEiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoEBCgsI74zbhAYQ1dG5DBIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqGhBmb3J3YXJkZXJfdXBsaW5rIhdyb3V0ZXItNWI1ZGM1NGNmNy1td2Y4bSokcGJyb3V0ZXIvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OpUBCgsI74zbhAYQp97DEBIXcm91dGVyLTViNWRjNTRjZjctbXdmOG0aIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTE6HQoICgYIyNAHEAcQBRoDNC81IKCEkZ4DMLuymaIKQgwI7ozbhAYQgZ3TxwNqBRCA0NMTchwKBHVzZXISFAn+KrKiWvpAwBEBAABQoXA2QCgD")
	req, err := http.NewRequest("POST", "", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	// Set request headers
	req.Header.Set("TTNMAPPERORG-USER", "test@ttnmapper.org")
	req.Header.Set("Content-Type", "application/octet-stream") // TTS uses octet-stream for protobufs
	req.Header.Set("X-Tts-Domain", "test.cloud.thethings.network")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(context.PostV3Uplink)

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
