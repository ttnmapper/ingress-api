package types

type TtnMapperUplinkMessage struct {
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
	GatewayId                   string  `json:"gtw_id"`
	GatewayEui                  string  `json:"gtw_eui,omitempty"`
	AntennaIndex                uint8   `json:"antenna_index"`
	Time                        int64   `json:"time,omitempty"`
	Timestamp                   uint32  `json:"timestamp,omitempty"`
	FineTimestamp               uint64  `json:"fine_timestamp,omitempty"`
	FineTimestampEncrypted      []byte  `json:"fine_timestamp_encrypted,omitempty"`
	FineTimestampEncryptedKeyId string  `json:"encrypted_fine_timestamp_key_id,omitempty"`
	ChannelIndex                uint32  `json:"channel,omitempty"`
	Rssi                        float32 `json:"rssi,omitempty"` // same as channel rssi
	SignalRssi                  float32 `json:"signal_rssi,omitempty"`
	Snr                         float32 `json:"snr,omitempty"`

	Latitude         float64 `json:"latitude,omitempty"`
	Longitude        float64 `json:"longitude,omitempty"`
	Altitude         int32   `json:"altitude,omitempty"`
	LocationAccuracy int32   `json:"location_accuracy,omitempty"`

	// The source can be: gps, config, registry, ip_geolocation or unknown (unknown may be left out)
	// See proto definition for more info
	LocationSource string `json:"location_source,omitempty"`
}
