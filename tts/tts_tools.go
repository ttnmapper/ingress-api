package tts

import (
	"encoding/json"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"strings"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"
)

func DecodeV3Payload(port int64, packetIn ttnpb.ApplicationUp, packetOut *types.TtnMapperUplinkMessage) error {
	// DecodedPayload is &Struct{Fields:map[string]*Value{},XXX_unrecognized:[],}.
	// Convert to a more standard map[string]interface{}

	// Marshal struct to json
	marshaler := jsonpb.TTN()
	decodedJson, err := marshaler.Marshal(packetIn.GetUplinkMessage().DecodedPayload)
	if err != nil {
		return err
	}

	// Unmarshal json to interface{}
	var decodedInterface interface{}
	err = json.Unmarshal(decodedJson, &decodedInterface)
	if err != nil {
		return err
	}

	// Parse fields from interface
	err = utils.ParsePayloadFields(port, decodedInterface, packetOut)
	if err != nil {
		return err
	}

	return nil
}

func CopyV3Fields(packetIn ttnpb.ApplicationUp, packetOut *types.TtnMapperUplinkMessage) {
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
	packetOut.AppID = packetIn.EndDeviceIdentifiers.ApplicationId
	packetOut.DevID = packetIn.EndDeviceIdentifiers.DeviceId

	if packetIn.EndDeviceIdentifiers.DevEui != nil {
		packetOut.DevEui = packetIn.EndDeviceIdentifiers.DevEui.String()
	}

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
	packetOut.Time = packetIn.ReceivedAt.UnixNano()

	/*
	   V3
	         "uplink_message":{
	            "f_port":1,
	            "f_cnt":527,
	*/
	packetOut.FPort = uint8(packetIn.GetUplinkMessage().FPort)
	packetOut.FCnt = int64(packetIn.GetUplinkMessage().FCnt)

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
	packetOut.Frequency = utils.SanitizeFrequency(float64(packetIn.GetUplinkMessage().Settings.Frequency))

	if packetIn.GetUplinkMessage().Settings.DataRate.GetLora() != nil {
		//log.Println("Is LORA")
		packetOut.Modulation = "LORA"
		packetOut.SpreadingFactor = uint8(packetIn.GetUplinkMessage().Settings.DataRate.GetLora().SpreadingFactor)
		packetOut.Bandwidth = uint64(packetIn.GetUplinkMessage().Settings.DataRate.GetLora().Bandwidth)
	}
	if packetIn.GetUplinkMessage().Settings.DataRate.GetFsk() != nil {
		//log.Println("Is FSK")
		packetOut.Modulation = "FSK"
		packetOut.Bitrate = uint64(packetIn.GetUplinkMessage().Settings.DataRate.GetFsk().BitRate)
	}
	if packetIn.GetUplinkMessage().Settings.DataRate.GetLrfhss() != nil {
		packetOut.Modulation = "LR_FHSS"
		packetOut.Bandwidth = uint64(packetIn.GetUplinkMessage().Settings.DataRate.GetLrfhss().GetOperatingChannelWidth())
		// TODO: grid steps, code rate
	}

	packetOut.CodingRate = packetIn.GetUplinkMessage().Settings.CodingRate

	/*
		V3
		   {
		       "rx_metadata": [
	*/
	for _, gatewayIn := range packetIn.GetUplinkMessage().RxMetadata {

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

		// The gateway's ID - unique per network
		gatewayOut.GatewayId = gatewayIn.GatewayIdentifiers.GetGatewayId()
		if gatewayIn.GatewayIdentifiers.Eui != nil {
			gatewayOut.GatewayEui = gatewayIn.GatewayIdentifiers.Eui.String()
		}

		/*
			Determine the network
			X-Tts-Domain suffix .cloud.thethings.network is The Things Stack Community Edition (The Things Network)
			X-Tts-Domain suffix .cloud.thethings.industries is The Things Stack Cloud (The Things Industries)
			home_network_net_id 000013 and home_network_tenant_id ttn is The Things Network
			home_network_net_id 000013 with any other home_network_tenant_id can be anything: The Things Stack Cloud, The Things Stack Enterprise, The Things Stack Open Source, even ChirpStack using Passive Roaming and using address space of TTN
		*/
		/*
		   "gateway_ids":{
		      "gateway_id":"packetbroker"
		   },
		   "packet_broker":{
		      "message_id":"01EZF0NS39X53JDXJGF8AC2R1Z",
		      "forwarder_net_id":"000013",
		      "forwarder_tenant_id":"ttn",
		      "forwarder_cluster_id":"ttn-v2-eu-3",
		      "home_network_net_id":"000013",
		      "home_network_tenant_id":"ttn",
		      "home_network_cluster_id":"ttn-eu1",
		*/
		if gatewayIn.PacketBroker != nil {
			/*
				ttsDomain = forwrder_tenant_id@forwarder_net_id // "ttn@000013"
			*/
			forwarderTenantId := gatewayIn.PacketBroker.ForwarderTenantId
			forwarderNetId := gatewayIn.PacketBroker.ForwarderNetId
			if forwarderTenantId == "ttnv2" {
				gatewayOut.NetworkId = "thethingsnetwork.org"
			} else {
				gatewayOut.NetworkId = types.NS_TTS_V3 + "://" + forwarderTenantId + "@" + forwarderNetId.String()
			}

			/*
				Use GatewayId and EUI if reported by PacketBroker
			*/
			if gatewayIn.PacketBroker.ForwarderGatewayEui != nil {
				gatewayOut.GatewayEui = gatewayIn.PacketBroker.ForwarderGatewayEui.String()
			}
			if gatewayIn.PacketBroker.ForwarderGatewayId != nil {
				gatewayOut.GatewayId = gatewayIn.PacketBroker.ForwarderGatewayId.Value
			}

		} else {
			gatewayOut.NetworkId = packetOut.NetworkId
		}

		// If the GatewayId is eui-deadbeef, strip the prefix, capitalize and use as EUI
		if strings.HasPrefix(gatewayOut.GatewayId, "eui-") && len(gatewayOut.GatewayId) == 20 {
			eui := strings.TrimPrefix(gatewayOut.GatewayId, "eui-")
			eui = strings.ToUpper(eui)
			gatewayOut.GatewayEui = eui
		}

		gatewayOut.Timestamp = gatewayIn.Timestamp

		if gatewayIn.Time != nil {
			gatewayOut.Time = gatewayIn.Time.UnixNano()
		}

		gatewayOut.FineTimestamp = gatewayIn.FineTimestamp
		gatewayOut.FineTimestampEncrypted = gatewayIn.EncryptedFineTimestamp
		gatewayOut.FineTimestampEncryptedKeyId = gatewayIn.EncryptedFineTimestampKeyId

		gatewayOut.ChannelIndex = gatewayIn.ChannelIndex

		if gatewayIn.ChannelRssi != 0 {
			gatewayOut.Rssi = gatewayIn.ChannelRssi
		}
		if gatewayIn.Rssi != 0 {
			gatewayOut.Rssi = gatewayIn.Rssi
		}
		if gatewayIn.SignalRssi != nil {
			gatewayOut.SignalRssi = gatewayIn.SignalRssi.Value
		}
		gatewayOut.Snr = gatewayIn.Snr
		gatewayOut.AntennaIndex = uint8(gatewayIn.AntennaIndex)
		if gatewayIn.Location != nil {
			gatewayOut.Latitude = gatewayIn.Location.Latitude
			gatewayOut.Longitude = gatewayIn.Location.Longitude
			gatewayOut.Altitude = gatewayIn.Location.Altitude
			gatewayOut.LocationAccuracy = gatewayIn.Location.Accuracy

			gatewayOut.LocationSource = gatewayIn.Location.Source.String()
		}

		packetOut.Gateways = append(packetOut.Gateways, gatewayOut)

	}
}
