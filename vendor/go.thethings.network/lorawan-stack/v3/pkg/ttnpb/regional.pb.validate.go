// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _regional_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on ConcentratorConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ConcentratorConfig) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ConcentratorConfigFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "channels":

			for idx, item := range m.GetChannels() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return ConcentratorConfigValidationError{
							field:  fmt.Sprintf("channels[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		case "lora_standard_channel":

			if v, ok := interface{}(m.GetLoraStandardChannel()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ConcentratorConfigValidationError{
						field:  "lora_standard_channel",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "fsk_channel":

			if v, ok := interface{}(m.GetFskChannel()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ConcentratorConfigValidationError{
						field:  "fsk_channel",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "lbt":

			if v, ok := interface{}(m.GetLbt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ConcentratorConfigValidationError{
						field:  "lbt",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "ping_slot":

			if v, ok := interface{}(m.GetPingSlot()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ConcentratorConfigValidationError{
						field:  "ping_slot",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "radios":

			for idx, item := range m.GetRadios() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return ConcentratorConfigValidationError{
							field:  fmt.Sprintf("radios[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		case "clock_source":
			// no validation rules for ClockSource
		default:
			return ConcentratorConfigValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ConcentratorConfigValidationError is the validation error returned by
// ConcentratorConfig.ValidateFields if the designated constraints aren't met.
type ConcentratorConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcentratorConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcentratorConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcentratorConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcentratorConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcentratorConfigValidationError) ErrorName() string {
	return "ConcentratorConfigValidationError"
}

// Error satisfies the builtin error interface
func (e ConcentratorConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcentratorConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcentratorConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcentratorConfigValidationError{}

// ValidateFields checks the field values on ConcentratorConfig_Channel with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ConcentratorConfig_Channel) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ConcentratorConfig_ChannelFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "frequency":
			// no validation rules for Frequency
		case "radio":
			// no validation rules for Radio
		default:
			return ConcentratorConfig_ChannelValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ConcentratorConfig_ChannelValidationError is the validation error returned
// by ConcentratorConfig_Channel.ValidateFields if the designated constraints
// aren't met.
type ConcentratorConfig_ChannelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcentratorConfig_ChannelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcentratorConfig_ChannelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcentratorConfig_ChannelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcentratorConfig_ChannelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcentratorConfig_ChannelValidationError) ErrorName() string {
	return "ConcentratorConfig_ChannelValidationError"
}

// Error satisfies the builtin error interface
func (e ConcentratorConfig_ChannelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcentratorConfig_Channel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcentratorConfig_ChannelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcentratorConfig_ChannelValidationError{}

// ValidateFields checks the field values on
// ConcentratorConfig_LoRaStandardChannel with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *ConcentratorConfig_LoRaStandardChannel) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ConcentratorConfig_LoRaStandardChannelFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "frequency":
			// no validation rules for Frequency
		case "radio":
			// no validation rules for Radio
		case "bandwidth":
			// no validation rules for Bandwidth
		case "spreading_factor":
			// no validation rules for SpreadingFactor
		default:
			return ConcentratorConfig_LoRaStandardChannelValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ConcentratorConfig_LoRaStandardChannelValidationError is the validation
// error returned by ConcentratorConfig_LoRaStandardChannel.ValidateFields if
// the designated constraints aren't met.
type ConcentratorConfig_LoRaStandardChannelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcentratorConfig_LoRaStandardChannelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcentratorConfig_LoRaStandardChannelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcentratorConfig_LoRaStandardChannelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcentratorConfig_LoRaStandardChannelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcentratorConfig_LoRaStandardChannelValidationError) ErrorName() string {
	return "ConcentratorConfig_LoRaStandardChannelValidationError"
}

// Error satisfies the builtin error interface
func (e ConcentratorConfig_LoRaStandardChannelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcentratorConfig_LoRaStandardChannel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcentratorConfig_LoRaStandardChannelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcentratorConfig_LoRaStandardChannelValidationError{}

// ValidateFields checks the field values on ConcentratorConfig_FSKChannel with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ConcentratorConfig_FSKChannel) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ConcentratorConfig_FSKChannelFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "frequency":
			// no validation rules for Frequency
		case "radio":
			// no validation rules for Radio
		default:
			return ConcentratorConfig_FSKChannelValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ConcentratorConfig_FSKChannelValidationError is the validation error
// returned by ConcentratorConfig_FSKChannel.ValidateFields if the designated
// constraints aren't met.
type ConcentratorConfig_FSKChannelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcentratorConfig_FSKChannelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcentratorConfig_FSKChannelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcentratorConfig_FSKChannelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcentratorConfig_FSKChannelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcentratorConfig_FSKChannelValidationError) ErrorName() string {
	return "ConcentratorConfig_FSKChannelValidationError"
}

// Error satisfies the builtin error interface
func (e ConcentratorConfig_FSKChannelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcentratorConfig_FSKChannel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcentratorConfig_FSKChannelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcentratorConfig_FSKChannelValidationError{}

// ValidateFields checks the field values on
// ConcentratorConfig_LBTConfiguration with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *ConcentratorConfig_LBTConfiguration) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ConcentratorConfig_LBTConfigurationFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "rssi_target":
			// no validation rules for RssiTarget
		case "rssi_offset":
			// no validation rules for RssiOffset
		case "scan_time":

			if v, ok := interface{}(m.GetScanTime()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ConcentratorConfig_LBTConfigurationValidationError{
						field:  "scan_time",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return ConcentratorConfig_LBTConfigurationValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ConcentratorConfig_LBTConfigurationValidationError is the validation error
// returned by ConcentratorConfig_LBTConfiguration.ValidateFields if the
// designated constraints aren't met.
type ConcentratorConfig_LBTConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcentratorConfig_LBTConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcentratorConfig_LBTConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcentratorConfig_LBTConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcentratorConfig_LBTConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcentratorConfig_LBTConfigurationValidationError) ErrorName() string {
	return "ConcentratorConfig_LBTConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e ConcentratorConfig_LBTConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcentratorConfig_LBTConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcentratorConfig_LBTConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcentratorConfig_LBTConfigurationValidationError{}
