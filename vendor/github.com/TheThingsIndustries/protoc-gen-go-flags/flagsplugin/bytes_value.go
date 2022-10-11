// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
)

// NewBytesFlag defines a new flag with base64 bytes value.
func NewBytesFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	return NewBase64BytesFlag(name, usage, opts...)
}

const (
	base64Encoding = "base64"
	hexEncoding    = "hex"
)

// NewBase64BytesFlag defines a new flag with base64 encoded bytes value.
func NewBase64BytesFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &BytesValue{Encoding: base64Encoding},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// NewHexBytesFlag defines a new flag with hex encoded bytes value.
func NewHexBytesFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &BytesValue{Encoding: hexEncoding},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetBytes returns a value from a bytes flag.
func GetBytes(fs *pflag.FlagSet, name string) (value []byte, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*BytesValue).Value, flag.Changed, nil
}

// BytesValue implements pflag.Value interface.
type BytesValue struct {
	Encoding string
	Value    []byte
}

// Base64Replacer transforms the string to RawStdEncoding string
var Base64Replacer = strings.NewReplacer("_", "/", "-", "+")

// Set implements pflag.Value interface.
func (bv *BytesValue) Set(s string) error {
	switch bv.Encoding {
	case base64Encoding, "":
		s = strings.TrimRight(s, "=")
		s = Base64Replacer.Replace(s)
		v, err := base64.RawStdEncoding.DecodeString(s)
		if err != nil {
			return err
		}
		bv.Value = v
		return nil
	case hexEncoding:
		v, err := hex.DecodeString(s)
		if err != nil {
			return err
		}
		bv.Value = v
		return nil
	default:
		return fmt.Errorf("unknown bytes encoding: %q", bv.Encoding)
	}
}

// Type implements pflag.Value interface.
func (bv *BytesValue) Type() string {
	switch bv.Encoding {
	case base64Encoding, "":
		return "bytesBase64"
	case hexEncoding:
		return "bytesHex"
	default:
		return fmt.Sprintf("%s-bytes", bv.Encoding)
	}
}

// String implements pflag.Value interface.
func (bv *BytesValue) String() string {
	switch bv.Encoding {
	case base64Encoding, "":
		return base64.StdEncoding.EncodeToString(bv.Value)
	case hexEncoding:
		return hex.EncodeToString(bv.Value)
	default:
		return ""
	}
}

// NewBytesSliceFlag defines a new bytes slice flag.
func NewBytesSliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	return NewBase64BytesSliceFlag(name, usage, opts...)
}

// NewBase64BytesSliceFlag defines a new base64 bytes slice flag.
func NewBase64BytesSliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &BytesSliceValue{Encoding: base64Encoding},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// NewHexBytesSliceFlag defines a new hex bytes slice flag.
func NewHexBytesSliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &BytesSliceValue{Encoding: hexEncoding},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetBytesSlice returns a value from a byte slice flag.
func GetBytesSlice(fs *pflag.FlagSet, name string) (value [][]byte, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([][]byte, len(flag.Value.(*BytesSliceValue).Values))
	for i, v := range flag.Value.(*BytesSliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// BytesSliceValue implements pflag.Value interface.
type BytesSliceValue struct {
	Encoding string
	Values   []BytesValue
}

// Set implements pflag.Value interface.
func (bsv *BytesSliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		bv := BytesValue{Encoding: bsv.Encoding}
		if err := bv.Set(v); err != nil {
			return err
		}
		bsv.Values = append(bsv.Values, bv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (bsv *BytesSliceValue) Type() string {
	switch bsv.Encoding {
	case base64Encoding, "":
		return "bytesBase64"
	case hexEncoding:
		return "bytesHex"
	default:
		return fmt.Sprintf("%s-bytes", bsv.Encoding)
	}
}

// String implements pflag.Value interface.
func (bsv *BytesSliceValue) String() string {
	if len(bsv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(bsv.Values))
	for i, v := range bsv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringBytesMapFlag defines a new string to bytes map flag.
func NewStringBytesMapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	return NewStringBase64BytesMapFlag(name, usage, opts...)
}

// NewStringBase64BytesMapFlag defines a new string to base64 bytes map flag.
func NewStringBase64BytesMapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringBytesMapValue{Encoding: base64Encoding},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// NewStringHexBytesMapFlag defines a new string to hex bytes map flag.
func NewStringHexBytesMapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringBytesMapValue{Encoding: hexEncoding},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringBytesMap returns a string to bytes map value from the flag.
func GetStringBytesMap(fs *pflag.FlagSet, name string) (value map[string][]byte, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string][]byte, len(flag.Value.(*StringBytesMapValue).Values))
	for i, v := range flag.Value.(*StringBytesMapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringBytesMapValue implements pflag.Value interface.
type StringBytesMapValue struct {
	Encoding string
	Values   map[string]BytesValue
}

// Set implements pflag.Value interface.
func (sbmv *StringBytesMapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		bv := BytesValue{Encoding: sbmv.Encoding}
		if err := bv.Set(v); err != nil {
			return err
		}
		if sbmv.Values == nil {
			sbmv.Values = make(map[string]BytesValue)
		}
		sbmv.Values[k] = bv
	}
	return nil
}

// Type implements pflag.Value interface.
func (sbmv *StringBytesMapValue) Type() string {
	switch sbmv.Encoding {
	case base64Encoding, "":
		return "stringToBytesBase64"
	case hexEncoding:
		return "stringToBytesHex"
	default:
		return fmt.Sprintf("%s-bytes", sbmv.Encoding)
	}
}

// String implements pflag.Value interface.
func (sbmv *StringBytesMapValue) String() string {
	if len(sbmv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(sbmv.Values))
	for k := range sbmv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(sbmv.Values))
	for _, k := range ks {
		v := sbmv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
