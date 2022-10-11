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

// NewInt64Flag defines a new flag with int64 value.
func NewInt64Flag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Int64Value{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetInt64 returns a value from a int64 flag.
func GetInt64(fs *pflag.FlagSet, name string) (value int64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*Int64Value).Value, flag.Changed, nil
}

// Int64Value implements pflag.Value interface.
type Int64Value struct {
	Value int64
}

// Set implements pflag.Value interface
func (iv *Int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	iv.Value = v
	return err
}

// Type implements pflag.Value interface.
func (*Int64Value) Type() string { return "int64" }

// String implements pflag.Value interface.
func (iv *Int64Value) String() string { return strconv.FormatInt(iv.Value, 10) }

// NewInt64SliceFlag defines a new flag that holds a slice of int64 values.
func NewInt64SliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Int64SliceValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetInt64Slice returns a value from a int64 slice flag.
func GetInt64Slice(fs *pflag.FlagSet, name string) (value []int64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]int64, len(flag.Value.(*Int64SliceValue).Values))
	for i, v := range flag.Value.(*Int64SliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// Int64SliceValue implements pflag.Value interface.
type Int64SliceValue struct {
	Values []Int64Value
}

// Set implements pflag.Value interface.
func (isv *Int64SliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var iv Int64Value
		if err := iv.Set(v); err != nil {
			return err
		}
		isv.Values = append(isv.Values, iv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (*Int64SliceValue) Type() string { return "int64Slice" }

// String implements pflag.Value interface.
func (isv *Int64SliceValue) String() string {
	if len(isv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(isv.Values))
	for i, v := range isv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringInt64MapFlag defines a new flag that holds a map of string to int64.
func NewStringInt64MapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringInt64MapValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringInt64Map returns a value from a string to int64 map flag.
func GetStringInt64Map(fs *pflag.FlagSet, name string) (value map[string]int64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]int64, len(flag.Value.(*StringInt64MapValue).Values))
	for i, v := range flag.Value.(*StringInt64MapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringInt64MapValue implements pflag.Value interface.
type StringInt64MapValue struct {
	Values map[string]Int64Value
}

// Set implements pflag.Value interface.
func (simv *StringInt64MapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv Int64Value
		if err := fv.Set(v); err != nil {
			return err
		}
		if simv.Values == nil {
			simv.Values = make(map[string]Int64Value)
		}
		simv.Values[k] = fv
	}
	return nil
}

// Type implements pflag.Value interface.
func (*StringInt64MapValue) Type() string { return "stringToInt64" }

// String implements pflag.Value interface.
func (simv *StringInt64MapValue) String() string {
	if len(simv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(simv.Values))
	for k := range simv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(simv.Values))
	for _, k := range ks {
		v := simv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
