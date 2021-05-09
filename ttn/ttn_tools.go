package ttn

import (
	"math"
	"strconv"
	"strings"
	"time"
	types2 "ttnmapper-ingress-api/ttn/ttn_types"
	"ttnmapper-ingress-api/types"
)

func CopyTtnV2Fields(packetIn types2.UplinkMessage, packetOut *types.TtnMapperUplinkMessage) {
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

		// TODO use network reported by packetbroker
		gatewayOut.NetworkId = packetOut.NetworkType + "://" + packetOut.NetworkAddress
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
