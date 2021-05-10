// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"fmt"
	"reflect"
)

type attributer interface {
	Attributes() map[string]interface{}
}

type publicAttributer interface {
	PublicAttributes() map[string]interface{}
}

var errOddKV = DefineInvalidArgument("odd_kv", "Odd number of key-value elements")

func supported(v interface{}) interface{} {
	if v == nil {
		return "<nil>"
	}
	rv := reflect.Indirect(reflect.ValueOf(v))
	switch rv.Type().Kind() {
	case reflect.Int:
		return int(rv.Int())
	case reflect.Float64:
		return rv.Float()
	case reflect.String:
		return rv.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func kvToMap(kv ...interface{}) (map[string]interface{}, error) {
	if len(kv)%2 != 0 {
		return nil, errOddKV.New()
	}
	m := make(map[string]interface{}, len(kv)/2)
	var key string
	for i, node := range kv {
		if i%2 == 0 {
			key = fmt.Sprintf("%v", node)
		} else {
			m[key] = node
		}
	}
	return m, nil
}

func (e *Error) mergeAttributes(kv ...interface{}) {
	if len(kv) == 0 {
		return
	}
	attributes, err := kvToMap(kv...)
	if err != nil {
		panic(err)
	}

	if e.attributes != nil {
		// Merge with existing attributes:
		for k, v := range e.attributes {
			if _, ok := attributes[k]; !ok {
				attributes[k] = v
			}
		}
	}
	e.attributes = attributes

	// Ensure that all attributes used as message format arguments are supported:
	for k, v := range e.attributes {
		for _, arg := range e.messageFormatArguments {
			if k == arg {
				e.attributes[k] = supported(v)
			}
		}
	}

	// Set the call stack if not already set:
	if e.stack == nil {
		e.stack = callers(4)
	}

	e.clearGRPCStatus()
}

// WithAttributes returns the error with the given attributes set.
// Any conflicting attributes in the Error will be overwritten.
func (e Error) WithAttributes(kv ...interface{}) Error {
	e.mergeAttributes(kv...)
	return e
}

// WithAttributes returns a new error from the definition, and sets the given attributes.
func (d Definition) WithAttributes(kv ...interface{}) Error {
	e := build(d, 0) // Don't refactor this to build(...).WithAttributes(...)
	e.mergeAttributes(kv...)
	return e
}

// Attributes of the error.
func (e Error) Attributes() map[string]interface{} { return e.attributes }

// PublicAttributes of the error.
func (e Error) PublicAttributes() map[string]interface{} {
	if len(e.attributes) == 0 {
		return nil
	}
	publicAttributes := make(map[string]interface{}, len(e.attributes))
nextAttr:
	for k, v := range e.attributes {
		for _, public := range e.publicAttributes {
			if k == public {
				publicAttributes[k] = v
				continue nextAttr
			}
		}
	}
	if len(publicAttributes) == 0 {
		return nil
	}
	return publicAttributes
}

// Attributes are not present in the error definition, so this just returns nil.
func (d Definition) Attributes() map[string]interface{} { return nil }

// PublicAttributes are not present in the error definition, so this just returns nil.
func (d Definition) PublicAttributes() map[string]interface{} { return nil }

// Attributes returns the attributes of the errors, if they implement Attributes().
// If more than one error is passed, subsequent error attributes will be added if not set.
func Attributes(err ...error) map[string]interface{} {
	attributes := make(map[string]interface{})
	for _, err := range err {
		if err, ok := err.(attributer); ok {
			for k, v := range err.Attributes() {
				if _, ok := attributes[k]; !ok {
					attributes[k] = v
				}
			}
		}
	}
	return attributes
}

// PublicAttributes returns the public attributes of the errors, if they implement PublicAttributes().
// If more than one error is passed, subsequent error attributes will be added if not set.
func PublicAttributes(err ...error) map[string]interface{} {
	attributes := make(map[string]interface{})
	for _, err := range err {
		if err, ok := err.(publicAttributer); ok {
			for k, v := range err.PublicAttributes() {
				if _, ok := attributes[k]; !ok {
					attributes[k] = v
				}
			}
		}
	}
	return attributes
}
