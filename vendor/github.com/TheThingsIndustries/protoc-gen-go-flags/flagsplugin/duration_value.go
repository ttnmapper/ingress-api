// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

// NewDurationFlag defines a new flag with bool value.
func NewDurationFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &DurationValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetDuration returns a value from a duration flag.
func GetDuration(fs *pflag.FlagSet, name string) (value time.Duration, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*DurationValue).Value, flag.Changed, nil
}

// DurationValue implements pflag.Value interface.
type DurationValue struct {
	Value time.Duration
}

// Set implements pflag.Value interface
func (dv *DurationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	dv.Value = v
	return err
}

// Type implements pflag.Value interface.
func (*DurationValue) Type() string { return "duration" }

// String implements pflag.Value interface.
func (dv *DurationValue) String() string {
	if dv.Value == 0 {
		return ""
	}
	return dv.Value.String()
}

// NewDurationSliceFlag defines a new flag with bool slice value.
func NewDurationSliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:     name,
		Usage:    usage,
		Value:    &DurationSliceValue{},
		DefValue: "0",
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetDurationSlice returns a value from a duration slice flag.
func GetDurationSlice(fs *pflag.FlagSet, name string) (value []time.Duration, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]time.Duration, len(flag.Value.(*DurationSliceValue).Values))
	for i, v := range flag.Value.(*DurationSliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// DurationSliceValue implements pflag.Value interface.
type DurationSliceValue struct {
	Values []DurationValue
}

// Set implements pflag.Value interface
func (dsv *DurationSliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var dv DurationValue
		if err := dv.Set(v); err != nil {
			return err
		}
		dsv.Values = append(dsv.Values, dv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (*DurationSliceValue) Type() string { return "durationSlice" }

func (dsv *DurationSliceValue) String() string {
	if len(dsv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(dsv.Values))
	for i, v := range dsv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringDurationMapFlag defines a new flag with string to duration map value.
func NewStringDurationMapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringDurationMapValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringDurationMap returns a value from a string to duration map flag.
func GetStringDurationMap(fs *pflag.FlagSet, name string) (value map[string]time.Duration, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]time.Duration, len(flag.Value.(*StringDurationMapValue).Values))
	for i, v := range flag.Value.(*StringDurationMapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringDurationMapValue implements pflag.Value interface.
type StringDurationMapValue struct {
	Values map[string]DurationValue
}

// Set implements pflag.Value interface
func (sdmv *StringDurationMapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv DurationValue
		if err := fv.Set(v); err != nil {
			return err
		}
		if sdmv.Values == nil {
			sdmv.Values = make(map[string]DurationValue)
		}
		sdmv.Values[k] = fv
	}
	return nil
}

// Type implements pflag.Value interface.
func (*StringDurationMapValue) Type() string { return "stringToDuration" }

// String implements pflag.Value interface.
func (sdmv *StringDurationMapValue) String() string {
	if len(sdmv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(sdmv.Values))
	for k := range sdmv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(sdmv.Values))
	for _, k := range ks {
		v := sdmv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
