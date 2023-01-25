package xtractr

import (
	"errors"
	"github.com/syke99/xtractr/internal"
	"github.com/syke99/xtractr/internal/unmarshal"
	"net/http"
	"reflect"
)

// ExtractParams takes a pattern of the path to match parameters against,
// the incoming *http.Request, and a pointer to the struct you would like
// to unmarshal to (dst), extracts all parameters for the request path
// and query, and unmarshals them to dst
func ExtractParams(pattern string, request *http.Request, dst any) error {
	str := reflect.ValueOf(dst)

	if str.Kind() != reflect.Pointer &&
		str.Elem().Kind() != reflect.Struct {
		return errors.New("dst provided is not a valid pointer to a struct")
	}

	if pattern[:1] == "/" {
		pattern = pattern[1:]
	}

	if pattern[len(pattern)-1:] == "/" {
		pattern = pattern[:len(pattern)-1]
	}

	reqPath := request.URL.Path

	if reqPath[:1] == "/" {
		reqPath = reqPath[1:]
	}

	if reqPath[len(reqPath)-1:] == "/" {
		reqPath = reqPath[:len(reqPath)-1]
	}

	pathParams := internal.GetMatchedPathParams(pattern, reqPath)
	if pathParams == nil {
		return errors.New("error parsing path for parameters")
	}

	return unmarshal.Unmarshal(request, str, pathParams)
}
