package types

type TtnMapperAndroidV2Message struct {
	/*
		{
		   "iid":"DD3D4575-1604-49EA-9C9F-AD8D30ECCD02",
		   "appeui":"wistrio-lora-rak5205-app-01",
		   "lon":5.3813155371357402,
		   "provider":"ios",
		   "mqtt_topic":"wistrio-lora-rak5205-app-01\/devices\/wistrio-lora-rak5205-01\/up",
		   "nodeaddr":"wistrio-lora-rak5205-01",
		   "time":"2019-01-02T21:35:01.268100556Z",
		   "datarate":"SF7BW125",
		   "lat":60.345758433352437,
		   "rssi":-115,
		   "freq":867.5,
		   "accuracy":12.001812410521316,
		   "user_agent":"iOS 12.1.2 - org.ttnmapper.ios.TTNMapper:26",
		   "alt":181.68609857229174,
		   "gwaddr":"mjs-bergen-gateway-4",
		   "snr":-3.25
		}
	*/
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
	Altitude  float64 `json:"alt,omitempty"`
	Accuracy  float64 `json:"accuracy,omitempty"`

	Iid       string `json:"iid,omitempty"`
	Provider  string `json:"provider,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`

	AppId     string `json:"appeui,omitempty"`
	DevId     string `json:"nodeaddr,omitempty"`
	MqttTopic string `json:"mqtt_topic,omitempty"`

	Time     JSONTime `json:"time,omitempty"`
	Datarate string   `json:"datarate,omitempty"`
	Freq     float32  `json:"freq,omitempty"`

	GwId string  `json:"gwaddr,omitempty"`
	Rssi float32 `json:"rssi,omitempty"`
	Snr  float32 `json:"snr,omitempty"`

	Experiment string `json:"experiment,omitempty"`
}
