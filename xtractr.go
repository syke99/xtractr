package xtractr

import (
	"github.com/syke99/xtractr/internal"
	"github.com/syke99/xtractr/internal/resources"
	"net/http"
	"reflect"
)

// ExtractParams takes a pattern of the path to match parameters against,
// the incoming *http.Request, and a pointer to the struct you would like
// to unmarshal to (dst), extracts all parameters for the request path
// and query, and unmarshals them to dst
func ExtractParams(pattern string, request *http.Request, dst any) error {
	str := reflect.ValueOf(dst)

	dstType := reflect.TypeOf(dst)

	if dstType.Kind() != reflect.Pointer {
		return resources.InvalidDst
	}

	if dstType.Kind() != reflect.Pointer &&
		dstType.Elem().Kind() != reflect.Struct {
		return resources.InvalidDst
	}

	reqPath := request.URL.Path

	pattern, reqPath = internal.SanitizePaths(pattern, reqPath)

	pathParams := internal.GetMatchedPathParams(pattern, reqPath)
	if pathParams == nil {
		return resources.PathParseErr
	}

	return internal.Unmarshal(request, str, pathParams)
}
