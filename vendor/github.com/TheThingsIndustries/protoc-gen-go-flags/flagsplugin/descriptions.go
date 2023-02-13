// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import (
	"fmt"
	"sort"
	"strings"
)

type enumWithStringValue struct {
	number int32
	name   string
}

type sortedEnumsWithStringValue []enumWithStringValue

func (a sortedEnumsWithStringValue) Len() int      { return len(a) }
func (a sortedEnumsWithStringValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortedEnumsWithStringValue) Less(i, j int) bool {
	if a[i].number == a[j].number {
		return a[i].name < a[j].name
	}

	return a[i].number < a[j].number
}

// EnumValueDesc returns a string with possible values for enum.
func EnumValueDesc(valueMaps ...map[string]int32) string {
	var all sortedEnumsWithStringValue
	for _, valueMap := range valueMaps {
		for name, number := range valueMap {
			all = append(all, enumWithStringValue{number: number, name: name})
		}
	}

	sort.Stable(all)
	stringValues := make([]string, len(all))
	for i, v := range all {
		stringValues[i] = v.name
	}
	return fmt.Sprintf("allowed values: %s", strings.Join(stringValues, ", "))
}

// SelectDesc returns a string with a description for select flags.
func SelectDesc(fieldName string, withSubFields bool) string {
	fieldName = toUnderscore.Replace(fieldName)
	if withSubFields {
		return fmt.Sprintf("select the %s field and all allowed sub-fields", fieldName)
	}
	return fmt.Sprintf("select the %s field", fieldName)
}
