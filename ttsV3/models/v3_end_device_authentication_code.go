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

// V3EndDeviceAuthenticationCode Authentication code for end devices.
//
// swagger:model v3EndDeviceAuthenticationCode
type V3EndDeviceAuthenticationCode struct {

	// valid from
	// Format: date-time
	ValidFrom strfmt.DateTime `json:"valid_from,omitempty"`

	// valid to
	// Format: date-time
	ValidTo strfmt.DateTime `json:"valid_to,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this v3 end device authentication code
func (m *V3EndDeviceAuthenticationCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3EndDeviceAuthenticationCode) validateValidFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("valid_from", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3EndDeviceAuthenticationCode) validateValidTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidTo) { // not required
		return nil
	}

	if err := validate.FormatOf("valid_to", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3EndDeviceAuthenticationCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3EndDeviceAuthenticationCode) UnmarshalBinary(b []byte) error {
	var res V3EndDeviceAuthenticationCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}