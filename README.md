
Xtractr
=====

[![Go Reference](https://pkg.go.dev/badge/github.com/syke99/xtractr.svg)](https://pkg.go.dev/github.com/syke99/xtractr)
[![Go Reportcard](https://goreportcard.com/badge/github.com/syke99/xtractr)](https://goreportcard.com/report/github.com/syke99/xtractr)
[![codecov](https://codecov.io/gh/syke99/Xtractr/branch/main/graph/badge.svg?token=NAmOLNyTgw)](https://codecov.io/gh/syke99/Xtractr)
[![LICENSE](https://img.shields.io/github/license/syke99/xtractr)](https://pkg.go.dev/github.com/syke99/xtractr/blob/master/LICENSE)

Xtractr helps keep code clean whenever extracting path and query parameters from an incoming HTTP request


Why use Xtractr?
=====

Sometimes, depending on how complex an endpoint is, you could potentially have a large number of parameters. Go's `net/http` also doesn't allow
for very easy extraction of path parameters. And with the gorilla toolkit now being archived, one can't rely on using `gorilla/mux` to not eventually
break and introduce bugs and vulnerabilities. With Xtractr, you can tackle both of these issues with relative ease, and also keep your code base
nice and clean. Xtractr makes it simple to unmarshal these parameters directly to a struct by determining what parameter to look for for each struct
field via an `xtractr-param` tag, and where to find that parameter (in the path or query) via an `xtractr` tag

How do I use Xtractr?
=====

Using Xtractr couldn't be simpler. First, simply import Xtractr into your project by entering the following command in your terminal

```
go get github.com/syke99/xtractr
```

Then, simply import Xtractr into your desired package(s)

```go
import "github.com/syke99/xtractr"
```

After importing, define the struct you would like to unmarshal the parameters into. Example:

```go
type MyStruct struct {
	FieldOne    int    `xtractr:"path" xtractr-param:"fieldOne"`
	FieldTwo    string `xtractr:"query" xtractr-param:"fieldTwo"`
}
```

You can even nest structs if you'd like, just simply give the `xtractr` tag the `struct` value and Xtractr 
will recurse into that struct and unmarshal its defined parameters.

```go
type MyFirstStruct struct {
    FieldOne    int    `xtractr:"path" xtractr-param:"fieldOne"`
    FieldTwo    string `xtractr:"query" xtractr-param:"fieldTwo"`
}

type MySecondStruc struct {
    FieldThree    int           `xtractr:"path" xtractr-param:"fieldThree"`
    FieldFour     MyFirstStruct `xtractr:"struct"`
}
```

Xtractr also suports time.Time (and sql.NullTime) fields. You just need to provide an additional `xtractr-time`
tag to specify the format to be used (Xtractr defaults to Go's default time format):

```go
type MyStruct struct {
    TimeField    time.Time `xtractr:"path" xtractr-param:"timeField" xtractr-time:"ISO80601"` // xtractr supports this (YYYY-MM-DD) ISO8601 format
}
```

You can also use any SQL NullType (besides sql.NullByte) fields, just specify by appending `,sql` to the end
of the value you supply in the `xtractr` tag

```go
type MyStruct struct {
    FieldOne      sql.NullTime   `xtractr:"path,sql" xtractr-param:"timeField" xtractr-time:"ISO80601"` // xtractr supports this (YYYY-MM-DD) ISO8601 format
    FieldTwo      sql.NullString `xtractr:"path,sql" xtractr-param:"fieldTwo"`
    ...
}
```

Once your struct has been defined, it's as simple as instantiating the struct, defining the exact path
pattern to match, and passing the path pattern to match, the incoming http.Request pointer and a pointer
to your defined struct to `xtractr.ExtractParams(pattern string, request *http.Request, dst any)`

```go
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	params := MyStruct {}
	
	err := xtractr.ExtractParams("/{fieldOne}", request, &params)
}
```

Put all together:

```go
package main

import (
	"fmt"
	"github.com/syke99/xtractr"
	"net/http"
	"os"
)

type MyStruct struct {
	FieldOne int    `xtractr:"path" xtractr-param:"fieldOne"`
	FieldTwo string `xtractr:"query" xtractr-param:"fieldTwo"`
}

func main() {
	params := MyStruct{}

	// Typically, one would use Xtractr in an Http.HandlerFunc
	// but the focus of creating this new http Request is on
	// the path being used. Given the path "/1234?fieldTwo=hello",
	// Xtractr will map 1234 to params.FieldOne and "hello" to
	// params.FieldTwo
	request := http.NewRequest(http.MethodGet, "/1234?fieldTwo=hello", nil)

	err := xtractr.ExtractParams("/{fieldOne}", request, &params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(params.FieldOne) // prints 1234
	fmt.Println(params.FieldTwo) // prints hello
}
```

Who?
====

This library was developed by Quinn Millican ([@syke99](https://github.com/syke99))


## License

This repo is under the MIT license, see [LICENSE](../LICENSE) for details.
