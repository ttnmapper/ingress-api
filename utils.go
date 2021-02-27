package main

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"ttnmapper-ingress-api/ttnV2"
	"ttnmapper-ingress-api/ttsV3/models"
	"ttnmapper-ingress-api/types"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func AppendToFile(filename string, line string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(line)
	if err != nil {
		return err
	}

	return nil
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func CopyTtnV2Fields(packetIn ttnV2.UplinkMessage, packetOut *types.TtnMapperUplinkMessage) {
	/*
		V2
		  "app_id": "jpm_mapping_nodes",
		  "dev_id": "rfm_teensy3_b",
		  "hardware_serial": "00F26916ED6E43B4",
	*/
	packetOut.AppID = packetIn.AppID
	packetOut.DevID = packetIn.DevID
	packetOut.DevEui = packetIn.HardwareSerial

	/*
		V2
		  "metadata": {
		    "time": "2018-01-16T15:35:09.315649867Z",
		    "gateways": [
		      {
		        "time": "2018-01-16T15:35:09Z",
	*/
	rxTime := packetIn.Metadata.Time
	packetOut.Time = time.Time(rxTime).UnixNano()

	/*
	   V2
	        "port": 1,
	        "counter": 388,
	*/
	packetOut.FPort = packetIn.FPort
	packetOut.FCnt = int64(packetIn.FCnt)

	/*
	   V2
	        "metadata": {
	          "frequency": 868.3,
	          "modulation": "LORA",
	          "data_rate": "SF7BW125",
	          "coding_rate": "4/5",
	*/
	// We need to round the frequency as the floating point type otherwise gives us strange decimals
	frequency := math.Round(float64(packetIn.Metadata.Frequency) * 1000000) // MHz to Hz
	packetOut.Frequency = uint64(frequency)

	packetOut.Modulation = packetIn.Metadata.Modulation

	if packetOut.Modulation == "LORA" {

		drParts := strings.Split(packetIn.Metadata.DataRate, "BW")
		bandwidth, _ := strconv.Atoi(drParts[1])
		packetOut.Bandwidth = uint64(bandwidth * 1000) // kHz to Hz
		sf, _ := strconv.Atoi(strings.TrimPrefix(drParts[0], "SF"))
		packetOut.SpreadingFactor = uint8(sf)
	}

	if packetOut.Modulation == "FSK" {
		packetOut.Bitrate = uint64(packetIn.Metadata.Bitrate)
	}

	packetOut.CodingRate = packetIn.Metadata.CodingRate

	/*
	   V2
	      {
	        "metadata": {
	          "gateways": [
	*/
	for _, gatewayIn := range packetIn.Metadata.Gateways {
		/*
			{
				"gtw_id": "stellenbosch-technopark",
				"gtw_trusted": true,
				"timestamp": 3221941195,
				"time": "2018-01-16T15:35:09Z",
				"fine_timestamp"
				"fine_timestamp_encrypted"
				"channel": 1,
				"rssi": -120,
				"snr": -6.5,
				"rf_chain": 1,
				"antenna"
				"latitude": -33.96445,
				"longitude": 18.836777,
				"altitude": 150
				"location_accuracy"
				"location_source"
			}
		*/
		gatewayOut := types.TtnMapperGateway{}

		gatewayOut.GatewayId = gatewayIn.GtwID

		// If the id is eui-deadbeef, strip the prefix, capitalize and use as EUI
		if strings.HasPrefix(gatewayIn.GtwID, "eui-") && len(gatewayIn.GtwID) == 20 {
			eui := strings.TrimPrefix(gatewayIn.GtwID, "eui-")
			eui = strings.ToUpper(eui)
			gatewayOut.GatewayEui = eui
		}

		gatewayOut.Timestamp = gatewayIn.Timestamp
		gatewayTime := time.Time(gatewayIn.Time)
		if !gatewayTime.IsZero() {
			gatewayOut.Time = gatewayTime.UnixNano()
		}

		gatewayOut.FineTimestamp = gatewayIn.FineTimestamp
		gatewayOut.FineTimestampEncrypted = gatewayIn.FineTimestampEncrypted
		gatewayOut.ChannelIndex = gatewayIn.Channel
		gatewayOut.Rssi = gatewayIn.RSSI
		gatewayOut.Snr = gatewayIn.SNR
		gatewayOut.AntennaIndex = gatewayIn.Antenna
		gatewayOut.Latitude = float64(gatewayIn.Latitude)
		gatewayOut.Longitude = float64(gatewayIn.Longitude)
		gatewayOut.Altitude = gatewayIn.Altitude
		gatewayOut.LocationAccuracy = gatewayIn.Accuracy
		gatewayOut.LocationSource = gatewayIn.Source

		packetOut.Gateways = append(packetOut.Gateways, gatewayOut)
	}
}

func CopyTtnV3Fields(packetIn models.V3ApplicationUp, packetOut *types.TtnMapperUplinkMessage) {
	/*
		V3
		  "end_device_ids" : {
		    "device_id" : "dev1",                    // Device ID
		    "application_ids" : {
		      "application_id" : "app1"              // Application ID
		    },
		    "dev_eui" : "0004A30B001C0530",          // DevEUI of the end device
		    "join_eui" : "800000000000000C",         // JoinEUI of the end device (also known as AppEUI in LoRaWAN versions below 1.1)
		    "dev_addr" : "00BCB929"                  // Device address known by the Network Server
		  },
	*/
	packetOut.AppID = packetIn.EndDeviceIds.ApplicationIds.ApplicationID
	packetOut.DevID = packetIn.EndDeviceIds.DeviceID
	packetOut.DevEui = packetIn.EndDeviceIds.DevEui

	/*
		V3
		  "received_at" : "2020-02-12T15:15..."      // ISO 8601 UTC timestamp at which the message has been received by the Application Server
		  "uplink_message" : {
		    "rx_metadata": [{                        // A list of metadata for each antenna of each gateway that received this message
		      "time": "2020-02-12T15:15:45.787Z",    // ISO 8601 UTC timestamp at which the uplink has been received by the gateway
		    }],
		    "settings": {                            // Settings for the transmission
		      "time": "2020-02-12T15:15:45.787Z"     // ISO 8601 UTC timestamp at which the uplink has been received by the gateway
		    },
		    "received_at": "2020-02-12T15:15..."     // ISO 8601 UTC timestamp at which the uplink has been received by the Network Server
	*/
	packetOut.Time = time.Time(packetIn.ReceivedAt).UnixNano()

	/*
	   V3
	         "uplink_message":{
	            "f_port":1,
	            "f_cnt":527,
	*/
	packetOut.FPort = uint8(packetIn.UplinkMessage.FPort)
	packetOut.FCnt = packetIn.UplinkMessage.FCnt

	/*
	   V3
	      "uplink_message" : {
	        "settings": {                            // Settings for the transmission
	          "data_rate": {                         // Data rate settings
	            "lora": {                            // LoRa modulation settings
	              "bandwidth": 125000,               // Bandwidth (Hz)
	              "spreading_factor": 7              // Spreading factor
	            }
	          },
	          "data_rate_index": 5,                  // LoRaWAN data rate index
	          "coding_rate": "4/6",                  // LoRa coding rate
	          "frequency": "868300000",              // Frequency (Hz)
	        },
	*/
	freq, err := strconv.ParseUint(packetIn.UplinkMessage.Settings.Frequency, 10, 64)
	if err == nil {
		packetOut.Frequency = freq
	}

	if packetIn.UplinkMessage.Settings.DataRate.Lora != nil {
		//log.Println("Is LORA")
		packetOut.Modulation = "LORA"
		packetOut.SpreadingFactor = uint8(packetIn.UplinkMessage.Settings.DataRate.Lora.SpreadingFactor)
		packetOut.Bandwidth = uint64(packetIn.UplinkMessage.Settings.DataRate.Lora.Bandwidth)
	}
	if packetIn.UplinkMessage.Settings.DataRate.Fsk != nil {
		//log.Println("Is FSK")
		packetOut.Modulation = "FSK"
		packetOut.Bitrate = uint64(packetIn.UplinkMessage.Settings.DataRate.Fsk.BitRate)
	}

	packetOut.CodingRate = packetIn.UplinkMessage.Settings.CodingRate

	/*
		V3
		   {
		       "rx_metadata": [
	*/
	for _, gatewayIn := range packetIn.UplinkMessage.RxMetadata {

		/*
			V3
			{                        // A list of metadata for each antenna of each gateway that received this message
					"gateway_ids": {
						"gateway_id": "gtw1",                // Gateway ID
						"eui": "9C5C8E00001A05C4"            // Gateway EUI
					},
					"time": "2020-02-12T15:15:45.787Z",    // ISO 8601 UTC timestamp at which the uplink has been received by the gateway
					"timestamp": 2463457000,               // Timestamp of the gateway concentrator when the message has been received
					fine_timestamp
					encrypted_fine_timestamp
					encrypted_fine_timestamp_key_id
					"rssi": -35,                           // Received signal strength indicator (dBm)
					"channel_rssi": -35,                   // Received signal strength indicator of the channel (dBm)
					signal_rssi
					rssi_standard_deviation
					"snr": 5,                              // Signal-to-noise ratio (dB)
					"uplink_token": "ChIKEA...",           // Uplink token injected by gateway, Gateway Server or fNS
					"channel_index": 2                     // Index of the gateway channel that received the message
					"antenna_index"
					"location":
						latitude
						longitude
						altitude
						accuracy
						source:
							SOURCE_UNKNOWN	0	The source of the location is not known or not set.
							SOURCE_GPS	1	The location is determined by GPS.
							SOURCE_REGISTRY	3	The location is set in and updated from a registry.
							SOURCE_IP_GEOLOCATION	4	The location is estimated with IP geolocation.
							SOURCE_WIFI_RSSI_GEOLOCATION	5	The location is estimated with WiFi RSSI geolocation.
							SOURCE_BT_RSSI_GEOLOCATION	6	The location is estimated with BT/BLE RSSI geolocation.
							SOURCE_LORA_RSSI_GEOLOCATION	7	The location is estimated with LoRa RSSI geolocation.
							SOURCE_LORA_TDOA_GEOLOCATION	8	The location is estimated with LoRa TDOA geolocation.
							SOURCE_COMBINED_GEOLOCATION	9	The location is estimated by a combination of geolocation sources.
			    }
		*/
		gatewayOut := types.TtnMapperGateway{}

		gatewayOut.GatewayId = gatewayIn.GatewayIds.GatewayID
		gatewayOut.GatewayEui = gatewayIn.GatewayIds.Eui

		// If the id is eui-deadbeef, strip the prefix, capitalize and use as EUI
		if strings.HasPrefix(gatewayIn.GatewayIds.GatewayID, "eui-") && len(gatewayIn.GatewayIds.GatewayID) == 20 {
			eui := strings.TrimPrefix(gatewayIn.GatewayIds.GatewayID, "eui-")
			strings.ToUpper(eui)
			gatewayOut.GatewayEui = eui
		}

		gatewayOut.Timestamp = uint32(gatewayIn.Timestamp)

		gatewayTime := time.Time(gatewayIn.Time)
		if !gatewayTime.IsZero() {
			gatewayOut.Time = gatewayTime.UnixNano()
		}

		fineTimestamp, err := strconv.ParseUint(gatewayIn.FineTimestamp, 10, 32)
		if err == nil {
			gatewayOut.FineTimestamp = fineTimestamp
		}

		gatewayOut.FineTimestampEncrypted = gatewayIn.EncryptedFineTimestamp
		gatewayOut.FineTimestampEncryptedKeyId = gatewayIn.EncryptedFineTimestampKeyID

		gatewayOut.ChannelIndex = uint32(gatewayIn.ChannelIndex)

		if gatewayIn.ChannelRssi != 0 {
			gatewayOut.Rssi = gatewayIn.ChannelRssi
		}
		if gatewayIn.Rssi != 0 {
			gatewayOut.Rssi = gatewayIn.Rssi
		}
		if gatewayIn.SignalRssi != 0 {
			gatewayOut.SignalRssi = gatewayIn.SignalRssi
		}
		gatewayOut.Snr = gatewayIn.Snr
		gatewayOut.AntennaIndex = uint8(gatewayIn.AntennaIndex)
		if gatewayIn.Location != nil {
			gatewayOut.Latitude = gatewayIn.Location.Latitude
			gatewayOut.Longitude = gatewayIn.Location.Longitude
			gatewayOut.Altitude = gatewayIn.Location.Altitude
			gatewayOut.LocationAccuracy = gatewayIn.Location.Accuracy

			gatewayOut.LocationSource = string(gatewayIn.Location.Source)
		}

		packetOut.Gateways = append(packetOut.Gateways, gatewayOut)

	}
}
