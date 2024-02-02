[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/values)
[![codecov](https://codecov.io/gh/go-corelibs/values/graph/badge.svg?token=FwmMUOrU4r)](https://codecov.io/gh/go-corelibs/values)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/values)](https://goreportcard.com/report/github.com/go-corelibs/values)

# values - arbitrary variable type utilities

A collection of utilities for working with `interface{}` type things.

# Installation

``` shell
> go get github.com/go-corelibs/values@latest
```

# Examples

## IsEmpty

``` go
func main() {
    empty := values.IsEmpty("") // empty == true
    empty = values.IsEmpty(0) // empty == true
    empty = values.IsEmpty(nil) // empty == true
    empty = values.IsEmpty(-1) // empty == false
    empty = values.IsEmpty(&struct{}{}) // empty == false
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
