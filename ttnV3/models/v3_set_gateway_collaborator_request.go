// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3SetGatewayCollaboratorRequest v3 set gateway collaborator request
//
// swagger:model v3SetGatewayCollaboratorRequest
type V3SetGatewayCollaboratorRequest struct {

	// collaborator
	Collaborator *V3Collaborator `json:"collaborator,omitempty"`

	// gateway ids
	GatewayIds *V3GatewayIdentifiers `json:"gateway_ids,omitempty"`
}

// Validate validates this v3 set gateway collaborator request
func (m *V3SetGatewayCollaboratorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollaborator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3SetGatewayCollaboratorRequest) validateCollaborator(formats strfmt.Registry) error {

	if swag.IsZero(m.Collaborator) { // not required
		return nil
	}

	if m.Collaborator != nil {
		if err := m.Collaborator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collaborator")
			}
			return err
		}
	}

	return nil
}

func (m *V3SetGatewayCollaboratorRequest) validateGatewayIds(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V3SetGatewayCollaboratorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3SetGatewayCollaboratorRequest) UnmarshalBinary(b []byte) error {
	var res V3SetGatewayCollaboratorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}