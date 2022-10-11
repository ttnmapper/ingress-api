// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

// SetterFromFlags defines interface for setting proto struct fields from flags.
type SetterFromFlags interface {
	SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error)
}

// SetEnumString parses an enum from its string flag representation using the value maps.
// If none of the value maps contains a mapping for the string value,
// it attempts to parse the string as a numeric value.
func SetEnumString(v string, valueMaps ...map[string]int32) (int32, error) {
	for _, valueMap := range valueMaps {
		if x, ok := valueMap[v]; ok {
			return x, nil
		}
	}
	x, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(x), nil
}

// IsAnyPrefixSet returns true if any flag with a supplied prefix is set.
func IsAnyPrefixSet(fs *pflag.FlagSet, prefix string) (flagSet bool) {
	prefix = toDash.Replace(prefix)
	fs.VisitAll(func(flag *pflag.Flag) {
		if flag.Changed && strings.HasPrefix(flag.Name, prefix) {
			flagSet = true
		}
	})
	return flagSet
}

// Prefix returns a field name with prefix of the form {prefix}.{field}.
func Prefix(field, prefix string) string {
	if prefix == "" {
		return field
	}
	return fmt.Sprintf("%s.%s", prefix, field)
}

// ErrFlagNotFound defines error on flagset.Lookup with unknown flag name.
type ErrFlagNotFound struct {
	FlagName string
}

func (err *ErrFlagNotFound) Error() string { return fmt.Sprintf("flag %q not found", err.FlagName) }

const trimChars = `'" `

// SplitSliceElements returns a slice of strings from a single string with commas.
func SplitSliceElements(s string) ([]string, error) {
	r := csv.NewReader(bytes.NewBufferString(s))
	elements, err := r.Read()
	if err != nil {
		return nil, err
	}
	return elements, nil
}

// JoinSliceElements returns a comma separated string from a slice of strings.
func JoinSliceElements(elements []string) string {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.Write(elements)
	w.Flush()
	return strings.TrimSpace(buf.String())
}

func splitStringMapElements(s string) (map[string]string, error) {
	if split := strings.SplitN(s, "=", 2); len(split) == 2 && !strings.Contains(split[1], "=") {
		return map[string]string{
			split[0]: split[1],
		}, nil
	}
	r := csv.NewReader(bytes.NewBufferString(s))
	elements, err := r.Read()
	if err != nil {
		return nil, err
	}
	res := make(map[string]string, len(elements))
	for _, element := range elements {
		split := strings.SplitN(element, "=", 2)
		if len(split) != 2 {
			return nil, fmt.Errorf("map flag values must be formatted as key=value")
		}
		res[split[0]] = split[1]
	}
	return res, nil
}
