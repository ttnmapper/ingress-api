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

// NewInt32Flag defines a new flag with int32 value.
func NewInt32Flag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Int32Value{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetInt32 returns a value from a int32 flag.
func GetInt32(fs *pflag.FlagSet, name string) (value int32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*Int32Value).Value, flag.Changed, nil
}

// Int32Value implements pflag.Value interface.
type Int32Value struct {
	Value int32
}

// Set implements pflag.Value interface
func (iv *Int32Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	iv.Value = int32(v)
	return err
}

// Type implements pflag.Value interface.
func (*Int32Value) Type() string { return "int32" }

// String implements pflag.Value interface.
func (iv *Int32Value) String() string { return strconv.FormatInt(int64(iv.Value), 10) }

// NewInt32SliceFlag defines a new flag that holds a slice of int32 values.
func NewInt32SliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Int32SliceValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetInt32Slice returns a value from a int32 slice flag.
func GetInt32Slice(fs *pflag.FlagSet, name string) (value []int32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]int32, len(flag.Value.(*Int32SliceValue).Values))
	for i, v := range flag.Value.(*Int32SliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// Int32SliceValue implements pflag.Value interface.
type Int32SliceValue struct {
	Values []Int32Value
}

// Set implements pflag.Value interface.
func (isv *Int32SliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var iv Int32Value
		if err := iv.Set(v); err != nil {
			return err
		}
		isv.Values = append(isv.Values, iv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (*Int32SliceValue) Type() string { return "int32Slice" }

// String implements pflag.Value interface.
func (isv *Int32SliceValue) String() string {
	if len(isv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(isv.Values))
	for i, v := range isv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringInt32MapFlag defines a new flag that holds a map of string to int32.
func NewStringInt32MapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringInt32MapValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringInt32Map returns a value from a string to int32 map flag.
func GetStringInt32Map(fs *pflag.FlagSet, name string) (value map[string]int32, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]int32, len(flag.Value.(*StringInt32MapValue).Values))
	for i, v := range flag.Value.(*StringInt32MapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringInt32MapValue implements pflag.Value interface.
type StringInt32MapValue struct {
	Values map[string]Int32Value
}

// Set implements pflag.Value interface.
func (simv *StringInt32MapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv Int32Value
		if err := fv.Set(v); err != nil {
			return err
		}
		if simv.Values == nil {
			simv.Values = make(map[string]Int32Value)
		}
		simv.Values[k] = fv
	}
	return nil
}

// Type implements pflag.Value interface.
func (*StringInt32MapValue) Type() string { return "stringToInt32" }

// String implements pflag.Value interface.
func (simv *StringInt32MapValue) String() string {
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
