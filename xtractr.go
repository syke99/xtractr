package main

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func ExtractParams(request *http.Request, dst any) error {
	str := reflect.ValueOf(dst)

	matchPath := str.Elem().FieldByName("XtractrPath").String()

	if matchPath[:1] == "/" {
		matchPath = matchPath[1:]
	}

	reqPath := request.URL.Path

	if reqPath[:1] == "/" {
		reqPath = reqPath[1:]
	}

	pathParams := getMatchedPathParams(matchPath, reqPath)
	if pathParams == nil {
		return errors.New("cannot parse request path with expected XtractrPath in dst")
	}

	return unmarshal(request, str, pathParams)
}

func getMatchedPathParams(toMatch string, requested string) map[string]string {

	matchPathParts := strings.Split(toMatch, "/")

	reqPathParts := strings.Split(requested, "/")

	if len(matchPathParts) != len(reqPathParts) {
		return nil
	}

	m := make(map[string]string)

	for i, reqPart := range reqPathParts {
		if matchPathParts[i][:1] == "{" &&
			matchPathParts[i][len(matchPathParts[i])-1:] == "}" {

			strippedMatch := matchPathParts[i][1 : len(matchPathParts[i])-1]

			m[strippedMatch] = reqPart
		}
	}

	return m
}

func unmarshal(request *http.Request, str reflect.Value, pathParams map[string]string) error {
	for i := 0; i < str.Elem().Type().NumField(); i++ {
		field := str.Elem().Type().Field(i)
		tag := field.Tag

		jsonTag := tag.Get("json")

		if jsonTag == "-" {
			continue
		}

		xtractrTag := tag.Get("xtractr")

		if xtractrTag == "query" &&
			str.Elem().Field(i).CanSet() {

			vals := request.URL.Query()[jsonTag]

			switch field.Type.Kind() {
			case reflect.Bool:
				v, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetBool(v)
			case reflect.String:
				str.Elem().Field(i).SetString(vals[0])
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v, err := strconv.Atoi(vals[0])
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetInt(int64(v))
			case reflect.Float32, reflect.Float64:
				bz := 32
				if field.Type.Kind() == reflect.Float64 {
					bz = 64
				}
				v, err := strconv.ParseFloat(vals[0], bz)
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetFloat(v)
			case reflect.Complex64, reflect.Complex128:
				bz := 64
				if field.Type.Kind() == reflect.Complex128 {
					bz = 128
				}
				v, err := strconv.ParseComplex(vals[0], bz)
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetComplex(v)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				bz := 0
				switch field.Type.Kind() {
				case reflect.Uint8:
					bz = 8
				case reflect.Uint16:
					bz = 16
				case reflect.Uint32:
					bz = 32
				case reflect.Uint64:
					bz = 64
				}
				v, err := strconv.ParseUint(vals[0], 10, bz)
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetUint(v)
			case reflect.Array, reflect.Slice:
				str.Elem().Field(i).Set(reflect.ValueOf(vals))
			}
		}

		if xtractrTag == "path" &&
			str.Elem().Field(i).CanSet() {

			switch field.Type.Kind() {
			case reflect.Bool:
				v, err := strconv.ParseBool(pathParams[jsonTag])
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetBool(v)
			case reflect.String:
				str.Elem().Field(i).SetString(pathParams[jsonTag])
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v, err := strconv.Atoi(pathParams[jsonTag])
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetInt(int64(v))
			case reflect.Float32, reflect.Float64:
				v, err := strconv.Atoi(pathParams[jsonTag])
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetFloat(float64(v))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				v, err := strconv.Atoi(pathParams[jsonTag])
				if err != nil {
					return err
				}
				str.Elem().Field(i).SetUint(uint64(v))
			case reflect.Array, reflect.Slice:
				str.Elem().Set(reflect.ValueOf(pathParams[jsonTag]))
			}
		}
	}

	return nil
}
