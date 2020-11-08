// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3APIKey v3 API key
//
// swagger:model v3APIKey
type V3APIKey struct {

	// Immutable and unique public identifier for the API key.
	// Generated by the Access Server.
	ID string `json:"id,omitempty"`

	// Immutable and unique secret value of the API key.
	// Generated by the Access Server.
	Key string `json:"key,omitempty"`

	// User-defined (friendly) name for the API key.
	Name string `json:"name,omitempty"`

	// Rights that are granted to this API key.
	Rights []V3Right `json:"rights"`
}

// Validate validates this v3 API key
func (m *V3APIKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3APIKey) validateRights(formats strfmt.Registry) error {

	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	for i := 0; i < len(m.Rights); i++ {

		if err := m.Rights[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3APIKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3APIKey) UnmarshalBinary(b []byte) error {
	var res V3APIKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}