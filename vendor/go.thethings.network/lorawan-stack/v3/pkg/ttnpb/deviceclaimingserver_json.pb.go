// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.0
// - protoc             v3.9.1
// source: lorawan-stack/api/deviceclaimingserver.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// MarshalProtoJSON marshals the ClaimEndDeviceRequest_AuthenticatedIdentifiers message to JSON.
func (x *ClaimEndDeviceRequest_AuthenticatedIdentifiers) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	if len(x.DevEui) > 0 || s.HasField("dev_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_eui")
		types.MarshalHEXBytes(s.WithField("dev_eui"), x.DevEui)
	}
	if x.AuthenticationCode != "" || s.HasField("authentication_code") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("authentication_code")
		s.WriteString(x.AuthenticationCode)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ClaimEndDeviceRequest_AuthenticatedIdentifiers to JSON.
func (x *ClaimEndDeviceRequest_AuthenticatedIdentifiers) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ClaimEndDeviceRequest_AuthenticatedIdentifiers message from JSON.
func (x *ClaimEndDeviceRequest_AuthenticatedIdentifiers) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		case "dev_eui", "devEui":
			s.AddField("dev_eui")
			x.DevEui = types.Unmarshal8Bytes(s.WithField("dev_eui", false))
		case "authentication_code", "authenticationCode":
			s.AddField("authentication_code")
			x.AuthenticationCode = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the ClaimEndDeviceRequest_AuthenticatedIdentifiers from JSON.
func (x *ClaimEndDeviceRequest_AuthenticatedIdentifiers) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ClaimEndDeviceRequest message to JSON.
func (x *ClaimEndDeviceRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.SourceDevice != nil {
		switch ov := x.SourceDevice.(type) {
		case *ClaimEndDeviceRequest_AuthenticatedIdentifiers_:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("authenticated_identifiers")
			ov.AuthenticatedIdentifiers.MarshalProtoJSON(s.WithField("authenticated_identifiers"))
		case *ClaimEndDeviceRequest_QrCode:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("qr_code")
			s.WriteBytes(ov.QrCode)
		}
	}
	if x.TargetApplicationIds != nil || s.HasField("target_application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.TargetApplicationIds)
	}
	if x.TargetDeviceId != "" || s.HasField("target_device_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_device_id")
		s.WriteString(x.TargetDeviceId)
	}
	if x.TargetNetworkServerAddress != "" || s.HasField("target_network_server_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_network_server_address")
		s.WriteString(x.TargetNetworkServerAddress)
	}
	if x.TargetNetworkServerKekLabel != "" || s.HasField("target_network_server_kek_label") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_network_server_kek_label")
		s.WriteString(x.TargetNetworkServerKekLabel)
	}
	if x.TargetApplicationServerAddress != "" || s.HasField("target_application_server_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_application_server_address")
		s.WriteString(x.TargetApplicationServerAddress)
	}
	if x.TargetApplicationServerKekLabel != "" || s.HasField("target_application_server_kek_label") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_application_server_kek_label")
		s.WriteString(x.TargetApplicationServerKekLabel)
	}
	if x.TargetApplicationServerId != "" || s.HasField("target_application_server_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_application_server_id")
		s.WriteString(x.TargetApplicationServerId)
	}
	if len(x.TargetNetId) > 0 || s.HasField("target_net_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_net_id")
		types.MarshalHEXBytes(s.WithField("target_net_id"), x.TargetNetId)
	}
	if x.InvalidateAuthenticationCode || s.HasField("invalidate_authentication_code") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("invalidate_authentication_code")
		s.WriteBool(x.InvalidateAuthenticationCode)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ClaimEndDeviceRequest to JSON.
func (x *ClaimEndDeviceRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ClaimEndDeviceRequest message from JSON.
func (x *ClaimEndDeviceRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "authenticated_identifiers", "authenticatedIdentifiers":
			ov := &ClaimEndDeviceRequest_AuthenticatedIdentifiers_{}
			x.SourceDevice = ov
			if s.ReadNil() {
				ov.AuthenticatedIdentifiers = nil
				return
			}
			ov.AuthenticatedIdentifiers = &ClaimEndDeviceRequest_AuthenticatedIdentifiers{}
			ov.AuthenticatedIdentifiers.UnmarshalProtoJSON(s.WithField("authenticated_identifiers", true))
		case "qr_code", "qrCode":
			s.AddField("qr_code")
			ov := &ClaimEndDeviceRequest_QrCode{}
			x.SourceDevice = ov
			ov.QrCode = s.ReadBytes()
		case "target_application_ids", "targetApplicationIds":
			s.AddField("target_application_ids")
			if s.ReadNil() {
				x.TargetApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.TargetApplicationIds = &v
		case "target_device_id", "targetDeviceId":
			s.AddField("target_device_id")
			x.TargetDeviceId = s.ReadString()
		case "target_network_server_address", "targetNetworkServerAddress":
			s.AddField("target_network_server_address")
			x.TargetNetworkServerAddress = s.ReadString()
		case "target_network_server_kek_label", "targetNetworkServerKekLabel":
			s.AddField("target_network_server_kek_label")
			x.TargetNetworkServerKekLabel = s.ReadString()
		case "target_application_server_address", "targetApplicationServerAddress":
			s.AddField("target_application_server_address")
			x.TargetApplicationServerAddress = s.ReadString()
		case "target_application_server_kek_label", "targetApplicationServerKekLabel":
			s.AddField("target_application_server_kek_label")
			x.TargetApplicationServerKekLabel = s.ReadString()
		case "target_application_server_id", "targetApplicationServerId":
			s.AddField("target_application_server_id")
			x.TargetApplicationServerId = s.ReadString()
		case "target_net_id", "targetNetId":
			s.AddField("target_net_id")
			x.TargetNetId = types.Unmarshal3Bytes(s.WithField("target_net_id", false))
		case "invalidate_authentication_code", "invalidateAuthenticationCode":
			s.AddField("invalidate_authentication_code")
			x.InvalidateAuthenticationCode = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the ClaimEndDeviceRequest from JSON.
func (x *ClaimEndDeviceRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetInfoByJoinEUIRequest message to JSON.
func (x *GetInfoByJoinEUIRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetInfoByJoinEUIRequest to JSON.
func (x *GetInfoByJoinEUIRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetInfoByJoinEUIRequest message from JSON.
func (x *GetInfoByJoinEUIRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		}
	})
}

// UnmarshalJSON unmarshals the GetInfoByJoinEUIRequest from JSON.
func (x *GetInfoByJoinEUIRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetInfoByJoinEUIResponse message to JSON.
func (x *GetInfoByJoinEUIResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	if x.SupportsClaiming || s.HasField("supports_claiming") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("supports_claiming")
		s.WriteBool(x.SupportsClaiming)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetInfoByJoinEUIResponse to JSON.
func (x *GetInfoByJoinEUIResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetInfoByJoinEUIResponse message from JSON.
func (x *GetInfoByJoinEUIResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		case "supports_claiming", "supportsClaiming":
			s.AddField("supports_claiming")
			x.SupportsClaiming = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the GetInfoByJoinEUIResponse from JSON.
func (x *GetInfoByJoinEUIResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetClaimStatusResponse message to JSON.
func (x *GetClaimStatusResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.EndDeviceIds != nil || s.HasField("end_device_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("end_device_ids")
		x.EndDeviceIds.MarshalProtoJSON(s.WithField("end_device_ids"))
	}
	if len(x.HomeNetId) > 0 || s.HasField("home_net_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("home_net_id")
		types.MarshalHEXBytes(s.WithField("home_net_id"), x.HomeNetId)
	}
	if len(x.HomeNsId) > 0 || s.HasField("home_ns_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("home_ns_id")
		types.MarshalHEXBytes(s.WithField("home_ns_id"), x.HomeNsId)
	}
	if x.VendorSpecific != nil || s.HasField("vendor_specific") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("vendor_specific")
		// NOTE: GetClaimStatusResponse_VendorSpecific does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.VendorSpecific)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetClaimStatusResponse to JSON.
func (x *GetClaimStatusResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetClaimStatusResponse message from JSON.
func (x *GetClaimStatusResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "end_device_ids", "endDeviceIds":
			if s.ReadNil() {
				x.EndDeviceIds = nil
				return
			}
			x.EndDeviceIds = &EndDeviceIdentifiers{}
			x.EndDeviceIds.UnmarshalProtoJSON(s.WithField("end_device_ids", true))
		case "home_net_id", "homeNetId":
			s.AddField("home_net_id")
			x.HomeNetId = types.Unmarshal3Bytes(s.WithField("home_net_id", false))
		case "home_ns_id", "homeNsId":
			s.AddField("home_ns_id")
			x.HomeNsId = types.Unmarshal8Bytes(s.WithField("home_ns_id", false))
		case "vendor_specific", "vendorSpecific":
			s.AddField("vendor_specific")
			if s.ReadNil() {
				x.VendorSpecific = nil
				return
			}
			// NOTE: GetClaimStatusResponse_VendorSpecific does not seem to implement UnmarshalProtoJSON.
			var v GetClaimStatusResponse_VendorSpecific
			gogo.UnmarshalMessage(s, &v)
			x.VendorSpecific = &v
		}
	})
}

// UnmarshalJSON unmarshals the GetClaimStatusResponse from JSON.
func (x *GetClaimStatusResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ClaimGatewayRequest_AuthenticatedIdentifiers message to JSON.
func (x *ClaimGatewayRequest_AuthenticatedIdentifiers) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.GatewayEui) > 0 || s.HasField("gateway_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_eui")
		types.MarshalHEXBytes(s.WithField("gateway_eui"), x.GatewayEui)
	}
	if len(x.AuthenticationCode) > 0 || s.HasField("authentication_code") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("authentication_code")
		s.WriteBytes(x.AuthenticationCode)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ClaimGatewayRequest_AuthenticatedIdentifiers to JSON.
func (x *ClaimGatewayRequest_AuthenticatedIdentifiers) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ClaimGatewayRequest_AuthenticatedIdentifiers message from JSON.
func (x *ClaimGatewayRequest_AuthenticatedIdentifiers) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_eui", "gatewayEui":
			s.AddField("gateway_eui")
			x.GatewayEui = types.Unmarshal8Bytes(s.WithField("gateway_eui", false))
		case "authentication_code", "authenticationCode":
			s.AddField("authentication_code")
			x.AuthenticationCode = s.ReadBytes()
		}
	})
}

// UnmarshalJSON unmarshals the ClaimGatewayRequest_AuthenticatedIdentifiers from JSON.
func (x *ClaimGatewayRequest_AuthenticatedIdentifiers) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ClaimGatewayRequest message to JSON.
func (x *ClaimGatewayRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.SourceGateway != nil {
		switch ov := x.SourceGateway.(type) {
		case *ClaimGatewayRequest_AuthenticatedIdentifiers_:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("authenticated_identifiers")
			ov.AuthenticatedIdentifiers.MarshalProtoJSON(s.WithField("authenticated_identifiers"))
		case *ClaimGatewayRequest_QrCode:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("qr_code")
			s.WriteBytes(ov.QrCode)
		}
	}
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Collaborator)
	}
	if x.TargetGatewayId != "" || s.HasField("target_gateway_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_gateway_id")
		s.WriteString(x.TargetGatewayId)
	}
	if x.TargetGatewayServerAddress != "" || s.HasField("target_gateway_server_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_gateway_server_address")
		s.WriteString(x.TargetGatewayServerAddress)
	}
	if x.CupsRedirection != nil || s.HasField("cups_redirection") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("cups_redirection")
		// NOTE: CUPSRedirection does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.CupsRedirection)
	}
	if x.TargetFrequencyPlanId != "" || s.HasField("target_frequency_plan_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("target_frequency_plan_id")
		s.WriteString(x.TargetFrequencyPlanId)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ClaimGatewayRequest to JSON.
func (x *ClaimGatewayRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ClaimGatewayRequest message from JSON.
func (x *ClaimGatewayRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "authenticated_identifiers", "authenticatedIdentifiers":
			ov := &ClaimGatewayRequest_AuthenticatedIdentifiers_{}
			x.SourceGateway = ov
			if s.ReadNil() {
				ov.AuthenticatedIdentifiers = nil
				return
			}
			ov.AuthenticatedIdentifiers = &ClaimGatewayRequest_AuthenticatedIdentifiers{}
			ov.AuthenticatedIdentifiers.UnmarshalProtoJSON(s.WithField("authenticated_identifiers", true))
		case "qr_code", "qrCode":
			s.AddField("qr_code")
			ov := &ClaimGatewayRequest_QrCode{}
			x.SourceGateway = ov
			ov.QrCode = s.ReadBytes()
		case "collaborator":
			s.AddField("collaborator")
			if s.ReadNil() {
				x.Collaborator = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Collaborator = &v
		case "target_gateway_id", "targetGatewayId":
			s.AddField("target_gateway_id")
			x.TargetGatewayId = s.ReadString()
		case "target_gateway_server_address", "targetGatewayServerAddress":
			s.AddField("target_gateway_server_address")
			x.TargetGatewayServerAddress = s.ReadString()
		case "cups_redirection", "cupsRedirection":
			s.AddField("cups_redirection")
			if s.ReadNil() {
				x.CupsRedirection = nil
				return
			}
			// NOTE: CUPSRedirection does not seem to implement UnmarshalProtoJSON.
			var v CUPSRedirection
			gogo.UnmarshalMessage(s, &v)
			x.CupsRedirection = &v
		case "target_frequency_plan_id", "targetFrequencyPlanId":
			s.AddField("target_frequency_plan_id")
			x.TargetFrequencyPlanId = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the ClaimGatewayRequest from JSON.
func (x *ClaimGatewayRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the AuthorizeGatewayRequest message to JSON.
func (x *AuthorizeGatewayRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.GatewayIds != nil || s.HasField("gateway_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_ids")
		x.GatewayIds.MarshalProtoJSON(s.WithField("gateway_ids"))
	}
	if x.ApiKey != "" || s.HasField("api_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("api_key")
		s.WriteString(x.ApiKey)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the AuthorizeGatewayRequest to JSON.
func (x *AuthorizeGatewayRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the AuthorizeGatewayRequest message from JSON.
func (x *AuthorizeGatewayRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_ids", "gatewayIds":
			if s.ReadNil() {
				x.GatewayIds = nil
				return
			}
			x.GatewayIds = &GatewayIdentifiers{}
			x.GatewayIds.UnmarshalProtoJSON(s.WithField("gateway_ids", true))
		case "api_key", "apiKey":
			s.AddField("api_key")
			x.ApiKey = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the AuthorizeGatewayRequest from JSON.
func (x *AuthorizeGatewayRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
