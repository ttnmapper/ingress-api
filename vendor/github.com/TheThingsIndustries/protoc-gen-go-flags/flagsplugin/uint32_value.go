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

// NewUint32Flag defines a new flag with uint32 value.
func NewUint32Flag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Uint32Value{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetUint32 returns a value from a uint32 flag.
func GetUint32(fs *pflag.FlagSet, name string) (value uint32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*Uint32Value).Value, flag.Changed, nil
}

// Uint32Value implements pflag.Value interface.
type Uint32Value struct {
	Value uint32
}

// Set implements pflag.Value interface
func (uv *Uint32Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}
	uv.Value = uint32(v)
	return err
}

// Type implements pflag.Value interface.
func (*Uint32Value) Type() string { return "uint32" }

// String implements pflag.Value interface.
func (uv *Uint32Value) String() string { return strconv.FormatUint(uint64(uv.Value), 10) }

// NewUint32SliceFlag defines a new flag that holds a slice of uint32 values.
func NewUint32SliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Uint32SliceValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetUint32Slice returns a value from a uint32 slice flag.
func GetUint32Slice(fs *pflag.FlagSet, name string) (value []uint32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]uint32, len(flag.Value.(*Uint32SliceValue).Values))
	for i, v := range flag.Value.(*Uint32SliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// Uint32SliceValue implements pflag.Value interface.
type Uint32SliceValue struct {
	Values []Uint32Value
}

// Set implements pflag.Value interface.
func (usv *Uint32SliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var uv Uint32Value
		if err := uv.Set(v); err != nil {
			return err
		}
		usv.Values = append(usv.Values, uv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (*Uint32SliceValue) Type() string { return "uint32Slice" }

// String implements pflag.Value interface.
func (usv *Uint32SliceValue) String() string {
	if len(usv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(usv.Values))
	for i, v := range usv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringUint32MapFlag defines a new flag that holds a map of string to uint32.
func NewStringUint32MapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringUint32MapValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringUint32Map returns a value from a string to uint32 map flag.
func GetStringUint32Map(fs *pflag.FlagSet, name string) (value map[string]uint32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]uint32, len(flag.Value.(*StringUint32MapValue).Values))
	for i, v := range flag.Value.(*StringUint32MapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringUint32MapValue implements pflag.Value interface.
type StringUint32MapValue struct {
	Values map[string]Uint32Value
}

// Set implements pflag.Value interface.
func (sumv *StringUint32MapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv Uint32Value
		if err := fv.Set(v); err != nil {
			return err
		}
		if sumv.Values == nil {
			sumv.Values = make(map[string]Uint32Value)
		}
		sumv.Values[k] = fv
	}
	return nil
}

// Type implements pflag.Value interface.
func (*StringUint32MapValue) Type() string { return "stringToUint32" }

// String implements pflag.Value interface.
func (sumv *StringUint32MapValue) String() string {
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
