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