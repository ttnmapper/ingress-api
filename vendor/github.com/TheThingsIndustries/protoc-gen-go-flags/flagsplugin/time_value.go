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

// NewTimestampFlag defines a new flag with time value.
func NewTimestampFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &TimestampValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetTimestamp returns a value from a time flag.
func GetTimestamp(fs *pflag.FlagSet, name string) (value time.Time, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return time.Time{}, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*TimestampValue).Value, flag.Changed, nil
}

// TimestampValue implements pflag.Value interface.
type TimestampValue struct {
	Value time.Time
}

// Set implements pflag.Value interface
func (tv *TimestampValue) Set(s string) error {
	v, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return err
	}
	tv.Value = v
	return err
}

// Type implements pflag.Value interface.
func (*TimestampValue) Type() string { return "timestamp" }

// String implements pflag.Value interface.
func (tv *TimestampValue) String() string {
	if tv.Value.IsZero() {
		return ""
	}
	return tv.Value.Format(time.RFC3339Nano)
}

// NewTimestampSliceFlag defines a new flag that holds a slice of time values.
func NewTimestampSliceFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:     name,
		Usage:    usage,
		Value:    &TimestampSliceValue{},
		DefValue: "0",
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetTimestampSlice returns a value from a time slice flag.
func GetTimestampSlice(fs *pflag.FlagSet, name string) (value []time.Time, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]time.Time, len(flag.Value.(*TimestampSliceValue).Values))
	for i, v := range flag.Value.(*TimestampSliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// TimestampSliceValue implements pflag.Value interface.
type TimestampSliceValue struct {
	Values []TimestampValue
}

// Set implements pflag.Value interface.
func (tsv *TimestampSliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var tv TimestampValue
		if err := tv.Set(v); err != nil {
			return err
		}
		tsv.Values = append(tsv.Values, tv)
	}
	return nil
}

// Type implements pflag.Value interface.
func (*TimestampSliceValue) Type() string { return "timestampSlice" }

// String implements pflag.Value interface.
func (tsv *TimestampSliceValue) String() string {
	if len(tsv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(tsv.Values))
	for i, v := range tsv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

// NewStringTimestampMapFlag defines a new flag that holds a map of string to time.
func NewStringTimestampMapFlag(name, usage string, opts ...FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringTimestampMapValue{},
	}
	ApplyOptions(flag, opts...)
	return flag
}

// GetStringTimestampMap returns a value from a string to time map flag.
func GetStringTimestampMap(fs *pflag.FlagSet, name string) (value map[string]time.Time, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]time.Time, len(flag.Value.(*StringTimestampMapValue).Values))
	for i, v := range flag.Value.(*StringTimestampMapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// StringTimestampMapValue implements pflag.Value interface.
type StringTimestampMapValue struct {
	Values map[string]TimestampValue
}

// Set implements pflag.Value interface.
func (stmv *StringTimestampMapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv TimestampValue
		if err := fv.Set(v); err != nil {
			return err
		}
		if stmv.Values == nil {
			stmv.Values = make(map[string]TimestampValue)
		}
		stmv.Values[k] = fv
	}
	return nil
}

// Type implements pflag.Value interface.
func (*StringTimestampMapValue) Type() string { return "stringToTimestamp" }

// String implements pflag.Value interface.
func (stmv *StringTimestampMapValue) String() string {
	if len(stmv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(stmv.Values))
	for k := range stmv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(stmv.Values))
	for _, k := range ks {
		v := stmv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
