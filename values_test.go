// Copyright (c) 2024  The Go-Enjin Authors
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

package values

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type stringerType struct {
	value []byte
}

func (s *stringerType) String() string {
	return "[" + string(s.value) + "]"
}

func Test(t *testing.T) {
	Convey("ToString", t, func() {
		Convey("is a stringer type", func() {
			v := &stringerType{value: []byte("this is a test")}
			So(ToString(v), ShouldEqual, "[this is a test]")
		})
		Convey("is not a stringer type", func() {
			v := 10
			So(ToString(v), ShouldEqual, "10")
		})
	})

	Convey("IsEmpty", t, func() {
		Convey("is actually empty", func() {
			So(IsEmpty(0), ShouldEqual, true)
			So(IsEmpty(""), ShouldEqual, true)
			So(IsEmpty(nil), ShouldEqual, true)
		})

		Convey("is not empty", func() {
			So(IsEmpty(-1), ShouldEqual, false)
			So(IsEmpty(1), ShouldEqual, false)
			So(IsEmpty("1"), ShouldEqual, false)
			So(IsEmpty(&stringerType{}), ShouldEqual, false)
			So(IsEmpty(struct{}{}), ShouldEqual, false)
		})
	})

	Convey("TypeOf", t, func() {
		So(TypeOf(""), ShouldEqual, "string")
		So(TypeOf(10), ShouldEqual, "int")
		So(TypeOf(2.5), ShouldEqual, "float64")
		So(TypeOf(&stringerType{}), ShouldEqual, "*values.stringerType")
	})

	Convey("GetKeyedValue", t, func() {
		m := map[string]interface{}{
			"one": 2,
		}
		v, ok := GetKeyedValue("one", m)
		So(ok, ShouldBeTrue)
		So(v.IsValid(), ShouldBeTrue)
		v, ok = GetKeyedValue("two", m)
		So(ok, ShouldBeFalse)
		So(v.IsValid(), ShouldBeFalse)

		s := &struct {
			One int
		}{One: 2}
		v, ok = GetKeyedValue("One", s)
		So(ok, ShouldBeTrue)
		So(v.IsValid(), ShouldBeTrue)
		v, ok = GetKeyedValue("Two", s)
		So(ok, ShouldBeFalse)
		So(v.IsValid(), ShouldBeFalse)
		v, ok = GetKeyedValue("One", *s)
		So(ok, ShouldBeTrue)
		So(v.IsValid(), ShouldBeTrue)
		v, ok = GetKeyedValue("Two", *s)
		So(ok, ShouldBeFalse)
		So(v.IsValid(), ShouldBeFalse)
	})

	Convey("GetKeyedType", t, func() {
		m := map[string]interface{}{
			"one": true,
		}
		v, ok := GetKeyedType(reflect.Bool, "one", m)
		So(ok, ShouldBeTrue)
		So(v.IsValid(), ShouldBeTrue)
		v, ok = GetKeyedType(reflect.Bool, "two", m)
		So(ok, ShouldBeFalse)
		So(v.IsValid(), ShouldBeFalse)
	})

	Convey("GetKeyedBool", t, func() {
		m := map[string]interface{}{
			"one": true,
		}
		v, ok := GetKeyedBool("one", m)
		So(ok, ShouldBeTrue)
		So(v, ShouldBeTrue)
		v, ok = GetKeyedBool("two", m)
		So(ok, ShouldBeFalse)
		So(v, ShouldBeFalse)
	})
}
