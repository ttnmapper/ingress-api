// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var RxMetadataFieldPathsNested = []string{
	"advanced",
	"antenna_index",
	"channel_index",
	"channel_rssi",
	"downlink_path_constraint",
	"encrypted_fine_timestamp",
	"encrypted_fine_timestamp_key_id",
	"fine_timestamp",
	"frequency_drift",
	"frequency_offset",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
	"gps_time",
	"hopping_width",
	"location",
	"location.accuracy",
	"location.altitude",
	"location.latitude",
	"location.longitude",
	"location.source",
	"packet_broker",
	"packet_broker.forwarder_cluster_id",
	"packet_broker.forwarder_gateway_eui",
	"packet_broker.forwarder_gateway_id",
	"packet_broker.forwarder_net_id",
	"packet_broker.forwarder_tenant_id",
	"packet_broker.home_network_cluster_id",
	"packet_broker.home_network_net_id",
	"packet_broker.home_network_tenant_id",
	"packet_broker.hops",
	"packet_broker.message_id",
	"received_at",
	"rssi",
	"rssi_standard_deviation",
	"signal_rssi",
	"snr",
	"time",
	"timestamp",
	"uplink_token",
}

var RxMetadataFieldPathsTopLevel = []string{
	"advanced",
	"antenna_index",
	"channel_index",
	"channel_rssi",
	"downlink_path_constraint",
	"encrypted_fine_timestamp",
	"encrypted_fine_timestamp_key_id",
	"fine_timestamp",
	"frequency_drift",
	"frequency_offset",
	"gateway_ids",
	"gps_time",
	"hopping_width",
	"location",
	"packet_broker",
	"received_at",
	"rssi",
	"rssi_standard_deviation",
	"signal_rssi",
	"snr",
	"time",
	"timestamp",
	"uplink_token",
}
var LocationFieldPathsNested = []string{
	"accuracy",
	"altitude",
	"latitude",
	"longitude",
	"source",
}

var LocationFieldPathsTopLevel = []string{
	"accuracy",
	"altitude",
	"latitude",
	"longitude",
	"source",
}
var PacketBrokerMetadataFieldPathsNested = []string{
	"forwarder_cluster_id",
	"forwarder_gateway_eui",
	"forwarder_gateway_id",
	"forwarder_net_id",
	"forwarder_tenant_id",
	"home_network_cluster_id",
	"home_network_net_id",
	"home_network_tenant_id",
	"hops",
	"message_id",
}

var PacketBrokerMetadataFieldPathsTopLevel = []string{
	"forwarder_cluster_id",
	"forwarder_gateway_eui",
	"forwarder_gateway_id",
	"forwarder_net_id",
	"forwarder_tenant_id",
	"home_network_cluster_id",
	"home_network_net_id",
	"home_network_tenant_id",
	"hops",
	"message_id",
}
var PacketBrokerRouteHopFieldPathsNested = []string{
	"received_at",
	"receiver_agent",
	"receiver_name",
	"sender_address",
	"sender_name",
}

var PacketBrokerRouteHopFieldPathsTopLevel = []string{
	"received_at",
	"receiver_agent",
	"receiver_name",
	"sender_address",
	"sender_name",
}
