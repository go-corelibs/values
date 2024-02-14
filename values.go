// Copyright (c) 2024  The Go-CoreLibs Authors
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

// Package values provides arbitrary variable type utilities
package values

import (
	"fmt"
	"reflect"
)

// ToString returns the string representation of an arbitrary value
func ToString(value interface{}) (v string) {
	if s, ok := value.(fmt.Stringer); ok {
		v = s.String()
		return
	}
	v = fmt.Sprintf("%v", value)
	return
}

// IsEmpty return true if the arbitrary value is empty (or zero)
func IsEmpty(value interface{}) (empty bool) {
	if empty = reflect.ValueOf(&value).Elem().IsZero(); !empty {
		v := ToString(value)
		empty = v == "" || v == "0"
	}
	return
}

// TypeOf is a convenience wrapper around `fmt.Sprintf("%T", value)`
func TypeOf(value interface{}) (name string) {
	name = fmt.Sprintf("%T", value)
	return
}

// GetKeyedValue uses reflection to inspect the data given and if it's a
// struct or map, check for the named key and return [reflect.Value] and
// ok equals true if the value is a valid [reflect.Value]
func GetKeyedValue(name string, data interface{}) (value reflect.Value, ok bool) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var kind reflect.Kind
	if kind = v.Kind(); kind == reflect.Struct {
		value = v.FieldByName(name)
		ok = value.IsValid()
	} else if kind == reflect.Map {
		if reflect.TypeOf(data).Key().Kind() == reflect.String {
			if mv := v.MapIndex(reflect.ValueOf(name)); mv.IsValid() {
				value = reflect.ValueOf(mv.Interface())
				ok = value.IsValid()
			}
		}
	}
	return
}

// GetKeyedType uses GetKeyedValue and only returns the value if it is of the
// specified `kind`
func GetKeyedType(kind reflect.Kind, name string, data interface{}) (value reflect.Value, ok bool) {
	if v, valid := GetKeyedValue(name, data); valid {
		if ok = v.Kind() == kind; ok {
			value = v
		}
	}
	return
}

// GetKeyedBool uses GetKeyedType to retrieve a bool value
func GetKeyedBool(name string, data interface{}) (value bool, ok bool) {
	if v, valid := GetKeyedType(reflect.Bool, name, data); valid {
		value = v.Bool()
		ok = true
	}
	return
}
