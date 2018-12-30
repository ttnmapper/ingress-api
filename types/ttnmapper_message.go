package types

type TtnMapperUplinkMessage struct {
	UplinkMessage
	TtnMLatitude   float64 `json:"ttnmapper_latitude,omitempty"`
	TtnMLongitude  float64 `json:"ttnmapper_longitude,omitempty"`
	TtnMAltitude   float64 `json:"ttnmapper_altitude,omitempty"`
	TtnMAccuracy   float64 `json:"ttnmapper_accuracy,omitempty"`
	TtnMSatellites int32   `json:"ttnmapper_satellites,omitempty"`
	TtnMHdop       float64 `json:"ttnmapper_hdop,omitempty"`

	// The source can be: gps, config, registry, ip_geolocation or unknown (unknown may be left out)
	// See proto definition for more info
	TtnMProvider   string `json:"ttnmapper_provider,omitempty"`
	TtnMExperiment string `json:"ttnmapper_experiment,omitempty"`
	TtnMUserId     string `json:"ttnmapper_userid,omitempty"`
	TtnMUserAgent  string `json:"ttnmapper_useragent,omitempty"`
}
