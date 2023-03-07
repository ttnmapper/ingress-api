// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.2
// - protoc             v3.21.1
// source: lorawan-stack/api/application.proto

package ttnpb

import (
	golang "github.com/TheThingsIndustries/protoc-gen-go-json/golang"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// MarshalProtoJSON marshals the Application message to JSON.
func (x *Application) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.Ids)
	}
	if x.CreatedAt != nil || s.HasField("created_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("created_at")
		if x.CreatedAt == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.CreatedAt)
		}
	}
	if x.UpdatedAt != nil || s.HasField("updated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("updated_at")
		if x.UpdatedAt == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.UpdatedAt)
		}
	}
	if x.DeletedAt != nil || s.HasField("deleted_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("deleted_at")
		if x.DeletedAt == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.DeletedAt)
		}
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if x.Description != "" || s.HasField("description") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("description")
		s.WriteString(x.Description)
	}
	if x.Attributes != nil || s.HasField("attributes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("attributes")
		s.WriteObjectStart()
		var wroteElement bool
		for k, v := range x.Attributes {
			s.WriteMoreIf(&wroteElement)
			s.WriteObjectStringField(k)
			s.WriteString(v)
		}
		s.WriteObjectEnd()
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
	if x.NetworkServerAddress != "" || s.HasField("network_server_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("network_server_address")
		s.WriteString(x.NetworkServerAddress)
	}
	if x.ApplicationServerAddress != "" || s.HasField("application_server_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_server_address")
		s.WriteString(x.ApplicationServerAddress)
	}
	if x.JoinServerAddress != "" || s.HasField("join_server_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_server_address")
		s.WriteString(x.JoinServerAddress)
	}
	if x.DevEuiCounter != 0 || s.HasField("dev_eui_counter") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_eui_counter")
		s.WriteUint32(x.DevEuiCounter)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Application to JSON.
func (x *Application) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the Application message from JSON.
func (x *Application) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "created_at", "createdAt":
			s.AddField("created_at")
			if s.ReadNil() {
				x.CreatedAt = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.CreatedAt = v
		case "updated_at", "updatedAt":
			s.AddField("updated_at")
			if s.ReadNil() {
				x.UpdatedAt = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.UpdatedAt = v
		case "deleted_at", "deletedAt":
			s.AddField("deleted_at")
			if s.ReadNil() {
				x.DeletedAt = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.DeletedAt = v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "description":
			s.AddField("description")
			x.Description = s.ReadString()
		case "attributes":
			s.AddField("attributes")
			if s.ReadNil() {
				x.Attributes = nil
				return
			}
			x.Attributes = make(map[string]string)
			s.ReadStringMap(func(key string) {
				x.Attributes[key] = s.ReadString()
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
		case "network_server_address", "networkServerAddress":
			s.AddField("network_server_address")
			x.NetworkServerAddress = s.ReadString()
		case "application_server_address", "applicationServerAddress":
			s.AddField("application_server_address")
			x.ApplicationServerAddress = s.ReadString()
		case "join_server_address", "joinServerAddress":
			s.AddField("join_server_address")
			x.JoinServerAddress = s.ReadString()
		case "dev_eui_counter", "devEuiCounter":
			s.AddField("dev_eui_counter")
			x.DevEuiCounter = s.ReadUint32()
		}
	})
}

// UnmarshalJSON unmarshals the Application from JSON.
func (x *Application) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the Applications message to JSON.
func (x *Applications) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Applications) > 0 || s.HasField("applications") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("applications")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Applications {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("applications"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Applications to JSON.
func (x *Applications) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the Applications message from JSON.
func (x *Applications) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "applications":
			s.AddField("applications")
			if s.ReadNil() {
				x.Applications = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Applications = append(x.Applications, nil)
					return
				}
				v := &Application{}
				v.UnmarshalProtoJSON(s.WithField("applications", false))
				if s.Err() != nil {
					return
				}
				x.Applications = append(x.Applications, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the Applications from JSON.
func (x *Applications) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the IssueDevEUIResponse message to JSON.
func (x *IssueDevEUIResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.DevEui) > 0 || s.HasField("dev_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_eui")
		types.MarshalHEXBytes(s.WithField("dev_eui"), x.DevEui)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the IssueDevEUIResponse to JSON.
func (x *IssueDevEUIResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the IssueDevEUIResponse message from JSON.
func (x *IssueDevEUIResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "dev_eui", "devEui":
			s.AddField("dev_eui")
			x.DevEui = types.Unmarshal8Bytes(s.WithField("dev_eui", false))
		}
	})
}

// UnmarshalJSON unmarshals the IssueDevEUIResponse from JSON.
func (x *IssueDevEUIResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetApplicationRequest message to JSON.
func (x *GetApplicationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ApplicationIds)
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

// MarshalJSON marshals the GetApplicationRequest to JSON.
func (x *GetApplicationRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetApplicationRequest message from JSON.
func (x *GetApplicationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
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

// UnmarshalJSON unmarshals the GetApplicationRequest from JSON.
func (x *GetApplicationRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ListApplicationsRequest message to JSON.
func (x *ListApplicationsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.Collaborator)
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

// MarshalJSON marshals the ListApplicationsRequest to JSON.
func (x *ListApplicationsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ListApplicationsRequest message from JSON.
func (x *ListApplicationsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "collaborator":
			s.AddField("collaborator")
			if s.ReadNil() {
				x.Collaborator = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.Collaborator = &v
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

// UnmarshalJSON unmarshals the ListApplicationsRequest from JSON.
func (x *ListApplicationsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CreateApplicationRequest message to JSON.
func (x *CreateApplicationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Application != nil || s.HasField("application") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application")
		x.Application.MarshalProtoJSON(s.WithField("application"))
	}
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.Collaborator)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CreateApplicationRequest to JSON.
func (x *CreateApplicationRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CreateApplicationRequest message from JSON.
func (x *CreateApplicationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application":
			if s.ReadNil() {
				x.Application = nil
				return
			}
			x.Application = &Application{}
			x.Application.UnmarshalProtoJSON(s.WithField("application", true))
		case "collaborator":
			s.AddField("collaborator")
			if s.ReadNil() {
				x.Collaborator = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.Collaborator = &v
		}
	})
}

// UnmarshalJSON unmarshals the CreateApplicationRequest from JSON.
func (x *CreateApplicationRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the UpdateApplicationRequest message to JSON.
func (x *UpdateApplicationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Application != nil || s.HasField("application") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application")
		x.Application.MarshalProtoJSON(s.WithField("application"))
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

// MarshalJSON marshals the UpdateApplicationRequest to JSON.
func (x *UpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the UpdateApplicationRequest message from JSON.
func (x *UpdateApplicationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application":
			if s.ReadNil() {
				x.Application = nil
				return
			}
			x.Application = &Application{}
			x.Application.UnmarshalProtoJSON(s.WithField("application", true))
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

// UnmarshalJSON unmarshals the UpdateApplicationRequest from JSON.
func (x *UpdateApplicationRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CreateApplicationAPIKeyRequest message to JSON.
func (x *CreateApplicationAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ApplicationIds)
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.ExpiresAt != nil || s.HasField("expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("expires_at")
		if x.ExpiresAt == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.ExpiresAt)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CreateApplicationAPIKeyRequest to JSON.
func (x *CreateApplicationAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CreateApplicationAPIKeyRequest message from JSON.
func (x *CreateApplicationAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "rights":
			s.AddField("rights")
			if s.ReadNil() {
				x.Rights = nil
				return
			}
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		case "expires_at", "expiresAt":
			s.AddField("expires_at")
			if s.ReadNil() {
				x.ExpiresAt = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.ExpiresAt = v
		}
	})
}

// UnmarshalJSON unmarshals the CreateApplicationAPIKeyRequest from JSON.
func (x *CreateApplicationAPIKeyRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the UpdateApplicationAPIKeyRequest message to JSON.
func (x *UpdateApplicationAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ApplicationIds)
	}
	if x.ApiKey != nil || s.HasField("api_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("api_key")
		x.ApiKey.MarshalProtoJSON(s.WithField("api_key"))
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

// MarshalJSON marshals the UpdateApplicationAPIKeyRequest to JSON.
func (x *UpdateApplicationAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the UpdateApplicationAPIKeyRequest message from JSON.
func (x *UpdateApplicationAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "api_key", "apiKey":
			if s.ReadNil() {
				x.ApiKey = nil
				return
			}
			x.ApiKey = &APIKey{}
			x.ApiKey.UnmarshalProtoJSON(s.WithField("api_key", true))
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

// UnmarshalJSON unmarshals the UpdateApplicationAPIKeyRequest from JSON.
func (x *UpdateApplicationAPIKeyRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the SetApplicationCollaboratorRequest message to JSON.
func (x *SetApplicationCollaboratorRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ApplicationIds)
	}
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		x.Collaborator.MarshalProtoJSON(s.WithField("collaborator"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SetApplicationCollaboratorRequest to JSON.
func (x *SetApplicationCollaboratorRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the SetApplicationCollaboratorRequest message from JSON.
func (x *SetApplicationCollaboratorRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "collaborator":
			if s.ReadNil() {
				x.Collaborator = nil
				return
			}
			x.Collaborator = &Collaborator{}
			x.Collaborator.UnmarshalProtoJSON(s.WithField("collaborator", true))
		}
	})
}

// UnmarshalJSON unmarshals the SetApplicationCollaboratorRequest from JSON.
func (x *SetApplicationCollaboratorRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}