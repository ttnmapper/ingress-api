// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3RejoinRequestPayload v3 rejoin request payload
//
// swagger:model v3RejoinRequestPayload
type V3RejoinRequestPayload struct {

	// dev eui
	// Format: byte
	DevEui strfmt.Base64 `json:"dev_eui,omitempty"`

	// join eui
	// Format: byte
	JoinEui strfmt.Base64 `json:"join_eui,omitempty"`

	// net id
	// Format: byte
	NetID strfmt.Base64 `json:"net_id,omitempty"`

	// rejoin cnt
	RejoinCnt int64 `json:"rejoin_cnt,omitempty"`

	// rejoin type
	RejoinType V3RejoinType `json:"rejoin_type,omitempty"`
}

// Validate validates this v3 rejoin request payload
func (m *V3RejoinRequestPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRejoinType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3RejoinRequestPayload) validateRejoinType(formats strfmt.Registry) error {

	if swag.IsZero(m.RejoinType) { // not required
		return nil
	}

	if err := m.RejoinType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rejoin_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3RejoinRequestPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3RejoinRequestPayload) UnmarshalBinary(b []byte) error {
	var res V3RejoinRequestPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}