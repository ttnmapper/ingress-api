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

// NewBoolFlag defines a new flag with bool value.
func NewBoolFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:        name,
		Usage:       usage,
		Value:       &BoolValue{},
		NoOptDefVal: "true",
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetBool returns a value from a bool flag.
func GetBool(fs *pflag.FlagSet, name string) (value bool, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return false, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*BoolValue).Value, flag.Changed, nil
}

// BoolValue implements pflag.Value interface.
type BoolValue struct {
	Value bool
}

// Set implements pflag.Value interface
func (bv *BoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	bv.Value = v
	return err
}

// Type implements pflag.Value interface.
func (*BoolValue) Type() string { return "bool" }

// String implements pflag.Value interface.
func (bv *BoolValue) String() string { return strconv.FormatBool(bv.Value) }

// NewBoolSliceFlag defines a new flag that holds a slice of bools.
func NewBoolSliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &BoolSliceValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetBoolSlice returns a value from a bool slice flag.
func GetBoolSlice(fs *pflag.FlagSet, name string) (value []bool, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]bool, len(flag.Value.(*BoolSliceValue).Values))
	for i, v := range flag.Value.(*BoolSliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// BoolSliceValue implements pflag.Value interface.
type BoolSliceValue struct {
	Values []BoolValue
}

// Set implements pflag.Value interface.
func (bsv *BoolSliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var bv BoolValue
		if err := bv.Set(v); err != nil {
			return err
		}
		bsv.Values = append(bsv.Values, bv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (*BoolSliceValue) Type() string { return "boolSlice" }

// String implements pflag.Value interface.
func (bsv *BoolSliceValue) String() string {
	if len(bsv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(bsv.Values))
	for i, v := range bsv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringBoolMapFlag defines a new flag that holds a map of string to bool.
func NewStringBoolMapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringBoolMapValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringBoolMap returns a value from a string to bool map flag.
func GetStringBoolMap(fs *pflag.FlagSet, name string) (value map[string]bool, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]bool, len(flag.Value.(*StringBoolMapValue).Values))
	for i, v := range flag.Value.(*StringBoolMapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringBoolMapValue implements pflag.Value interface.
type StringBoolMapValue struct {
	Values map[string]BoolValue
}

// Set implements pflag.Value interface.
func (sbmv *StringBoolMapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var bv BoolValue
		if err := bv.Set(v); err != nil {
			return err
		}
		if sbmv.Values == nil {
			sbmv.Values = make(map[string]BoolValue)
		}
		sbmv.Values[k] = bv
	}
	return nil
}

// Type implements pflag.Value interface.
func (*StringBoolMapValue) Type() string { return "stringToBool" }

// String implements pflag.Value interface.
func (sbmv *StringBoolMapValue) String() string {
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
