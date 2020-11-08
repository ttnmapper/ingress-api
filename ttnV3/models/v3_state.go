// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V3State State enum defines states that an entity can be in.
//
//  - STATE_REQUESTED: Denotes that the entity has been requested and is pending review by an admin.
//  - STATE_APPROVED: Denotes that the entity has been reviewed and approved by an admin.
//  - STATE_REJECTED: Denotes that the entity has been reviewed and rejected by an admin.
//  - STATE_FLAGGED: Denotes that the entity has been flagged and is pending review by an admin.
//  - STATE_SUSPENDED: Denotes that the entity has been reviewed and suspended by an admin.
//
// swagger:model v3State
type V3State string

const (

	// V3StateSTATEREQUESTED captures enum value "STATE_REQUESTED"
	V3StateSTATEREQUESTED V3State = "STATE_REQUESTED"

	// V3StateSTATEAPPROVED captures enum value "STATE_APPROVED"
	V3StateSTATEAPPROVED V3State = "STATE_APPROVED"

	// V3StateSTATEREJECTED captures enum value "STATE_REJECTED"
	V3StateSTATEREJECTED V3State = "STATE_REJECTED"

	// V3StateSTATEFLAGGED captures enum value "STATE_FLAGGED"
	V3StateSTATEFLAGGED V3State = "STATE_FLAGGED"

	// V3StateSTATESUSPENDED captures enum value "STATE_SUSPENDED"
	V3StateSTATESUSPENDED V3State = "STATE_SUSPENDED"
)

// for schema
var v3StateEnum []interface{}

func init() {
	var res []V3State
	if err := json.Unmarshal([]byte(`["STATE_REQUESTED","STATE_APPROVED","STATE_REJECTED","STATE_FLAGGED","STATE_SUSPENDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3StateEnum = append(v3StateEnum, v)
	}
}

func (m V3State) validateV3StateEnum(path, location string, value V3State) error {
	if err := validate.Enum(path, location, value, v3StateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 state
func (m V3State) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3StateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}