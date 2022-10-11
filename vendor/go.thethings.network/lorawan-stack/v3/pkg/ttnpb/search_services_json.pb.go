// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.0
// - protoc             v3.9.1
// source: lorawan-stack/api/search_services.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the SearchClientsRequest message to JSON.
func (x *SearchClientsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Query != "" || s.HasField("query") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("query")
		s.WriteString(x.Query)
	}
	if x.IdContains != "" || s.HasField("id_contains") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("id_contains")
		s.WriteString(x.IdContains)
	}
	if x.NameContains != "" || s.HasField("name_contains") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name_contains")
		s.WriteString(x.NameContains)
	}
	if x.DescriptionContains != "" || s.HasField("description_contains") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("description_contains")
		s.WriteString(x.DescriptionContains)
	}
	if x.AttributesContain != nil || s.HasField("attributes_contain") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("attributes_contain")
		s.WriteObjectStart()
		var wroteElement bool
		for k, v := range x.AttributesContain {
			s.WriteMoreIf(&wroteElement)
			s.WriteObjectStringField(k)
			s.WriteString(v)
		}
		s.WriteObjectEnd()
	}
	if len(x.State) > 0 || s.HasField("state") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.State {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	if x.Order != "" || s.HasField("order") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("order")
		s.WriteString(x.Order)
	}
	if x.Limit != 0 || s.HasField("limit") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("limit")
		s.WriteUint32(x.Limit)
	}
	if x.Page != 0 || s.HasField("page") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("page")
		s.WriteUint32(x.Page)
	}
	if x.Deleted || s.HasField("deleted") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("deleted")
		s.WriteBool(x.Deleted)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SearchClientsRequest to JSON.
func (x *SearchClientsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the SearchClientsRequest message from JSON.
func (x *SearchClientsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "query":
			s.AddField("query")
			x.Query = s.ReadString()
		case "id_contains", "idContains":
			s.AddField("id_contains")
			x.IdContains = s.ReadString()
		case "name_contains", "nameContains":
			s.AddField("name_contains")
			x.NameContains = s.ReadString()
		case "description_contains", "descriptionContains":
			s.AddField("description_contains")
			x.DescriptionContains = s.ReadString()
		case "attributes_contain", "attributesContain":
			s.AddField("attributes_contain")
			if s.ReadNil() {
				x.AttributesContain = nil
				return
			}
			x.AttributesContain = make(map[string]string)
			s.ReadStringMap(func(key string) {
				x.AttributesContain[key] = s.ReadString()
			})
		case "state":
			s.AddField("state")
			if s.ReadNil() {
				x.State = nil
				return
			}
			s.ReadArray(func() {
				var v State
				v.UnmarshalProtoJSON(s)
				x.State = append(x.State, v)
			})
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		case "order":
			s.AddField("order")
			x.Order = s.ReadString()
		case "limit":
			s.AddField("limit")
			x.Limit = s.ReadUint32()
		case "page":
			s.AddField("page")
			x.Page = s.ReadUint32()
		case "deleted":
			s.AddField("deleted")
			x.Deleted = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the SearchClientsRequest from JSON.
func (x *SearchClientsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the SearchUsersRequest message to JSON.
func (x *SearchUsersRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Query != "" || s.HasField("query") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("query")
		s.WriteString(x.Query)
	}
	if x.IdContains != "" || s.HasField("id_contains") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("id_contains")
		s.WriteString(x.IdContains)
	}
	if x.NameContains != "" || s.HasField("name_contains") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name_contains")
		s.WriteString(x.NameContains)
	}
	if x.DescriptionContains != "" || s.HasField("description_contains") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("description_contains")
		s.WriteString(x.DescriptionContains)
	}
	if x.AttributesContain != nil || s.HasField("attributes_contain") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("attributes_contain")
		s.WriteObjectStart()
		var wroteElement bool
		for k, v := range x.AttributesContain {
			s.WriteMoreIf(&wroteElement)
			s.WriteObjectStringField(k)
			s.WriteString(v)
		}
		s.WriteObjectEnd()
	}
	if len(x.State) > 0 || s.HasField("state") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.State {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	if x.Order != "" || s.HasField("order") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("order")
		s.WriteString(x.Order)
	}
	if x.Limit != 0 || s.HasField("limit") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("limit")
		s.WriteUint32(x.Limit)
	}
	if x.Page != 0 || s.HasField("page") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("page")
		s.WriteUint32(x.Page)
	}
	if x.Deleted || s.HasField("deleted") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("deleted")
		s.WriteBool(x.Deleted)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SearchUsersRequest to JSON.
func (x *SearchUsersRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the SearchUsersRequest message from JSON.
func (x *SearchUsersRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "query":
			s.AddField("query")
			x.Query = s.ReadString()
		case "id_contains", "idContains":
			s.AddField("id_contains")
			x.IdContains = s.ReadString()
		case "name_contains", "nameContains":
			s.AddField("name_contains")
			x.NameContains = s.ReadString()
		case "description_contains", "descriptionContains":
			s.AddField("description_contains")
			x.DescriptionContains = s.ReadString()
		case "attributes_contain", "attributesContain":
			s.AddField("attributes_contain")
			if s.ReadNil() {
				x.AttributesContain = nil
				return
			}
			x.AttributesContain = make(map[string]string)
			s.ReadStringMap(func(key string) {
				x.AttributesContain[key] = s.ReadString()
			})
		case "state":
			s.AddField("state")
			if s.ReadNil() {
				x.State = nil
				return
			}
			s.ReadArray(func() {
				var v State
				v.UnmarshalProtoJSON(s)
				x.State = append(x.State, v)
			})
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		case "order":
			s.AddField("order")
			x.Order = s.ReadString()
		case "limit":
			s.AddField("limit")
			x.Limit = s.ReadUint32()
		case "page":
			s.AddField("page")
			x.Page = s.ReadUint32()
		case "deleted":
			s.AddField("deleted")
			x.Deleted = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the SearchUsersRequest from JSON.
func (x *SearchUsersRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the SearchAccountsRequest message to JSON.
func (x *SearchAccountsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Query != "" || s.HasField("query") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("query")
		s.WriteString(x.Query)
	}
	if x.OnlyUsers || s.HasField("only_users") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("only_users")
		s.WriteBool(x.OnlyUsers)
	}
	if x.CollaboratorOf != nil {
		switch ov := x.CollaboratorOf.(type) {
		case *SearchAccountsRequest_ApplicationIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("application_ids")
			// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
			gogo.MarshalMessage(s, ov.ApplicationIds)
		case *SearchAccountsRequest_ClientIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("client_ids")
			// NOTE: ClientIdentifiers does not seem to implement MarshalProtoJSON.
			gogo.MarshalMessage(s, ov.ClientIds)
		case *SearchAccountsRequest_GatewayIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("gateway_ids")
			ov.GatewayIds.MarshalProtoJSON(s.WithField("gateway_ids"))
		case *SearchAccountsRequest_OrganizationIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("organization_ids")
			// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
			gogo.MarshalMessage(s, ov.OrganizationIds)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SearchAccountsRequest to JSON.
func (x *SearchAccountsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the SearchAccountsRequest message from JSON.
func (x *SearchAccountsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "query":
			s.AddField("query")
			x.Query = s.ReadString()
		case "only_users", "onlyUsers":
			s.AddField("only_users")
			x.OnlyUsers = s.ReadBool()
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			ov := &SearchAccountsRequest_ApplicationIds{}
			x.CollaboratorOf = ov
			if s.ReadNil() {
				ov.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			ov.ApplicationIds = &v
		case "client_ids", "clientIds":
			s.AddField("client_ids")
			ov := &SearchAccountsRequest_ClientIds{}
			x.CollaboratorOf = ov
			if s.ReadNil() {
				ov.ClientIds = nil
				return
			}
			// NOTE: ClientIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ClientIdentifiers
			gogo.UnmarshalMessage(s, &v)
			ov.ClientIds = &v
		case "gateway_ids", "gatewayIds":
			ov := &SearchAccountsRequest_GatewayIds{}
			x.CollaboratorOf = ov
			if s.ReadNil() {
				ov.GatewayIds = nil
				return
			}
			ov.GatewayIds = &GatewayIdentifiers{}
			ov.GatewayIds.UnmarshalProtoJSON(s.WithField("gateway_ids", true))
		case "organization_ids", "organizationIds":
			s.AddField("organization_ids")
			ov := &SearchAccountsRequest_OrganizationIds{}
			x.CollaboratorOf = ov
			if s.ReadNil() {
				ov.OrganizationIds = nil
				return
			}
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			ov.OrganizationIds = &v
		}
	})
}

// UnmarshalJSON unmarshals the SearchAccountsRequest from JSON.
func (x *SearchAccountsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
