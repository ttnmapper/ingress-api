// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MACCommandLinkCheckAns m a c command link check ans
//
// swagger:model MACCommandLinkCheckAns
type MACCommandLinkCheckAns struct {

	// gateway count
	GatewayCount int64 `json:"gateway_count,omitempty"`

	// Indicates the link margin in dB of the received LinkCheckReq, relative to the demodulation floor.
	Margin int64 `json:"margin,omitempty"`
}

// Validate validates this m a c command link check ans
func (m *MACCommandLinkCheckAns) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MACCommandLinkCheckAns) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MACCommandLinkCheckAns) UnmarshalBinary(b []byte) error {
	var res MACCommandLinkCheckAns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}