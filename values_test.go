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
}
