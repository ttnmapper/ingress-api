// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MACCommandLinkADRAns m a c command link a d r ans
//
// swagger:model MACCommandLinkADRAns
type MACCommandLinkADRAns struct {

	// channel mask ack
	ChannelMaskAck bool `json:"channel_mask_ack,omitempty"`

	// data rate index ack
	DataRateIndexAck bool `json:"data_rate_index_ack,omitempty"`

	// tx power index ack
	TxPowerIndexAck bool `json:"tx_power_index_ack,omitempty"`
}

// Validate validates this m a c command link a d r ans
func (m *MACCommandLinkADRAns) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MACCommandLinkADRAns) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MACCommandLinkADRAns) UnmarshalBinary(b []byte) error {
	var res MACCommandLinkADRAns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}