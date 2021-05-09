package types

const (
	NS_TTN_V2 = "NS_TTN_V2"
	NS_TTS_V3 = "NS_TTS_V3"
	NS_CHIRP  = "NS_CHIRP"
)

type TtnMapperUplinkMessage struct {
	/*
		One of the constant network type strings from the above const()
	*/
	NetworkType string `json:"network_type,omitempty"`
	/*
		A hostname or IP address to uniquely identify the network server
	*/
	NetworkAddress string `json:"network_address,omitempty"`

	// Combine network type and network address into a single networkid field which is globally unique.
	// We will start using a combination of the LoRaWAN NetID and a TenantID soon.
	NetworkId string `json:"network_id,omitempty"`

	AppID  string `json:"app_id"`
	DevID  string `json:"dev_id"`
	DevEui string `json:"dev_eui,omitempty"`

	Time int64 `json:"time,omitempty,omitempty"`

	FPort uint8 `json:"port,omitempty"`
	FCnt  int64 `json:"counter,omitempty"`

	Frequency       uint64 `json:"frequency,omitempty"`
	Modulation      string `json:"modulation,omitempty"`
	Bandwidth       uint64 `json:"bandwidth,omitempty"`
	SpreadingFactor uint8  `json:"spreading_factor,omitempty"`
	Bitrate         uint64 `json:"bit_rate,omitempty"`
	CodingRate      string `json:"coding_rate,omitempty"`

	Gateways []TtnMapperGateway `json:"gateways"`

	Latitude       float64 `json:"latitude,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	Altitude       float64 `json:"altitude,omitempty"`
	AccuracyMeters float64 `json:"accuracy_meters,omitempty"`
	Satellites     int32   `json:"satellites,omitempty"`
	Hdop           float64 `json:"hdop,omitempty"`
	AccuracySource string  `json:"accuracy_source,omitempty"`

	Experiment string `json:"experiment,omitempty"`
	UserId     string `json:"userid,omitempty"`
	UserAgent  string `json:"useragent,omitempty"`
}

type TtnMapperGateway struct {
	// Globally unique identifier for the specific network instance.
	// Normally `packetOut.NetworkType + "://" + packetOut.NetworkAddress` unless the data is forwarded via the packet broker (ie peering/roaming)
	// TODO: See comment in TtnMapperUplinkMessage
	NetworkId string `json:"network_id,omitempty"`
	// Unique ID for this gateway for the respective network server
	// Use `"eui-" + strings.ToLower(gatewayEui)` if not available
	GatewayId string `json:"gtw_id"`
	// Globally unique identifier for the gateway. EUI64 as an upper case hex string, with 0 prefixes, thus always 16 characters long.
	GatewayEui string `json:"gtw_eui,omitempty"`
	// Some gateways have more than one concentrator and more than one antenna. They could be on different frequency plans.
	AntennaIndex uint8 `json:"antenna_index"`
	// Time info
	// see https://github.com/Lora-net/packet_forwarder/blob/master/PROTOCOL.TXT section 4
	// Wall clock time
	Time int64 `json:"time,omitempty"`
	// Gateway concentrator internal clock counter value
	Timestamp uint32 `json:"timestamp,omitempty"`
	// Fine timestamp if not encrypted
	FineTimestamp uint64 `json:"fine_timestamp,omitempty"`
	// Fine timestamp if encrypted
	FineTimestampEncrypted []byte `json:"fine_timestamp_encrypted,omitempty"`
	// Fine timestamp AES key ID if encrypted
	FineTimestampEncryptedKeyId string `json:"encrypted_fine_timestamp_key_id,omitempty"`
	// Gateway concentrator channel index
	ChannelIndex uint32 `json:"channel,omitempty"`
	// RSSI / Channel RSSI
	Rssi float32 `json:"rssi,omitempty"`
	// Only new gateway architectures provide Signal RSSI values.
	SignalRssi float32 `json:"signal_rssi,omitempty"`
	Snr        float32 `json:"snr,omitempty"`

	// Location of the gateway. Actually it should be antenna location, but we assume they are very close together.
	Latitude         float64 `json:"latitude,omitempty"`
	Longitude        float64 `json:"longitude,omitempty"`
	Altitude         int32   `json:"altitude,omitempty"`
	LocationAccuracy int32   `json:"location_accuracy,omitempty"`

	// The source can be: gps, config, registry, ip_geolocation or unknown (unknown may be left out)
	// See proto definition for more info
	LocationSource string `json:"location_source,omitempty"`

	// Some sources of statuses provide the description
	Description string `json:"description,omitempty"`
}
