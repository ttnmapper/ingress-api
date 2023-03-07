// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.2
// - protoc             v3.21.1
// source: lorawan-stack/api/packetbrokeragent.proto

package ttnpb

import (
	golang "github.com/TheThingsIndustries/protoc-gen-go-json/golang"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// MarshalProtoJSON marshals the PacketBrokerGateway_GatewayIdentifiers message to JSON.
func (x *PacketBrokerGateway_GatewayIdentifiers) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.GatewayId != "" || s.HasField("gateway_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_id")
		s.WriteString(x.GatewayId)
	}
	if len(x.Eui) > 0 || s.HasField("eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("eui")
		types.MarshalHEXBytes(s.WithField("eui"), x.Eui)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the PacketBrokerGateway_GatewayIdentifiers to JSON.
func (x *PacketBrokerGateway_GatewayIdentifiers) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the PacketBrokerGateway_GatewayIdentifiers message from JSON.
func (x *PacketBrokerGateway_GatewayIdentifiers) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_id", "gatewayId":
			s.AddField("gateway_id")
			x.GatewayId = s.ReadString()
		case "eui":
			s.AddField("eui")
			x.Eui = types.Unmarshal8Bytes(s.WithField("eui", false))
		}
	})
}

// UnmarshalJSON unmarshals the PacketBrokerGateway_GatewayIdentifiers from JSON.
func (x *PacketBrokerGateway_GatewayIdentifiers) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the PacketBrokerGateway message to JSON.
func (x *PacketBrokerGateway) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		x.Ids.MarshalProtoJSON(s.WithField("ids"))
	}
	if len(x.ContactInfo) > 0 || s.HasField("contact_info") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_info")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ContactInfo {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("contact_info"))
		}
		s.WriteArrayEnd()
	}
	if x.AdministrativeContact != nil || s.HasField("administrative_contact") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("administrative_contact")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.AdministrativeContact)
	}
	if x.TechnicalContact != nil || s.HasField("technical_contact") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("technical_contact")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.TechnicalContact)
	}
	if len(x.Antennas) > 0 || s.HasField("antennas") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("antennas")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Antennas {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("antennas"))
		}
		s.WriteArrayEnd()
	}
	if x.StatusPublic || s.HasField("status_public") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("status_public")
		s.WriteBool(x.StatusPublic)
	}
	if x.LocationPublic || s.HasField("location_public") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("location_public")
		s.WriteBool(x.LocationPublic)
	}
	if len(x.FrequencyPlanIds) > 0 || s.HasField("frequency_plan_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency_plan_ids")
		s.WriteStringArray(x.FrequencyPlanIds)
	}
	if x.UpdateLocationFromStatus || s.HasField("update_location_from_status") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("update_location_from_status")
		s.WriteBool(x.UpdateLocationFromStatus)
	}
	if x.Online || s.HasField("online") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("online")
		s.WriteBool(x.Online)
	}
	if x.RxRate != nil || s.HasField("rx_rate") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rx_rate")
		if x.RxRate == nil {
			s.WriteNil()
		} else {
			s.WriteFloat32(x.RxRate.Value)
		}
	}
	if x.TxRate != nil || s.HasField("tx_rate") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("tx_rate")
		if x.TxRate == nil {
			s.WriteNil()
		} else {
			s.WriteFloat32(x.TxRate.Value)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the PacketBrokerGateway to JSON.
func (x *PacketBrokerGateway) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the PacketBrokerGateway message from JSON.
func (x *PacketBrokerGateway) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			x.Ids = &PacketBrokerGateway_GatewayIdentifiers{}
			x.Ids.UnmarshalProtoJSON(s.WithField("ids", true))
		case "contact_info", "contactInfo":
			s.AddField("contact_info")
			if s.ReadNil() {
				x.ContactInfo = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ContactInfo = append(x.ContactInfo, nil)
					return
				}
				v := &ContactInfo{}
				v.UnmarshalProtoJSON(s.WithField("contact_info", false))
				if s.Err() != nil {
					return
				}
				x.ContactInfo = append(x.ContactInfo, v)
			})
		case "administrative_contact", "administrativeContact":
			s.AddField("administrative_contact")
			if s.ReadNil() {
				x.AdministrativeContact = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.AdministrativeContact = &v
		case "technical_contact", "technicalContact":
			s.AddField("technical_contact")
			if s.ReadNil() {
				x.TechnicalContact = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.TechnicalContact = &v
		case "antennas":
			s.AddField("antennas")
			if s.ReadNil() {
				x.Antennas = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Antennas = append(x.Antennas, nil)
					return
				}
				v := &GatewayAntenna{}
				v.UnmarshalProtoJSON(s.WithField("antennas", false))
				if s.Err() != nil {
					return
				}
				x.Antennas = append(x.Antennas, v)
			})
		case "status_public", "statusPublic":
			s.AddField("status_public")
			x.StatusPublic = s.ReadBool()
		case "location_public", "locationPublic":
			s.AddField("location_public")
			x.LocationPublic = s.ReadBool()
		case "frequency_plan_ids", "frequencyPlanIds":
			s.AddField("frequency_plan_ids")
			if s.ReadNil() {
				x.FrequencyPlanIds = nil
				return
			}
			x.FrequencyPlanIds = s.ReadStringArray()
		case "update_location_from_status", "updateLocationFromStatus":
			s.AddField("update_location_from_status")
			x.UpdateLocationFromStatus = s.ReadBool()
		case "online":
			s.AddField("online")
			x.Online = s.ReadBool()
		case "rx_rate", "rxRate":
			s.AddField("rx_rate")
			if s.ReadNil() {
				x.RxRate = nil
				return
			}
			v := s.ReadWrappedFloat32()
			if s.Err() != nil {
				return
			}
			x.RxRate = &wrapperspb.FloatValue{Value: v}
		case "tx_rate", "txRate":
			s.AddField("tx_rate")
			if s.ReadNil() {
				x.TxRate = nil
				return
			}
			v := s.ReadWrappedFloat32()
			if s.Err() != nil {
				return
			}
			x.TxRate = &wrapperspb.FloatValue{Value: v}
		}
	})
}

// UnmarshalJSON unmarshals the PacketBrokerGateway from JSON.
func (x *PacketBrokerGateway) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the UpdatePacketBrokerGatewayRequest message to JSON.
func (x *UpdatePacketBrokerGatewayRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Gateway != nil || s.HasField("gateway") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway")
		x.Gateway.MarshalProtoJSON(s.WithField("gateway"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			golang.MarshalLegacyFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the UpdatePacketBrokerGatewayRequest to JSON.
func (x *UpdatePacketBrokerGatewayRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the UpdatePacketBrokerGatewayRequest message from JSON.
func (x *UpdatePacketBrokerGatewayRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway":
			if s.ReadNil() {
				x.Gateway = nil
				return
			}
			x.Gateway = &PacketBrokerGateway{}
			x.Gateway.UnmarshalProtoJSON(s.WithField("gateway", true))
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := golang.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// UnmarshalJSON unmarshals the UpdatePacketBrokerGatewayRequest from JSON.
func (x *UpdatePacketBrokerGatewayRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the PacketBrokerDevAddrBlock message to JSON.
func (x *PacketBrokerDevAddrBlock) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.DevAddrPrefix != nil || s.HasField("dev_addr_prefix") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_addr_prefix")
		x.DevAddrPrefix.MarshalProtoJSON(s.WithField("dev_addr_prefix"))
	}
	if x.HomeNetworkClusterId != "" || s.HasField("home_network_cluster_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("home_network_cluster_id")
		s.WriteString(x.HomeNetworkClusterId)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the PacketBrokerDevAddrBlock to JSON.
func (x *PacketBrokerDevAddrBlock) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the PacketBrokerDevAddrBlock message from JSON.
func (x *PacketBrokerDevAddrBlock) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "dev_addr_prefix", "devAddrPrefix":
			if s.ReadNil() {
				x.DevAddrPrefix = nil
				return
			}
			x.DevAddrPrefix = &DevAddrPrefix{}
			x.DevAddrPrefix.UnmarshalProtoJSON(s.WithField("dev_addr_prefix", true))
		case "home_network_cluster_id", "homeNetworkClusterId":
			s.AddField("home_network_cluster_id")
			x.HomeNetworkClusterId = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the PacketBrokerDevAddrBlock from JSON.
func (x *PacketBrokerDevAddrBlock) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the PacketBrokerNetwork message to JSON.
func (x *PacketBrokerNetwork) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Id != nil || s.HasField("id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("id")
		// NOTE: PacketBrokerNetworkIdentifier does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.Id)
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if len(x.DevAddrBlocks) > 0 || s.HasField("dev_addr_blocks") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_addr_blocks")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.DevAddrBlocks {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("dev_addr_blocks"))
		}
		s.WriteArrayEnd()
	}
	if len(x.ContactInfo) > 0 || s.HasField("contact_info") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_info")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ContactInfo {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("contact_info"))
		}
		s.WriteArrayEnd()
	}
	if x.AdministrativeContact != nil || s.HasField("administrative_contact") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("administrative_contact")
		x.AdministrativeContact.MarshalProtoJSON(s.WithField("administrative_contact"))
	}
	if x.TechnicalContact != nil || s.HasField("technical_contact") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("technical_contact")
		x.TechnicalContact.MarshalProtoJSON(s.WithField("technical_contact"))
	}
	if x.Listed || s.HasField("listed") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("listed")
		s.WriteBool(x.Listed)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the PacketBrokerNetwork to JSON.
func (x *PacketBrokerNetwork) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the PacketBrokerNetwork message from JSON.
func (x *PacketBrokerNetwork) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "id":
			s.AddField("id")
			if s.ReadNil() {
				x.Id = nil
				return
			}
			// NOTE: PacketBrokerNetworkIdentifier does not seem to implement UnmarshalProtoJSON.
			var v PacketBrokerNetworkIdentifier
			golang.UnmarshalMessage(s, &v)
			x.Id = &v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "dev_addr_blocks", "devAddrBlocks":
			s.AddField("dev_addr_blocks")
			if s.ReadNil() {
				x.DevAddrBlocks = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.DevAddrBlocks = append(x.DevAddrBlocks, nil)
					return
				}
				v := &PacketBrokerDevAddrBlock{}
				v.UnmarshalProtoJSON(s.WithField("dev_addr_blocks", false))
				if s.Err() != nil {
					return
				}
				x.DevAddrBlocks = append(x.DevAddrBlocks, v)
			})
		case "contact_info", "contactInfo":
			s.AddField("contact_info")
			if s.ReadNil() {
				x.ContactInfo = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ContactInfo = append(x.ContactInfo, nil)
					return
				}
				v := &ContactInfo{}
				v.UnmarshalProtoJSON(s.WithField("contact_info", false))
				if s.Err() != nil {
					return
				}
				x.ContactInfo = append(x.ContactInfo, v)
			})
		case "administrative_contact", "administrativeContact":
			if s.ReadNil() {
				x.AdministrativeContact = nil
				return
			}
			x.AdministrativeContact = &ContactInfo{}
			x.AdministrativeContact.UnmarshalProtoJSON(s.WithField("administrative_contact", true))
		case "technical_contact", "technicalContact":
			if s.ReadNil() {
				x.TechnicalContact = nil
				return
			}
			x.TechnicalContact = &ContactInfo{}
			x.TechnicalContact.UnmarshalProtoJSON(s.WithField("technical_contact", true))
		case "listed":
			s.AddField("listed")
			x.Listed = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the PacketBrokerNetwork from JSON.
func (x *PacketBrokerNetwork) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the PacketBrokerNetworks message to JSON.
func (x *PacketBrokerNetworks) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Networks) > 0 || s.HasField("networks") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("networks")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Networks {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("networks"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the PacketBrokerNetworks to JSON.
func (x *PacketBrokerNetworks) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the PacketBrokerNetworks message from JSON.
func (x *PacketBrokerNetworks) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "networks":
			s.AddField("networks")
			if s.ReadNil() {
				x.Networks = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Networks = append(x.Networks, nil)
					return
				}
				v := &PacketBrokerNetwork{}
				v.UnmarshalProtoJSON(s.WithField("networks", false))
				if s.Err() != nil {
					return
				}
				x.Networks = append(x.Networks, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the PacketBrokerNetworks from JSON.
func (x *PacketBrokerNetworks) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the PacketBrokerInfo message to JSON.
func (x *PacketBrokerInfo) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Registration != nil || s.HasField("registration") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("registration")
		x.Registration.MarshalProtoJSON(s.WithField("registration"))
	}
	if x.ForwarderEnabled || s.HasField("forwarder_enabled") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("forwarder_enabled")
		s.WriteBool(x.ForwarderEnabled)
	}
	if x.HomeNetworkEnabled || s.HasField("home_network_enabled") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("home_network_enabled")
		s.WriteBool(x.HomeNetworkEnabled)
	}
	if x.RegisterEnabled || s.HasField("register_enabled") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("register_enabled")
		s.WriteBool(x.RegisterEnabled)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the PacketBrokerInfo to JSON.
func (x *PacketBrokerInfo) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the PacketBrokerInfo message from JSON.
func (x *PacketBrokerInfo) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "registration":
			if s.ReadNil() {
				x.Registration = nil
				return
			}
			x.Registration = &PacketBrokerNetwork{}
			x.Registration.UnmarshalProtoJSON(s.WithField("registration", true))
		case "forwarder_enabled", "forwarderEnabled":
			s.AddField("forwarder_enabled")
			x.ForwarderEnabled = s.ReadBool()
		case "home_network_enabled", "homeNetworkEnabled":
			s.AddField("home_network_enabled")
			x.HomeNetworkEnabled = s.ReadBool()
		case "register_enabled", "registerEnabled":
			s.AddField("register_enabled")
			x.RegisterEnabled = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the PacketBrokerInfo from JSON.
func (x *PacketBrokerInfo) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}