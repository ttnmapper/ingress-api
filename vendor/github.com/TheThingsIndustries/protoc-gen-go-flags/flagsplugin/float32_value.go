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

// NewFloat32Flag defines a new flag with float32 value.
func NewFloat32Flag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Float32Value{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetFloat32 returns a value from a float32 flag.
func GetFloat32(fs *pflag.FlagSet, name string) (value float32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*Float32Value).Value, flag.Changed, nil
}

// Float32Value implements pflag.Value interface.
type Float32Value struct {
	Value float32
}

// Set implements pflag.Value interface
func (fv *Float32Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	fv.Value = float32(v)
	return err
}

// Type implements pflag.Value interface.
func (*Float32Value) Type() string { return "float32" }

// String implements pflag.Value interface.
func (fv *Float32Value) String() string { return strconv.FormatFloat(float64(fv.Value), 'g', -1, 32) }

// NewFloat32SliceFlag defines a new flag that holds a slice of float32 values.
func NewFloat32SliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Float32SliceValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetFloat32Slice returns a value from a float32 slice flag.
func GetFloat32Slice(fs *pflag.FlagSet, name string) (value []float32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]float32, len(flag.Value.(*Float32SliceValue).Values))
	for i, v := range flag.Value.(*Float32SliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// Float32SliceValue implements pflag.Value interface.
type Float32SliceValue struct {
	Values []Float32Value
}

// Set implements pflag.Value interface.
func (fsv *Float32SliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var fv Float32Value
		if err := fv.Set(v); err != nil {
			return err
		}
		fsv.Values = append(fsv.Values, fv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (*Float32SliceValue) Type() string { return "float32Slice" }

// String implements pflag.Value interface.
func (fsv *Float32SliceValue) String() string {
	if len(fsv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(fsv.Values))
	for i, v := range fsv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringFloat32MapFlag defines a new flag that holds a map of string to float32.
func NewStringFloat32MapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringFloat32MapValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringFloat32Map returns a value from a string to float32 map flag.
func GetStringFloat32Map(fs *pflag.FlagSet, name string) (value map[string]float32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]float32, len(flag.Value.(*StringFloat32MapValue).Values))
	for i, v := range flag.Value.(*StringFloat32MapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringFloat32MapValue implements pflag.Value interface.
type StringFloat32MapValue struct {
	Values map[string]Float32Value
}

// Set implements pflag.Value interface.
func (sfmv *StringFloat32MapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv Float32Value
		if err := fv.Set(v); err != nil {
			return err
		}
		if sfmv.Values == nil {
			sfmv.Values = make(map[string]Float32Value)
		}
		sfmv.Values[k] = fv
	}
	return nil
}

// Type implements pflag.Value interface.
func (*StringFloat32MapValue) Type() string { return "stringToFloat32" }

// String implements pflag.Value interface.
func (sfmv *StringFloat32MapValue) String() string {
	if len(sfmv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(sfmv.Values))
	for k := range sfmv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(sfmv.Values))
	for _, k := range ks {
		v := sfmv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
