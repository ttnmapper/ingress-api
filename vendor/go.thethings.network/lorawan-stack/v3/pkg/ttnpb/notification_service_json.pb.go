// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.2
// - protoc             v3.21.1
// source: lorawan-stack/api/notification_service.proto

package ttnpb

import (
	golang "github.com/TheThingsIndustries/protoc-gen-go-json/golang"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the NotificationReceiver to JSON.
func (x NotificationReceiver) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), NotificationReceiver_name)
}

// MarshalText marshals the NotificationReceiver to text.
func (x NotificationReceiver) MarshalText() ([]byte, error) {
	return []byte(jsonplugin.GetEnumString(int32(x), NotificationReceiver_name)), nil
}

// MarshalJSON marshals the NotificationReceiver to JSON.
func (x NotificationReceiver) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// NotificationReceiver_customvalue contains custom string values that extend NotificationReceiver_value.
var NotificationReceiver_customvalue = map[string]int32{
	"UNKNOWN":                0,
	"COLLABORATOR":           1,
	"ADMINISTRATIVE_CONTACT": 3,
	"TECHNICAL_CONTACT":      4,
}

// UnmarshalProtoJSON unmarshals the NotificationReceiver from JSON.
func (x *NotificationReceiver) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(NotificationReceiver_value, NotificationReceiver_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read NotificationReceiver enum: %v", err)
		return
	}
	*x = NotificationReceiver(v)
}

// UnmarshalText unmarshals the NotificationReceiver from text.
func (x *NotificationReceiver) UnmarshalText(b []byte) error {
	i, err := jsonplugin.ParseEnumString(string(b), NotificationReceiver_customvalue, NotificationReceiver_value)
	if err != nil {
		return err
	}
	*x = NotificationReceiver(i)
	return nil
}

// UnmarshalJSON unmarshals the NotificationReceiver from JSON.
func (x *NotificationReceiver) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the NotificationStatus to JSON.
func (x NotificationStatus) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), NotificationStatus_name)
}

// MarshalText marshals the NotificationStatus to text.
func (x NotificationStatus) MarshalText() ([]byte, error) {
	return []byte(jsonplugin.GetEnumString(int32(x), NotificationStatus_name)), nil
}

// MarshalJSON marshals the NotificationStatus to JSON.
func (x NotificationStatus) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// NotificationStatus_customvalue contains custom string values that extend NotificationStatus_value.
var NotificationStatus_customvalue = map[string]int32{
	"UNSEEN":   0,
	"SEEN":     1,
	"ARCHIVED": 2,
}

// UnmarshalProtoJSON unmarshals the NotificationStatus from JSON.
func (x *NotificationStatus) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(NotificationStatus_value, NotificationStatus_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read NotificationStatus enum: %v", err)
		return
	}
	*x = NotificationStatus(v)
}

// UnmarshalText unmarshals the NotificationStatus from text.
func (x *NotificationStatus) UnmarshalText(b []byte) error {
	i, err := jsonplugin.ParseEnumString(string(b), NotificationStatus_customvalue, NotificationStatus_value)
	if err != nil {
		return err
	}
	*x = NotificationStatus(i)
	return nil
}

// UnmarshalJSON unmarshals the NotificationStatus from JSON.
func (x *NotificationStatus) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the Notification message to JSON.
func (x *Notification) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Id != "" || s.HasField("id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("id")
		s.WriteString(x.Id)
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
	if x.EntityIds != nil || s.HasField("entity_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("entity_ids")
		x.EntityIds.MarshalProtoJSON(s.WithField("entity_ids"))
	}
	if x.NotificationType != "" || s.HasField("notification_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("notification_type")
		s.WriteString(x.NotificationType)
	}
	if x.Data != nil || s.HasField("data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("data")
		if x.Data == nil {
			s.WriteNil()
		} else {
			golang.MarshalAny(s, x.Data, true)
		}
	}
	if x.SenderIds != nil || s.HasField("sender_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("sender_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.SenderIds)
	}
	if len(x.Receivers) > 0 || s.HasField("receivers") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("receivers")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Receivers {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.Email || s.HasField("email") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("email")
		s.WriteBool(x.Email)
	}
	if x.Status != 0 || s.HasField("status") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("status")
		x.Status.MarshalProtoJSON(s)
	}
	if x.StatusUpdatedAt != nil || s.HasField("status_updated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("status_updated_at")
		if x.StatusUpdatedAt == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.StatusUpdatedAt)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Notification to JSON.
func (x *Notification) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the Notification message from JSON.
func (x *Notification) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "id":
			s.AddField("id")
			x.Id = s.ReadString()
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
		case "entity_ids", "entityIds":
			if s.ReadNil() {
				x.EntityIds = nil
				return
			}
			x.EntityIds = &EntityIdentifiers{}
			x.EntityIds.UnmarshalProtoJSON(s.WithField("entity_ids", true))
		case "notification_type", "notificationType":
			s.AddField("notification_type")
			x.NotificationType = s.ReadString()
		case "data":
			s.AddField("data")
			if s.ReadNil() {
				x.Data = nil
				return
			}
			v := golang.UnmarshalAny(s)
			if s.Err() != nil {
				return
			}
			x.Data = v
		case "sender_ids", "senderIds":
			s.AddField("sender_ids")
			if s.ReadNil() {
				x.SenderIds = nil
				return
			}
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.SenderIds = &v
		case "receivers":
			s.AddField("receivers")
			if s.ReadNil() {
				x.Receivers = nil
				return
			}
			s.ReadArray(func() {
				var v NotificationReceiver
				v.UnmarshalProtoJSON(s)
				x.Receivers = append(x.Receivers, v)
			})
		case "email":
			s.AddField("email")
			x.Email = s.ReadBool()
		case "status":
			s.AddField("status")
			x.Status.UnmarshalProtoJSON(s)
		case "status_updated_at", "statusUpdatedAt":
			s.AddField("status_updated_at")
			if s.ReadNil() {
				x.StatusUpdatedAt = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.StatusUpdatedAt = v
		}
	})
}

// UnmarshalJSON unmarshals the Notification from JSON.
func (x *Notification) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CreateNotificationRequest message to JSON.
func (x *CreateNotificationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.EntityIds != nil || s.HasField("entity_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("entity_ids")
		x.EntityIds.MarshalProtoJSON(s.WithField("entity_ids"))
	}
	if x.NotificationType != "" || s.HasField("notification_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("notification_type")
		s.WriteString(x.NotificationType)
	}
	if x.Data != nil || s.HasField("data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("data")
		if x.Data == nil {
			s.WriteNil()
		} else {
			golang.MarshalAny(s, x.Data, true)
		}
	}
	if x.SenderIds != nil || s.HasField("sender_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("sender_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.SenderIds)
	}
	if len(x.Receivers) > 0 || s.HasField("receivers") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("receivers")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Receivers {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.Email || s.HasField("email") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("email")
		s.WriteBool(x.Email)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CreateNotificationRequest to JSON.
func (x *CreateNotificationRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CreateNotificationRequest message from JSON.
func (x *CreateNotificationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "entity_ids", "entityIds":
			if s.ReadNil() {
				x.EntityIds = nil
				return
			}
			x.EntityIds = &EntityIdentifiers{}
			x.EntityIds.UnmarshalProtoJSON(s.WithField("entity_ids", true))
		case "notification_type", "notificationType":
			s.AddField("notification_type")
			x.NotificationType = s.ReadString()
		case "data":
			s.AddField("data")
			if s.ReadNil() {
				x.Data = nil
				return
			}
			v := golang.UnmarshalAny(s)
			if s.Err() != nil {
				return
			}
			x.Data = v
		case "sender_ids", "senderIds":
			s.AddField("sender_ids")
			if s.ReadNil() {
				x.SenderIds = nil
				return
			}
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.SenderIds = &v
		case "receivers":
			s.AddField("receivers")
			if s.ReadNil() {
				x.Receivers = nil
				return
			}
			s.ReadArray(func() {
				var v NotificationReceiver
				v.UnmarshalProtoJSON(s)
				x.Receivers = append(x.Receivers, v)
			})
		case "email":
			s.AddField("email")
			x.Email = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the CreateNotificationRequest from JSON.
func (x *CreateNotificationRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ListNotificationsRequest message to JSON.
func (x *ListNotificationsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ReceiverIds != nil || s.HasField("receiver_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("receiver_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ReceiverIds)
	}
	if len(x.Status) > 0 || s.HasField("status") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("status")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Status {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
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
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ListNotificationsRequest to JSON.
func (x *ListNotificationsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ListNotificationsRequest message from JSON.
func (x *ListNotificationsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "receiver_ids", "receiverIds":
			s.AddField("receiver_ids")
			if s.ReadNil() {
				x.ReceiverIds = nil
				return
			}
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ReceiverIds = &v
		case "status":
			s.AddField("status")
			if s.ReadNil() {
				x.Status = nil
				return
			}
			s.ReadArray(func() {
				var v NotificationStatus
				v.UnmarshalProtoJSON(s)
				x.Status = append(x.Status, v)
			})
		case "limit":
			s.AddField("limit")
			x.Limit = s.ReadUint32()
		case "page":
			s.AddField("page")
			x.Page = s.ReadUint32()
		}
	})
}

// UnmarshalJSON unmarshals the ListNotificationsRequest from JSON.
func (x *ListNotificationsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ListNotificationsResponse message to JSON.
func (x *ListNotificationsResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Notifications) > 0 || s.HasField("notifications") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("notifications")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Notifications {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("notifications"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ListNotificationsResponse to JSON.
func (x *ListNotificationsResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ListNotificationsResponse message from JSON.
func (x *ListNotificationsResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "notifications":
			s.AddField("notifications")
			if s.ReadNil() {
				x.Notifications = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Notifications = append(x.Notifications, nil)
					return
				}
				v := &Notification{}
				v.UnmarshalProtoJSON(s.WithField("notifications", false))
				if s.Err() != nil {
					return
				}
				x.Notifications = append(x.Notifications, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the ListNotificationsResponse from JSON.
func (x *ListNotificationsResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the UpdateNotificationStatusRequest message to JSON.
func (x *UpdateNotificationStatusRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ReceiverIds != nil || s.HasField("receiver_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("receiver_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ReceiverIds)
	}
	if len(x.Ids) > 0 || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		s.WriteStringArray(x.Ids)
	}
	if x.Status != 0 || s.HasField("status") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("status")
		x.Status.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the UpdateNotificationStatusRequest to JSON.
func (x *UpdateNotificationStatusRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the UpdateNotificationStatusRequest message from JSON.
func (x *UpdateNotificationStatusRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "receiver_ids", "receiverIds":
			s.AddField("receiver_ids")
			if s.ReadNil() {
				x.ReceiverIds = nil
				return
			}
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ReceiverIds = &v
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			x.Ids = s.ReadStringArray()
		case "status":
			s.AddField("status")
			x.Status.UnmarshalProtoJSON(s)
		}
	})
}

// UnmarshalJSON unmarshals the UpdateNotificationStatusRequest from JSON.
func (x *UpdateNotificationStatusRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the EntityStateChangedNotification message to JSON.
func (x *EntityStateChangedNotification) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.State != 0 || s.HasField("state") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state")
		x.State.MarshalProtoJSON(s)
	}
	if x.StateDescription != "" || s.HasField("state_description") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state_description")
		s.WriteString(x.StateDescription)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the EntityStateChangedNotification to JSON.
func (x *EntityStateChangedNotification) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the EntityStateChangedNotification message from JSON.
func (x *EntityStateChangedNotification) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "state":
			s.AddField("state")
			x.State.UnmarshalProtoJSON(s)
		case "state_description", "stateDescription":
			s.AddField("state_description")
			x.StateDescription = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the EntityStateChangedNotification from JSON.
func (x *EntityStateChangedNotification) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}