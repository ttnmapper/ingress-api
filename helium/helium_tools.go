package helium

import (
	"strconv"
	"strings"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"
)

func CopyHeliumFields(packetIn UplinkMessage, packetOut *types.TtnMapperUplinkMessage) {

	packetOut.NetworkType = types.NS_HELIUM
	/*
		A hostname or IP address to uniquely identify the network server
	*/
	packetOut.NetworkAddress = "000024"

	// Combine network type and network address into a single networkid field which is globally unique.
	// We will start using a combination of the LoRaWAN NetID and a TenantID soon.
	packetOut.NetworkId = types.NS_HELIUM + "://000024"

	packetOut.AppID = packetIn.AppEui
	packetOut.DevID = packetIn.Name
	packetOut.DevEui = packetIn.DevEui

	packetOut.Time = packetIn.ReportedAt * 1000000 // ms to ns

	packetOut.FPort = packetIn.Port
	packetOut.FCnt = packetIn.Fcnt

	packetOut.Gateways = make([]types.TtnMapperGateway, 0)

	for _, hotspot := range packetIn.Hotspots {
		packetOut.Frequency = utils.SanitizeFrequency(hotspot.Frequency * 1000000) // MHz to Hz
		packetOut.Modulation = "LORA"

		drParts := strings.Split(hotspot.Spreading, "BW")
		bandwidth, _ := strconv.Atoi(drParts[1])
		packetOut.Bandwidth = uint64(bandwidth * 1000) // kHz to Hz
		sf, _ := strconv.Atoi(strings.TrimPrefix(drParts[0], "SF"))
		packetOut.SpreadingFactor = uint8(sf)

		packetOut.Bitrate = 0
		packetOut.CodingRate = ""

		gateway := types.TtnMapperGateway{
			NetworkId:                   packetOut.NetworkId,
			GatewayId:                   hotspot.Name,
			GatewayEui:                  "",
			AntennaIndex:                0,
			Time:                        hotspot.ReportedAt * 1000000, // ms to ns
			Timestamp:                   0,
			FineTimestamp:               0,
			FineTimestampEncrypted:      nil,
			FineTimestampEncryptedKeyId: "",
			ChannelIndex:                hotspot.Channel,
			Rssi:                        hotspot.Rssi,
			SignalRssi:                  0,
			Snr:                         hotspot.Snr,
			Latitude:                    hotspot.Latitude,
			Longitude:                   hotspot.Longitude,
			Altitude:                    0,
			LocationAccuracy:            0,
			LocationSource:              "",
			Description:                 "",
		}

		packetOut.Gateways = append(packetOut.Gateways, gateway)
	}
}
