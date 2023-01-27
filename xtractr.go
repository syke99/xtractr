package xtractr

import (
	"github.com/syke99/xtractr/internal"
	"github.com/syke99/xtractr/internal/pkg/resources"
	"github.com/syke99/xtractr/internal/unmarshal"
	"net/http"
	"reflect"
)

// ExtractParams takes a pattern of the path to match parameters against,
// the incoming *http.Request, and a pointer to the struct you would like
// to unmarshal to (dst), extracts all parameters for the request path
// and query, and unmarshals them to dst
func ExtractParams(pattern string, request *http.Request, dst any) error {
	if pattern == "" {
		return resources.MissingPattern
	}

	if request == nil {
		return resources.MissingIncomingRequest
	}

	dstType := reflect.TypeOf(dst)

	if dstType.Kind() != reflect.Pointer ||
		dstType.Elem().Kind() != reflect.Struct {
		return resources.InvalidDst
	}

	str := reflect.ValueOf(dst)

	reqPath := request.URL.EscapedPath()

	pattern, reqPath = internal.SanitizePaths(pattern, reqPath)

	pathParams := internal.GetMatchedPathParams(pattern, reqPath)
	if pathParams == nil {
		return resources.PathParseErr
	}

	return unmarshal.Unmarshal(request, str, pathParams)
}
