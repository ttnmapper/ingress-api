package helium

type UplinkMessage struct {
	AppEui  string `json:"app_eui"`
	DevEui  string `json:"dev_eui"`
	DevAddr string `json:"devaddr"`
	Fcnt    int64

	Decoded struct {
		Payload interface{}
	}

	Hotspots []Hotspot

	Id          string
	Metadata    interface{}
	Name        string
	Payload     string
	PayloadSize int64 `json:"payload_size"`
	Port        uint8
	ReportedAt  int64 `json:"reported_at"`
}

type Hotspot struct {
	Channel    uint32
	Frequency  float64
	Id         string
	Name       string
	ReportedAt int64 `json:"reported_at"`
	Rssi       float32
	Snr        float32
	Spreading  string

	Latitude  interface{} `json:"lat"`
	Longitude interface{} `json:"long"`
}
