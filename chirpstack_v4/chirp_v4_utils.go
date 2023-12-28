package chirpstack_v4

import (
	"errors"
	"github.com/chirpstack/chirpstack/api/go/v4/gw"
	"github.com/chirpstack/chirpstack/api/go/v4/integration"
	"log"
	"strconv"
	"strings"
	"time"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"
)

func ParseChirpV4Payload(packetIn integration.UplinkEvent, packetOut *types.TtnMapperUplinkMessage) error {
	if err := utils.ParsePayloadFields(int64(packetOut.FPort), packetIn.Object.AsMap(), packetOut); err != nil {
		return err
	}

	return nil
}

func CopyChirpV4Fields(packetIn integration.UplinkEvent, packetOut *types.TtnMapperUplinkMessage) error {
	packetOut.Time = time.Now().UnixNano()

	packetOut.AppID = packetIn.DeviceInfo.ApplicationName
	packetOut.DevID = packetIn.DeviceInfo.DeviceName
	packetOut.DevEui = strings.ToUpper(packetIn.DeviceInfo.DevEui)

	packetOut.FPort = uint8(packetIn.FPort)
	packetOut.FCnt = int64(packetIn.FCnt)

	if packetIn.TxInfo == nil {
		return errors.New("tx info is not set")
	}
	packetOut.Frequency = utils.SanitizeFrequency(float64(packetIn.TxInfo.Frequency))
	packetOut.Modulation = packetIn.TxInfo.Modulation.String()

	loraInfo := packetIn.TxInfo.Modulation.GetLora()
	if loraInfo != nil {
		packetOut.Modulation = types.MOD_LORA
		packetOut.SpreadingFactor = uint8(loraInfo.SpreadingFactor)
		packetOut.Bandwidth = uint64(loraInfo.Bandwidth * 1000) // kHz to Hz
		packetOut.CodingRate = ChirpCrToMapperCr(loraInfo.CodeRate)
	}
	fskInfo := packetIn.TxInfo.Modulation.GetFsk()
	if fskInfo != nil { // FSK
		packetOut.Modulation = types.MOD_FSK
		packetOut.Bitrate = uint64(fskInfo.Datarate)
	}
	lrFhssInfo := packetIn.TxInfo.Modulation.GetLrFhss()
	if lrFhssInfo != nil {
		packetOut.Modulation = types.MOD_LRFHSS
		packetOut.SpreadingFactor = uint8(lrFhssInfo.GridSteps)
		packetOut.Bandwidth = uint64(lrFhssInfo.OperatingChannelWidth)
		packetOut.CodingRate = ChirpCrToMapperCr(lrFhssInfo.CodeRate)
	}

	if packetIn.RxInfo == nil {
		return errors.New("rx info is not set")
	}
	for _, gatewayIn := range packetIn.RxInfo {
		if gatewayIn == nil {
			continue
		}

		var err error

		gatewayOut := types.TtnMapperGateway{}
		gatewayOut.Attributes = make(map[string]interface{}, 0)

		// Default to ChirpStack provided ID, unless it is overridden by Helium, ThingsIX, etc
		gatewayOut.NetworkId = packetOut.NetworkId

		if gatewayOut.GatewayId == "packetbroker" {
			err = CopyPacketBrokerFields(gatewayIn, &gatewayOut)
		} else if isThingsIx(gatewayIn) {
			err = CopyThingsIxFields(gatewayIn, &gatewayOut)
		} else if isHelium(gatewayIn) {
			err = CopyHeliumFields(gatewayIn, &gatewayOut)
		} else {
			err = CopyChirpGwFields(gatewayIn, &gatewayOut)
		}

		if err != nil {
			log.Println(err.Error())
			continue
		}

		packetOut.Gateways = append(packetOut.Gateways, gatewayOut)
	}

	return nil
}

func CopyPacketBrokerFields(gatewayIn *gw.UplinkRxInfo, gatewayOut *types.TtnMapperGateway) error {
	// If the gateway id is packet broker, ignore because we do not know if any of the useful metadata comes along
	return errors.New("not implemented")
}

func CopyHeliumFields(gatewayIn *gw.UplinkRxInfo, gatewayOut *types.TtnMapperGateway) error {
	gatewayOut.NetworkId = types.NS_HELIUM + "://000024"
	gatewayOut.GatewayId = gatewayIn.Metadata["gateway_id"]
	//gatewayOut.GatewayEui =
	gatewayOut.AntennaIndex = uint8(gatewayIn.Antenna)

	// gateway Time is the wall clock time
	if gatewayIn.GwTime != nil {
		gatewayOut.Time = time.Unix(gatewayIn.GwTime.Seconds, int64(gatewayIn.GwTime.Nanos)).UnixNano()
	} else if gatewayIn.NsTime != nil {
		// Fall back to server time
		gatewayOut.Time = time.Unix(gatewayIn.NsTime.Seconds, int64(gatewayIn.NsTime.Nanos)).UnixNano()
	}

	// gateway Timestamp is the internal clock counter of the concentrator
	// not provided by ChirpStack/ThingsIX

	// Fine timestamp - not encrypted?
	if gatewayIn.GetFineTimeSinceGpsEpoch() != nil {
		gatewayOut.FineTimestamp = uint64(time.Unix(gatewayIn.GetFineTimeSinceGpsEpoch().Seconds, int64(gatewayIn.GetFineTimeSinceGpsEpoch().Nanos)).UnixNano())
	}

	gatewayOut.ChannelIndex = gatewayIn.Channel
	gatewayOut.Rssi = float32(gatewayIn.Rssi)
	// Missing Channel/Signal RSSI
	gatewayOut.Snr = gatewayIn.Snr

	if val, ok := gatewayIn.Metadata["gateway_lat"]; ok {
		latitude, err := strconv.ParseFloat(val, 64)
		if err == nil {
			gatewayOut.Latitude = latitude
		}
	}

	if val, ok := gatewayIn.Metadata["gateway_long"]; ok {
		longitude, err := strconv.ParseFloat(val, 64)
		if err == nil {
			gatewayOut.Longitude = longitude
		}
	}

	//if val, ok := gatewayIn.Metadata[""]; ok {
	//	altitude, err := strconv.ParseFloat(val, 64)
	//	if err == nil {
	//		gatewayOut.Altitude = int32(altitude)
	//	}
	//}

	//gatewayOut.LocationAccuracy =
	gatewayOut.LocationSource = "HELIUM"

	gatewayOut.Name = gatewayIn.Metadata["gateway_name"]

	// Copy all metadata fields into attributes
	for k, v := range gatewayIn.Metadata {
		gatewayOut.Attributes[k] = v
	}

	return nil
}

func CopyThingsIxFields(gatewayIn *gw.UplinkRxInfo, gatewayOut *types.TtnMapperGateway) error {
	gatewayOut.NetworkId = types.NS_THINGSIX
	gatewayOut.GatewayId = gatewayIn.Metadata["thingsix_gateway_id"]
	//gatewayOut.GatewayEui =
	gatewayOut.AntennaIndex = uint8(gatewayIn.Antenna)

	// gateway Time is the wall clock time
	if gatewayIn.GwTime != nil {
		gatewayOut.Time = time.Unix(gatewayIn.GwTime.Seconds, int64(gatewayIn.GwTime.Nanos)).UnixNano()
	} else if gatewayIn.NsTime != nil {
		// Fall back to server time
		gatewayOut.Time = time.Unix(gatewayIn.NsTime.Seconds, int64(gatewayIn.NsTime.Nanos)).UnixNano()
	}

	// gateway Timestamp is the internal clock counter of the concentrator
	// not provided by ChirpStack/ThingsIX

	// Fine timestamp - not encrypted?
	if gatewayIn.GetFineTimeSinceGpsEpoch() != nil {
		gatewayOut.FineTimestamp = uint64(time.Unix(gatewayIn.GetFineTimeSinceGpsEpoch().Seconds, int64(gatewayIn.GetFineTimeSinceGpsEpoch().Nanos)).UnixNano())
	}

	gatewayOut.ChannelIndex = gatewayIn.Channel
	gatewayOut.Rssi = float32(gatewayIn.Rssi)
	// Missing Channel/Signal RSSI
	gatewayOut.Snr = gatewayIn.Snr

	if val, ok := gatewayIn.Metadata["thingsix_location_latitude"]; ok {
		latitude, err := strconv.ParseFloat(val, 64)
		if err == nil {
			gatewayOut.Latitude = latitude
		}
	}

	if val, ok := gatewayIn.Metadata["thingsix_location_longitude"]; ok {
		longitude, err := strconv.ParseFloat(val, 64)
		if err == nil {
			gatewayOut.Longitude = longitude
		}
	}

	if val, ok := gatewayIn.Metadata["thingsix_altitude"]; ok {
		altitude, err := strconv.ParseFloat(val, 64)
		if err == nil {
			gatewayOut.Altitude = int32(altitude)
		}
	}

	//gatewayOut.LocationAccuracy =
	gatewayOut.LocationSource = "THINGSIX"

	//gatewayOut.Name =

	// Copy all metadata fields into attributes
	for k, v := range gatewayIn.Metadata {
		gatewayOut.Attributes[k] = v
	}

	return nil
}

func CopyChirpGwFields(gatewayIn *gw.UplinkRxInfo, gatewayOut *types.TtnMapperGateway) error {
	// V4 gateway ID is always the EUI
	gatewayOut.GatewayId = "eui-" + strings.ToLower(gatewayIn.GatewayId)
	gatewayOut.GatewayEui = strings.ToUpper(gatewayIn.GatewayId)
	// gatewayOut.Description = ... // TODO: Get Gateway Name

	// gateway Time is the wall clock time
	if gatewayIn.GwTime != nil {
		gatewayOut.Time = time.Unix(gatewayIn.GwTime.Seconds, int64(gatewayIn.GwTime.Nanos)).UnixNano()
	} else if gatewayIn.NsTime != nil {
		// Fall back to server time
		gatewayOut.Time = time.Unix(gatewayIn.NsTime.Seconds, int64(gatewayIn.NsTime.Nanos)).UnixNano()
	}

	// gateway Timestamp is the internal clock counter of the concentrator
	// not provided by ChirpStack

	// Fine timestamp - not encrypted?
	if gatewayIn.GetFineTimeSinceGpsEpoch() != nil {
		gatewayOut.FineTimestamp = uint64(time.Unix(gatewayIn.GetFineTimeSinceGpsEpoch().Seconds, int64(gatewayIn.GetFineTimeSinceGpsEpoch().Nanos)).UnixNano())
	}

	gatewayOut.AntennaIndex = uint8(gatewayIn.Antenna)
	gatewayOut.ChannelIndex = gatewayIn.Channel
	gatewayOut.Rssi = float32(gatewayIn.Rssi)
	// Missing Channel/Signal RSSI
	gatewayOut.Snr = gatewayIn.Snr

	if gatewayIn.Location != nil {
		gatewayOut.Latitude = gatewayIn.Location.Latitude
		gatewayOut.Longitude = gatewayIn.Location.Longitude
		gatewayOut.Altitude = int32(gatewayIn.Location.Altitude)
		gatewayOut.LocationAccuracy = int32(gatewayIn.Location.Accuracy)
		gatewayOut.LocationSource = gatewayIn.Location.Source.String()
	}

	// Copy all metadata fields into attributes
	for k, v := range gatewayIn.Metadata {
		gatewayOut.Attributes[k] = v
	}

	return nil
}

// When PRs are accepted, do a simple if(network == helium) or if(network == thingsix)
// https://github.com/ThingsIXFoundation/packet-handling/pull/75
// https://github.com/helium/helium-packet-router/pull/286/files

// For now we need to guess by looking at the other fields in the metadata
func isThingsIx(gatewayIn *gw.UplinkRxInfo) bool {
	// ThingsIX
	// https://github.com/ThingsIXFoundation/packet-handling/blob/0245669dd103a88c366aaa859c0b159a7517d2c3/forwarder/metadata.go#L32
	if _, ok := gatewayIn.Metadata["thingsix_gateway_id"]; ok {
		return true
	}
	return false
}

func isHelium(gatewayIn *gw.UplinkRxInfo) bool {
	// Helium
	// https://github.com/helium/helium-packet-router/blob/b06b2758f443c04079e74daf809a156f281162bd/src/protocols/gwmp/hpr_gwmp_worker.erl#L240
	// https://github.com/helium/helium-packet-router/blob/b06b2758f443c04079e74daf809a156f281162bd/src/protocols/gwmp/hpr_gwmp_worker.erl#L252
	_, hasId := gatewayIn.Metadata["gateway_id"]
	_, hasName := gatewayIn.Metadata["gateway_name"]
	_, hasRegi := gatewayIn.Metadata["regi"]
	// Also look at optional location fields, as a gateway without location isn't useful
	_, hasH3 := gatewayIn.Metadata["gateway_h3index"]
	_, hasLat := gatewayIn.Metadata["gateway_lat"]
	_, hasLon := gatewayIn.Metadata["gateway_long"]

	if hasId && hasName && hasRegi && hasH3 && hasLat && hasLon {
		return true
	}
	return false
}

func ChirpCrToMapperCr(cr gw.CodeRate) string {
	/*
		0:  "CR_UNDEFINED",
		1:  "CR_4_5",
		2:  "CR_4_6",
		3:  "CR_4_7",
		4:  "CR_4_8",
		5:  "CR_3_8",
		6:  "CR_2_6",
		7:  "CR_1_4",
		8:  "CR_1_6",
		9:  "CR_5_6",
		10: "CR_LI_4_5",
		11: "CR_LI_4_6",
		12: "CR_LI_4_8",
	*/

	// Translate to previous representation
	if cr == gw.CodeRate_CR_UNDEFINED {
		return "OFF"
	}
	if cr == gw.CodeRate_CR_4_5 {
		return "4/5"
	}
	if cr == gw.CodeRate_CR_4_6 {
		return "4/6"
	}
	if cr == gw.CodeRate_CR_4_7 {
		return "4/7"
	}
	if cr == gw.CodeRate_CR_4_8 {
		return "4/8"
	}

	// Otherwise return ChirpStack's representation
	return cr.String()
}
