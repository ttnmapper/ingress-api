// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3SetApplicationLinkRequest v3 set application link request
//
// swagger:model v3SetApplicationLinkRequest
type V3SetApplicationLinkRequest struct {

	// application ids
	ApplicationIds *V3ApplicationIdentifiers `json:"application_ids,omitempty"`

	// field mask
	FieldMask *ProtobufFieldMask `json:"field_mask,omitempty"`

	// link
	Link *V3ApplicationLink `json:"link,omitempty"`
}

// Validate validates this v3 set application link request
func (m *V3SetApplicationLinkRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldMask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3SetApplicationLinkRequest) validateApplicationIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicationIds) { // not required
		return nil
	}

	if m.ApplicationIds != nil {
		if err := m.ApplicationIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application_ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3SetApplicationLinkRequest) validateFieldMask(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldMask) { // not required
		return nil
	}

	if m.FieldMask != nil {
		if err := m.FieldMask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("field_mask")
			}
			return err
		}
	}

	return nil
}

func (m *V3SetApplicationLinkRequest) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3SetApplicationLinkRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3SetApplicationLinkRequest) UnmarshalBinary(b []byte) error {
	var res V3SetApplicationLinkRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}