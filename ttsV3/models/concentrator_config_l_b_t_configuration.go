// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConcentratorConfigLBTConfiguration concentrator config l b t configuration
//
// swagger:model ConcentratorConfigLBTConfiguration
type ConcentratorConfigLBTConfiguration struct {

	// Received signal strength offset (dBm).
	RssiOffset float32 `json:"rssi_offset,omitempty"`

	// Received signal strength (dBm).
	RssiTarget float32 `json:"rssi_target,omitempty"`

	// scan time
	ScanTime string `json:"scan_time,omitempty"`
}

// Validate validates this concentrator config l b t configuration
func (m *ConcentratorConfigLBTConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConcentratorConfigLBTConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConcentratorConfigLBTConfiguration) UnmarshalBinary(b []byte) error {
	var res ConcentratorConfigLBTConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}