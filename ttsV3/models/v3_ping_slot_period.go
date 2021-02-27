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

// V3PingSlotPeriod v3 ping slot period
//
// swagger:model v3PingSlotPeriod
type V3PingSlotPeriod string

const (

	// V3PingSlotPeriodPINGEVERY1S captures enum value "PING_EVERY_1S"
	V3PingSlotPeriodPINGEVERY1S V3PingSlotPeriod = "PING_EVERY_1S"

	// V3PingSlotPeriodPINGEVERY2S captures enum value "PING_EVERY_2S"
	V3PingSlotPeriodPINGEVERY2S V3PingSlotPeriod = "PING_EVERY_2S"

	// V3PingSlotPeriodPINGEVERY4S captures enum value "PING_EVERY_4S"
	V3PingSlotPeriodPINGEVERY4S V3PingSlotPeriod = "PING_EVERY_4S"

	// V3PingSlotPeriodPINGEVERY8S captures enum value "PING_EVERY_8S"
	V3PingSlotPeriodPINGEVERY8S V3PingSlotPeriod = "PING_EVERY_8S"

	// V3PingSlotPeriodPINGEVERY16S captures enum value "PING_EVERY_16S"
	V3PingSlotPeriodPINGEVERY16S V3PingSlotPeriod = "PING_EVERY_16S"

	// V3PingSlotPeriodPINGEVERY32S captures enum value "PING_EVERY_32S"
	V3PingSlotPeriodPINGEVERY32S V3PingSlotPeriod = "PING_EVERY_32S"

	// V3PingSlotPeriodPINGEVERY64S captures enum value "PING_EVERY_64S"
	V3PingSlotPeriodPINGEVERY64S V3PingSlotPeriod = "PING_EVERY_64S"

	// V3PingSlotPeriodPINGEVERY128S captures enum value "PING_EVERY_128S"
	V3PingSlotPeriodPINGEVERY128S V3PingSlotPeriod = "PING_EVERY_128S"
)

// for schema
var v3PingSlotPeriodEnum []interface{}

func init() {
	var res []V3PingSlotPeriod
	if err := json.Unmarshal([]byte(`["PING_EVERY_1S","PING_EVERY_2S","PING_EVERY_4S","PING_EVERY_8S","PING_EVERY_16S","PING_EVERY_32S","PING_EVERY_64S","PING_EVERY_128S"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3PingSlotPeriodEnum = append(v3PingSlotPeriodEnum, v)
	}
}

func (m V3PingSlotPeriod) validateV3PingSlotPeriodEnum(path, location string, value V3PingSlotPeriod) error {
	if err := validate.Enum(path, location, value, v3PingSlotPeriodEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 ping slot period
func (m V3PingSlotPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3PingSlotPeriodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}