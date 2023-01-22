package xtractr

import (
	"github.com/syke99/xtractr/internal"
	"github.com/syke99/xtractr/internal/unmarshal"
	"net/http"
	"reflect"
)

func ExtractParams(request *http.Request, dst any) {
	str := reflect.ValueOf(dst)

	if str.Kind() != reflect.Pointer &&
		str.Elem().Kind() != reflect.Struct {
		return
	}

	matchPath := str.Elem().FieldByName("Xtractr").String()

	if matchPath[:1] == "/" {
		matchPath = matchPath[1:]
	}

	if matchPath[len(matchPath)-1:] == "/" {
		matchPath = matchPath[:len(matchPath)-1]
	}

	reqPath := request.URL.Path

	if reqPath[:1] == "/" {
		reqPath = reqPath[1:]
	}

	if reqPath[len(reqPath)-1:] == "/" {
		reqPath = reqPath[:len(reqPath)-1]
	}

	pathParams := internal.GetMatchedPathParams(matchPath, reqPath)
	if pathParams == nil {
		return
	}

	unmarshal.Unmarshal(request, str, pathParams)
}
