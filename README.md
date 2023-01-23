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
field via a `json` tag, and where to find that parameter (in the path or query) via a `xtractr` tag

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

After importing, define the struct you would like to unmarshal the parameters into. Here is where you
can define the path pattern with path parameters to match against by defining a field in your struct
named XtractrPath, as well as whether each subsequent struct field's value should be found in the request
path or query using the `xtractr` tag. Example:

```go
type MyStruct struct {
	Xtractr     string `xtractr:"-"` // Xtractr only needs a `xtractr` tag with the value "-"
	FieldOne    int    `json:"fieldOne" xtractr:"path"`
	FieldTwo    string `json:"fieldTwo" xtractr:"query"`
}
```

You can even nest structs if you'd like, just simply give the `xtractr` tag the `struct` value and Xtractr 
will recurse into that struct and unmarshal its defined parameters.

```go
type MyFirstStruct struct {
    FieldOne    int    `json:"fieldOne" xtractr:"path"`
    FieldTwo    string `json:"fieldTwo" xtractr:"query"`
}

type MySecondStruc struct {
    Xtractr       string        `xtractr:"-"` // Xtractr only needs a `xtractr` tag with the value "-"
    FieldThree    int           `json:"fieldThree" xtractr:"path"`
    FieldFour     MyFirstStruct `xtractr:"struct"`
}
```

Xtractr also suports time.Time (and sql.NullTime) fields. You just need to provide an additional `xtractr-time`
tag to specify the format to be used (Xtractr defaults to Go's default time format):

```go
type MyStruct struct {
    Xtractr      string    `xtractr:"-"` // Xtractr only needs a `xtractr` tag with the value "-"
    TimeField    time.Time `json:"timeField" xtractr:"path" xtractr-time:"ISO80601"` // xtractr supports this (YYYY-MM-DD) ISO8601 format}
```

You can also use any SQL NullType (besides sql.NullByte) fields, just specify by appending `,sql` to the end
of the value you supply in the `xtractr` tag

```go
type MyStruct struct {
    Xtractr       string         `xtractr:"-"` // Xtractr only needs a `xtractr` tag with the value "-"
    FieldOne      sql.NullTime   `json:"timeField" xtractr:"path,sql" xtractr-time:"ISO80601"` // xtractr supports this (YYYY-MM-DD) ISO8601 format
    FieldTwo      sql.NullString `json:fieldTwo" xtractr:"path,sql"`
    ...
}
```

Once your struct has been defined, it's as simple as instantiating the struct, defining the exact path
pattern to match, and passing the incoming http.Request pointer and a pointer to your defined struct 
to `xtractr.ExtractParams(request *http.Request, dst any)`

```go
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	params := MyStruct {Xtractr: "/{fieldOne}"}
	
	xtractr.ExtractParams(r, &params)
}
```

Put all together:

```go
package main

import (
	"fmt"
	"github.com/syke99/xtractr"
	"net/http"
)

type MyStruct struct {
	Xtractr     string `xtractr:"-"` // Xtractr only needs a `xtractr` tag with the value "-"
	FieldOne    int    `json:"fieldOne" xtractr:"path"`
	FieldTwo    string `json:"fieldTwo" xtractr:"query"`
}

func main() {
	params := MyStruct{Xtractr: "/{fieldOne}"}

	// Typically, one would use Xtractr in an Http.HandlerFunc
	// but the focus of creating this new http Request is on
	// the path being used. Given the path "/1234?fieldTwo=hello",
	// Xtractr will map 1234 to params.FieldOne and "hello" to
	// params.FieldTwo
	request := http.NewRequest(http.MethodGet, "/1234?fieldTwo=hello", nil)

	xtractr.ExtractParams(request, &params)

	fmt.Println(params.FieldOne) // prints 1234
	fmt.Println(params.FieldTwo) // prints hello
}
```

Who?
====

This library was developed by Quinn Millican ([@syke99](https://github.com/syke99))


## License

This repo is under the BSD 3 license, see [LICENSE](../LICENSE) for details.
