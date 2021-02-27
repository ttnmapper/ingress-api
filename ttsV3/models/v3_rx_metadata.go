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

// V3RxMetadata Contains metadata for a received message. Each antenna that receives
// a message corresponds to one RxMetadata.
//
// swagger:model v3RxMetadata
type V3RxMetadata struct {

	// Advanced metadata fields
	// - can be used for advanced information or experimental features that are not yet formally defined in the API
	// - field names are written in snake_case
	Advanced interface{} `json:"advanced,omitempty"`

	// antenna index
	AntennaIndex int64 `json:"antenna_index,omitempty"`

	// Index of the gateway channel that received the message.
	ChannelIndex int64 `json:"channel_index,omitempty"`

	// Received signal strength indicator of the channel (dBm).
	ChannelRssi float32 `json:"channel_rssi,omitempty"`

	// Gateway downlink path constraint; injected by the Gateway Server.
	DownlinkPathConstraint V3DownlinkPathConstraint `json:"downlink_path_constraint,omitempty"`

	// Encrypted gateway's internal fine timestamp when the Rx finished (nanoseconds).
	// Format: byte
	EncryptedFineTimestamp strfmt.Base64 `json:"encrypted_fine_timestamp,omitempty"`

	// encrypted fine timestamp key id
	EncryptedFineTimestampKeyID string `json:"encrypted_fine_timestamp_key_id,omitempty"`

	// Gateway's internal fine timestamp when the Rx finished (nanoseconds).
	FineTimestamp string `json:"fine_timestamp,omitempty"`

	// Frequency offset (Hz).
	FrequencyOffset string `json:"frequency_offset,omitempty"`

	// gateway ids
	GatewayIds *V3GatewayIdentifiers `json:"gateway_ids,omitempty"`

	// Antenna location; injected by the Gateway Server.
	Location *Lorawanv3Location `json:"location,omitempty"`

	// Received signal strength indicator (dBm).
	// This value equals `channel_rssi`.
	Rssi float32 `json:"rssi,omitempty"`

	// Standard deviation of the RSSI during preamble.
	RssiStandardDeviation float32 `json:"rssi_standard_deviation,omitempty"`

	// Received signal strength indicator of the signal (dBm).
	SignalRssi float32 `json:"signal_rssi,omitempty"`

	// Signal-to-noise ratio (dB).
	Snr float32 `json:"snr,omitempty"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// Gateway concentrator timestamp when the Rx finished (microseconds).
	Timestamp int64 `json:"timestamp,omitempty"`

	// Uplink token to be included in the Tx request in class A downlink; injected by gateway, Gateway Server or fNS.
	// Format: byte
	UplinkToken strfmt.Base64 `json:"uplink_token,omitempty"`
}

// Validate validates this v3 rx metadata
func (m *V3RxMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownlinkPathConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3RxMetadata) validateDownlinkPathConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkPathConstraint) { // not required
		return nil
	}

	if err := m.DownlinkPathConstraint.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("downlink_path_constraint")
		}
		return err
	}

	return nil
}

func (m *V3RxMetadata) validateGatewayIds(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayIds) { // not required
		return nil
	}

	if m.GatewayIds != nil {
		if err := m.GatewayIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway_ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3RxMetadata) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *V3RxMetadata) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3RxMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3RxMetadata) UnmarshalBinary(b []byte) error {
	var res V3RxMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}