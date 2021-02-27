package tests

import (
	"encoding/json"
	"log"
	"testing"
	"ttnmapper-ingress-api/ttsV3/models"
)

func TestDecodeV3(t *testing.T) {
	postbody := "{\"end_device_ids\":{\"device_id\":\"eui-0004a30b001c684f\",\"application_ids\":{\"application_id\":\"jpm-test\"},\"dev_eui\":\"0004A30B001C684F\",\"join_eui\":\"0000000000000000\",\"dev_addr\":\"260B652C\"},\"correlation_ids\":[\"as:up:01EX973RKE79GPAM85E5JS4JFX\",\"ns:uplink:01EX973RCQC3G8FMCC5E2DV9AT\",\"pba:conn:up:01EWN8R418E173ZP0AJ57741C7\",\"pba:uplink:01EX973RBVHZDN440HJV0RNGPM\",\"rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01EX973RCQX0FMDC75HYGHYEEH\",\"rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01EX973RKDF6F652NABJ1QB4AK\"],\"received_at\":\"2021-01-30T08:38:51.246859741Z\",\"uplink_message\":{\"session_key_id\":\"AXdOCkz50LYMrJIii1dEvw==\",\"f_port\":1,\"f_cnt\":4609,\"frm_payload\":\"AA==\",\"decoded_payload\":{\"ledState\":\"off\"},\"rx_metadata\":[{\"gateway_ids\":{\"gateway_id\":\"packetbroker\"},\"packet_broker\":{\"message_id\":\"01EX973RBVHZDN440HJV0RNGPM\",\"forwarder_net_id\":\"000013\",\"forwarder_tenant_id\":\"ttn\",\"forwarder_cluster_id\":\"ttn-v2-eu-4\",\"home_network_net_id\":\"000013\",\"home_network_tenant_id\":\"ttn\",\"home_network_cluster_id\":\"ttn-eu1\",\"hops\":[{\"received_at\":\"2021-01-30T08:38:51.003841504Z\",\"sender_address\":\"52.169.150.138\",\"receiver_name\":\"router-dataplane-56fc9fb8fd-8nrzc\",\"receiver_agent\":\"pbdataplane/1.2.0-rc.1 go/1.15.6 linux/amd64\"},{\"received_at\":\"2021-01-30T08:38:51.006027538Z\",\"sender_name\":\"router-dataplane-56fc9fb8fd-8nrzc\",\"sender_address\":\"kafkapb://router?topic=forwarder_uplink\",\"receiver_name\":\"router-56fb7bcf99-fg4tp\",\"receiver_agent\":\"pbrouter/1.2.0-rc.1 go/1.15.6 linux/amd64\"},{\"received_at\":\"2021-01-30T08:38:51.009209494Z\",\"sender_name\":\"router-56fb7bcf99-fg4tp\",\"sender_address\":\"kafkapb://ttn-eu1?topic=deliver_000013.ttn.ttn-eu1_uplink\",\"receiver_name\":\"router-dataplane-56fc9fb8fd-d9wh4\",\"receiver_agent\":\"pbdataplane/1.2.0-rc.1 go/1.15.6 linux/amd64\"}]},\"rssi\":-103,\"channel_rssi\":-103,\"snr\":6.8,\"uplink_token\":\"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU1hkUmFteDZWVmMxTTFKVmNESlZSVGx2V2tWWmVFbHBkMmxrUjBadVNXcHZhVll5TlZSaVdGWXlZV3MxTlZWVVVqVmxSRVp6WlZaS1ZXSkhhRFZWVTBvNUxrdHRhbWhIT1RGU1YydE1PQzB5WVd4TGNYYzRlbmN1WVVNd1dHNWxhVmhQYTFkelRYWjBWQzVTUWpoSVZpMUlOVUp2Wm01a1lXZDJXVTl4V2xaVVlVNDRORVZrWjNsTk5GbzVhblJpVVZwclRVWndkSGxDYm5GbmRFWkpOMVZhV2pORWNIaHlTbTEzUkMxSU16SnZhSGd3WmxwSmFqaHdPSGhNTlhsMU0wSmtSamhOZW1kcldsWkdTalZUYUhOWWMzQlNVeTFGUzJKSE5FRmhjM1p0Wm5KUVlrWnJiMFpXV0hNeFRsbFpVMDV4UjBkd2FIb3pZa1pyU0ZOTmFFOVJMV0p5ZG1WR1JDMXBVRFJ5ZVc4M1Z6TkVaRmRxTG5SMlNsTXpjSFJJYURReE1VSlllbXhhUTNsMFdXYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS00In19\"},{\"gateway_ids\":{\"gateway_id\":\"packetbroker\"},\"packet_broker\":{\"message_id\":\"01EX973RC87Y6CKNW536GY8G0Q\",\"forwarder_net_id\":\"000013\",\"forwarder_tenant_id\":\"ttn\",\"forwarder_cluster_id\":\"ttn-v2-eu-3\",\"home_network_net_id\":\"000013\",\"home_network_tenant_id\":\"ttn\",\"home_network_cluster_id\":\"ttn-eu1\",\"hops\":[{\"received_at\":\"2021-01-30T08:38:51.016342101Z\",\"sender_address\":\"40.113.68.198\",\"receiver_name\":\"router-dataplane-56fc9fb8fd-8nrzc\",\"receiver_agent\":\"pbdataplane/1.2.0-rc.1 go/1.15.6 linux/amd64\"},{\"received_at\":\"2021-01-30T08:38:51.023747493Z\",\"sender_name\":\"router-dataplane-56fc9fb8fd-8nrzc\",\"sender_address\":\"kafkapb://router?topic=forwarder_uplink\",\"receiver_name\":\"router-56fb7bcf99-fg4tp\",\"receiver_agent\":\"pbrouter/1.2.0-rc.1 go/1.15.6 linux/amd64\"},{\"received_at\":\"2021-01-30T08:38:51.029807445Z\",\"sender_name\":\"router-56fb7bcf99-fg4tp\",\"sender_address\":\"kafkapb://ttn-eu1?topic=deliver_000013.ttn.ttn-eu1_uplink\",\"receiver_name\":\"router-dataplane-56fc9fb8fd-d9wh4\",\"receiver_agent\":\"pbdataplane/1.2.0-rc.1 go/1.15.6 linux/amd64\"}]},\"rssi\":-96,\"channel_rssi\":-96,\"snr\":7,\"uplink_token\":\"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU25WTk1uUTFVbGhTY1ZwRVRrWmhWVGxFVjIxc2IwbHBkMmxrUjBadVNXcHZhVmt5TlROVVJWSlJXVlZqZW1FeFRrdGpNVXAyVTFac1RrNHdkekJWVTBvNUxuWXhiMGhGUjNoVVFuWlFaVFZ0Y1V4T09GaFpVbWN1Y2xjd1kyUmZaa0pqVm10b05XTkVUQzVVUnpkeFUwcFdXRmRRT0hveVNtOWlXREJwVjFrMlpreFlUV1paYTBJMmNtZG1XVlpFVUc4MVpFcHhNWEoxU1VWYWEySXRORzFPVEcxc016TndNemRUU2toVlpqWk5ZVzl5VDJOME5ETlJRMm90U1ZKU1ExZG9VR2swVG1aQ2IxSlJOVWRJUlhscmFVRkxSblkxWW5SNVRqVjNVMVZxVWtabVRWcHRTVmxzTm5saU9Ib3lOVmh4WWpWSE4yTmpWazVRY0ZJMFNsZEphRUo0ZGs1SVRqTnpTVkF5Ymt4U01DNXRkSEJTYmxGdVgzaEpRMmhJV0hKS00zQXlSV1pSIiwiYSI6eyJmbmlkIjoiMDAwMDEzIiwiZnRpZCI6InR0biIsImZjaWQiOiJ0dG4tdjItZXUtMyJ9fQ==\"}],\"settings\":{\"data_rate\":{\"lora\":{\"bandwidth\":125000,\"spreading_factor\":7}},\"data_rate_index\":5,\"frequency\":\"867900000\"},\"received_at\":\"2021-01-30T08:38:51.031315619Z\"}}"

	var packetIn models.V3ApplicationUp
	if err := json.Unmarshal([]byte(postbody), &packetIn); err != nil {
		log.Print(err.Error())
	}

	log.Println("Decoded payload: ", packetIn.UplinkMessage.DecodedPayload)
}