// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3ApplicationUp v3 application up
//
// swagger:model v3ApplicationUp
type V3ApplicationUp struct {

	// correlation ids
	CorrelationIds []string `json:"correlation_ids"`

	// downlink ack
	DownlinkAck *V3ApplicationDownlink `json:"downlink_ack,omitempty"`

	// downlink failed
	DownlinkFailed *V3ApplicationDownlinkFailed `json:"downlink_failed,omitempty"`

	// downlink nack
	DownlinkNack *V3ApplicationDownlink `json:"downlink_nack,omitempty"`

	// downlink queue invalidated
	DownlinkQueueInvalidated *V3ApplicationInvalidatedDownlinks `json:"downlink_queue_invalidated,omitempty"`

	// downlink queued
	DownlinkQueued *V3ApplicationDownlink `json:"downlink_queued,omitempty"`

	// downlink sent
	DownlinkSent *V3ApplicationDownlink `json:"downlink_sent,omitempty"`

	// end device ids
	EndDeviceIds *V3EndDeviceIdentifiers `json:"end_device_ids,omitempty"`

	// join accept
	JoinAccept *V3ApplicationJoinAccept `json:"join_accept,omitempty"`

	// location solved
	LocationSolved *V3ApplicationLocation `json:"location_solved,omitempty"`

	// Server time when the Application Server received the message.
	// Format: date-time
	ReceivedAt strfmt.DateTime `json:"received_at,omitempty"`

	// uplink message
	UplinkMessage *V3ApplicationUplink `json:"uplink_message,omitempty"`
}

// Validate validates this v3 application up
func (m *V3ApplicationUp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownlinkAck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownlinkFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownlinkNack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownlinkQueueInvalidated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownlinkQueued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownlinkSent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDeviceIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoinAccept(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationSolved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUplinkMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3ApplicationUp) validateDownlinkAck(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkAck) { // not required
		return nil
	}

	if m.DownlinkAck != nil {
		if err := m.DownlinkAck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downlink_ack")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateDownlinkFailed(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkFailed) { // not required
		return nil
	}

	if m.DownlinkFailed != nil {
		if err := m.DownlinkFailed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downlink_failed")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateDownlinkNack(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkNack) { // not required
		return nil
	}

	if m.DownlinkNack != nil {
		if err := m.DownlinkNack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downlink_nack")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateDownlinkQueueInvalidated(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkQueueInvalidated) { // not required
		return nil
	}

	if m.DownlinkQueueInvalidated != nil {
		if err := m.DownlinkQueueInvalidated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downlink_queue_invalidated")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateDownlinkQueued(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkQueued) { // not required
		return nil
	}

	if m.DownlinkQueued != nil {
		if err := m.DownlinkQueued.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downlink_queued")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateDownlinkSent(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkSent) { // not required
		return nil
	}

	if m.DownlinkSent != nil {
		if err := m.DownlinkSent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downlink_sent")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateEndDeviceIds(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDeviceIds) { // not required
		return nil
	}

	if m.EndDeviceIds != nil {
		if err := m.EndDeviceIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("end_device_ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateJoinAccept(formats strfmt.Registry) error {

	if swag.IsZero(m.JoinAccept) { // not required
		return nil
	}

	if m.JoinAccept != nil {
		if err := m.JoinAccept.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("join_accept")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateLocationSolved(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationSolved) { // not required
		return nil
	}

	if m.LocationSolved != nil {
		if err := m.LocationSolved.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location_solved")
			}
			return err
		}
	}

	return nil
}

func (m *V3ApplicationUp) validateReceivedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ReceivedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("received_at", "body", "date-time", m.ReceivedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3ApplicationUp) validateUplinkMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.UplinkMessage) { // not required
		return nil
	}

	if m.UplinkMessage != nil {
		if err := m.UplinkMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uplink_message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3ApplicationUp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ApplicationUp) UnmarshalBinary(b []byte) error {
	var res V3ApplicationUp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}