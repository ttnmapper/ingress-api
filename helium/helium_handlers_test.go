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
		`{"app_eui":"6081F9946B592C18","decoded":{"payload":{"accuracy":4,"altitude":19,"hdop":2.1,"latitude":52.21151365110359,"longitude":5.983654617289005},"status":"success"},"dev_eui":"6081F91006C1390E","devaddr":"77000048","downlink_url":"https://console.helium.com/api/v1/down/31de78f1-ee4b-48ba-8164-3fb6896832f3/Pb5ThzFBGi32kCqNbZNWE1Hzd6lop0B4/31abe016-8684-4ed7-bea9-c683373c4915","fcnt":10,"hotspots":[{"channel":4,"frequency":867.2999877929688,"hold_time":577,"id":"146AfLNcT146oVCPK17Jb7CdMxEkbhc9XkJdbT2j8afVE3NJ3Bx","lat":"unknown","long":"unknown","name":"scruffy-white-aphid","reported_at":1634672288083,"rssi":-45.0,"snr":9.199999809265137,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":720,"id":"13K5PyJDau5SKyJq9kEVurCcujF6mqc67jhSG67GVTZZ2YenkeB","lat":52.20860898339664,"long":5.983685532406203,"name":"feisty-candy-corgi","reported_at":1634672288093,"rssi":-37.0,"snr":9.5,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":1170,"id":"11JisG5QnEzkLAtDGbbdMbgA3wBUEmwpneTAp7FKDHcRKPPEHFa","lat":52.222011315157566,"long":5.993294065143459,"name":"lone-glass-wasp","reported_at":1634672288728,"rssi":-96.0,"snr":9.199999809265137,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":1389,"id":"112mdqmXBCMXg7wgiBXmtmbjZppPmY8ABZTfnE7KNdFzRhSAYbMD","lat":52.214138787048356,"long":5.968713530820866,"name":"magnificent-eggshell-wombat","reported_at":1634672288950,"rssi":-117.0,"snr":-11.0,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":1727,"id":"11gbT6L6mVDjHDyEs5RHomA4TzRjqAKP3N11Pu6y7hcjFEXgjwm","lat":52.20126564957077,"long":5.97837478274946,"name":"fluffy-lavender-gorilla","reported_at":1634672289310,"rssi":-113.0,"snr":-7.800000190734863,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":1772,"id":"112aHKWgbCiLdPUpmhKKFrwjkEKy1nq7ZsRbS8p7E2cNaUAuoSz8","lat":52.23313620594774,"long":5.975775292696945,"name":"oblong-pebble-sheep","reported_at":1634672289334,"rssi":-118.0,"snr":-10.800000190734863,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":1770,"id":"11gke1NX5FRTSdmXGoLjj34cPCRL6xJ1PGPromZHXhDYqLEdaDZ","lat":52.212277361290965,"long":5.977788077467616,"name":"helpful-arctic-meerkat","reported_at":1634672289350,"rssi":-102.0,"snr":-11.5,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2169,"id":"11D1tt5f2brYCAhtfCbbMnU4DhHTxqTFd3GQFcVSehQSYTWehKb","lat":52.20258827503975,"long":6.005519220039531,"name":"modern-clay-mole","reported_at":1634672289407,"rssi":-114.0,"snr":-13.0,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":1915,"id":"112JAPRHCjca9uWKNdgHtV436f6wv4gGK6a7EUyqX8FMyUf9yxer","lat":52.23654091796439,"long":6.000338744740332,"name":"long-taffy-mongoose","reported_at":1634672289453,"rssi":-120.0,"snr":-11.5,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2022,"id":"112G58KSGSAhcVzQGx53weEGSJvwkdrVTRwRk9vWfeNVMXSAeGAc","lat":52.19360374679962,"long":5.973726132893089,"name":"wonderful-citron-penguin","reported_at":1634672289531,"rssi":-111.0,"snr":3.0,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2024,"id":"1129VwKVBWrYRsiKdJWAZTYUDub9voYXjRvdKBbV5YU5ABn634xK","lat":52.22517160839464,"long":5.9754466285820405,"name":"tame-lemon-chameleon","reported_at":1634672289583,"rssi":-117.0,"snr":-5.0,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2105,"id":"119ENrLqRpiNaKRZhUgcacD9YTo9QvJDffxYSTUHboGs2JGJF3r","lat":52.1984371761415,"long":5.976240672895449,"name":"short-ginger-coyote","reported_at":1634672289614,"rssi":-119.0,"snr":-2.200000047683716,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":1810,"id":"112KR5qSspWdgAX2mWUwVALAVsXVXa2TdXoGDSEPLo2arwjjtuPv","lat":52.19182750779413,"long":5.983484087278604,"name":"tame-hotpink-iguana","reported_at":1634672289711,"rssi":-117.0,"snr":-14.199999809265137,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2210,"id":"1122DiCDj1pMnazKByDCY5JxSqbkHRWivw1jtXUrScppGeeh3MRN","lat":52.22049059621654,"long":5.988006016781468,"name":"jolly-glossy-worm","reported_at":1634672289731,"rssi":-107.0,"snr":4.800000190734863,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2195,"id":"11vpunHhM87ZBP8uZyxQBvyBRURBeebCfMn45WAvjfwvJ5NuHXQ","lat":52.19492885732675,"long":5.970320809585335,"name":"mythical-indigo-dragonfly","reported_at":1634672289736,"rssi":-118.0,"snr":-8.800000190734863,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2187,"id":"112sUENUaNdZwdFWw9xkcDA6wR8UVDKd3bME7c9zyM1XyhRTvESR","lat":52.21017259855869,"long":6.013147500466674,"name":"tiny-basil-zebra","reported_at":1634672289749,"rssi":-118.0,"snr":-7.800000190734863,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2228,"id":"112R5ud1nA3GkLqEurXQwvpxCH3nALrm5C3S3FYRoy98cxJ9EwHc","lat":52.21158125760912,"long":5.983786880104201,"name":"hollow-malachite-pigeon","reported_at":1634672289768,"rssi":-34.0,"snr":15.0,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2275,"id":"11N7mviee9gNmvgK2MECqq5r2MYJ6osP4tZGeud7Y4eeucyJMvk","lat":52.21502574531855,"long":5.9638331496130235,"name":"best-crimson-coyote","reported_at":1634672289855,"rssi":-112.0,"snr":4.800000190734863,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2300,"id":"112RvYJjYvw1Ub5VEG3hz8m2GjbB6ECYzNAahSdAjCcL2r6YeHzk","lat":52.20262468092881,"long":5.964930950143436,"name":"high-merlot-swift","reported_at":1634672289857,"rssi":-119.0,"snr":-11.199999809265137,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2352,"id":"112XAm2DzCWFTCzfjfcWAXBQ2pdpA5AQx71c2Z2NWM6vZWMsUn9V","lat":52.23942248504769,"long":5.980681524765956,"name":"wild-olive-scorpion","reported_at":1634672289863,"rssi":-116.0,"snr":-0.5,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2346,"id":"11jVqFqSHyNs98Dt8ywSCyGpepr3nkvtM5zMXrJrLQL3LrPLt1X","lat":52.21391647963938,"long":5.975678030286078,"name":"mini-tartan-beetle","reported_at":1634672289898,"rssi":-109.0,"snr":6.5,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2407,"id":"1127Ds21gvN7rDs5mRQW1Ff7MvWAZ6yG8XpdDRVKYU3JKBKo3XPV","lat":52.215506064113384,"long":6.020039527457058,"name":"macho-tawny-capybara","reported_at":1634672289937,"rssi":-118.0,"snr":-10.199999809265137,"spreading":"SF9BW125","status":"success"},{"channel":4,"frequency":867.2999877929688,"hold_time":2411,"id":"11jDtQnSwXP4UT1GfQbnQGsgYjnaHpRiegAacpb5eN6JwRvajVp","lat":52.25158680957983,"long":5.980933968820737,"name":"strong-mossy-sheep","reported_at":1634672289954,"rssi":-113.0,"snr":0.5,"spreading":"SF9BW125","status":"success"}],"id":"31abe016-8684-4ed7-bea9-c683373c4915","metadata":{"adr_allowed":false,"cf_list_enabled":false,"labels":[{"id":"a3d58a10-7ca9-493e-b8a6-9a21f1803793","name":"ttnmapper","organization_id":"6fcb1c75-00c1-41fb-a7d8-1af0be2cf795"},{"id":"d1de4a4b-ad3b-4811-94ae-675e619bb5e4","name":"cargo","organization_id":"6fcb1c75-00c1-41fb-a7d8-1af0be2cf795"},{"id":"d48ca6bd-83d3-40ef-9f19-16a21eb68126","name":"Mapper","organization_id":"6fcb1c75-00c1-41fb-a7d8-1af0be2cf795"}],"multi_buy":9999,"organization_id":"6fcb1c75-00c1-41fb-a7d8-1af0be2cf795"},"name":"RFSee_Mapper","payload":"ykGhhEFKABMV","payload_size":9,"port":1,"reported_at":1634672288083,"uuid":"c7fe2a63-5c22-4a40-99cb-56895937c7de"}`,
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
