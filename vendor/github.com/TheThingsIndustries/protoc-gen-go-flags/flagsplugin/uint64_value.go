// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

// NewUint64Flag defines a new flag with uint64 value.
func NewUint64Flag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Uint64Value{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetUint64 returns a value from a uint64 flag.
func GetUint64(fs *pflag.FlagSet, name string) (value uint64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*Uint64Value).Value, flag.Changed, nil
}

// Uint64Value implements pflag.Value interface.
type Uint64Value struct {
	Value uint64
}

// Set implements pflag.Value interface
func (uv *Uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	uv.Value = v
	return err
}

// Type implements pflag.Value interface.
func (*Uint64Value) Type() string { return "uint64" }

// String implements pflag.Value interface.
func (uv *Uint64Value) String() string { return strconv.FormatUint(uv.Value, 10) }

// NewUint64SliceFlag defines a new flag that holds a slice of uint64 values.
func NewUint64SliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Uint64SliceValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetUint64Slice returns a value from a uint64 slice flag.
func GetUint64Slice(fs *pflag.FlagSet, name string) (value []uint64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]uint64, len(flag.Value.(*Uint64SliceValue).Values))
	for i, v := range flag.Value.(*Uint64SliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// Uint64SliceValue implements pflag.Value interface.
type Uint64SliceValue struct {
	Values []Uint64Value
}

// Set implements pflag.Value interface.
func (usv *Uint64SliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var uv Uint64Value
		if err := uv.Set(v); err != nil {
			return err
		}
		usv.Values = append(usv.Values, uv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (*Uint64SliceValue) Type() string { return "uint64Slice" }

// String implements pflag.Value interface.
func (usv *Uint64SliceValue) String() string {
	if len(usv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(usv.Values))
	for i, v := range usv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringUint64MapFlag defines a new flag that holds a map of string to uint64.
func NewStringUint64MapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringUint64MapValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringUint64Map returns a value from a string to uint64 map flag.
func GetStringUint64Map(fs *pflag.FlagSet, name string) (value map[string]uint64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]uint64, len(flag.Value.(*StringUint64MapValue).Values))
	for i, v := range flag.Value.(*StringUint64MapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringUint64MapValue implements pflag.Value interface.
type StringUint64MapValue struct {
	Values map[string]Uint64Value
}

// Set implements pflag.Value interface.
func (sumv *StringUint64MapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv Uint64Value
		if err := fv.Set(v); err != nil {
			return err
		}
		if sumv.Values == nil {
			sumv.Values = make(map[string]Uint64Value)
		}
		sumv.Values[k] = fv
	}
	return nil
}

// Type implements pflag.Value interface.
func (*StringUint64MapValue) Type() string { return "stringToUint64" }

// String implements pflag.Value interface.
func (sumv *StringUint64MapValue) String() string {
	if len(sumv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(sumv.Values))
	for k := range sumv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(sumv.Values))
	for _, k := range ks {
		v := sumv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
