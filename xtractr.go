package xtractr

import (
	"errors"
	"github.com/syke99/xtractr/internal"
	"github.com/syke99/xtractr/internal/unmarshal"
	"net/http"
	"reflect"
)

func ExtractParams(request *http.Request, dst any) error {
	str := reflect.ValueOf(dst)

	if str.Kind() != reflect.Pointer &&
		str.Elem().Kind() != reflect.Struct {
		return errors.New("dst provided is not a valid pointer to a struct")
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
		return errors.New("error parsing path for parameters")
	}

	return unmarshal.Unmarshal(request, str, pathParams)
}
