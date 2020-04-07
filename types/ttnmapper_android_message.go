package types

import "ttnmapper-ingress-api/ttnV2"

type TtnMapperAndroidMessage struct {
	ttnV2.UplinkMessage

	//@field:Json(name = "phone_lat")
	//var phoneLat: Double?, // -34.0484124
	PhoneLat float64 `json:"phone_lat,omitempty"`

	//@field:Json(name = "phone_lon")
	//var phoneLon: Double?, // 18.8214014
	PhoneLon float64 `json:"phone_lon,omitempty"`

	//@field:Json(name = "phone_alt")
	//var phoneAlt: Double?, // 184.9958825503345
	PhoneAlt float64 `json:"phone_alt,omitempty"`

	//@field:Json(name = "phone_loc_acc")
	//var phoneLocAcc: Double?, // 10
	PhoneLocAccuracy float64 `json:"phone_loc_acc,omitempty"`

	//@field:Json(name = "phone_loc_provider")
	//var phoneLocProvider: String?, // fused
	PhoneLocProvider string `json:"phone_loc_provider,omitempty"`

	//@field:Json(name = "phone_time")
	//var phoneTime: String?, // 2018-03-18T10:05:44Z
	PhoneTime string `json:"phone_time,omitempty"`

	//@field:Json(name = "user_agent")
	//var userAgent: String?, // Android7.0 App30:2018.03.04
	UserAgent string `json:"user_agent,omitempty"`

	//@field:Json(name = "iid")
	//var iid: String?, //some random number
	Iid string `json:"iid,omitempty"`

	//@field:Json(name = "experiment")
	//var experiment: String? // experiment name
	Experiment string `json:"experiment,omitempty"`

	//MqttTopic string `json:"mqtt_topic,omitempty"`
}
