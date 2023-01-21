package xtractr

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func Extract(request *http.Request, dst any) error {
	str := reflect.ValueOf(dst).Elem()

	matchPath := str.FieldByName("XtractrPath").String()

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

	for i := 0; i < str.NumField(); i++ {
		field := str.Type().Field(i)
		tag := field.Tag

		if tag.Get("json") == "-" {
			continue
		}

		if tag.Get("xtractr") == "path" {
			switch str.Field(i).Type().Kind() {
			case reflect.Bool:
				v, err := strconv.ParseBool(pathParams[tag.Get("json")])
				if err != nil {
					return err
				}
				if str.Field(i).CanSet() {
					str.Field(i).SetBool(v)
				}
			case reflect.String:
				if str.Field(i).CanSet() {
					str.Field(i).SetString(pathParams[tag.Get("json")])
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v, err := strconv.Atoi(pathParams[tag.Get("json")])
				if err != nil {
					return err
				}
				if str.Field(i).CanSet() {
					str.Field(i).SetInt(int64(v))
				}
			case reflect.Float32, reflect.Float64:
				v, err := strconv.Atoi(pathParams[tag.Get("json")])
				if err != nil {
					return err
				}
				if str.Field(i).CanSet() {
					str.Field(i).SetFloat(float64(v))
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				v, err := strconv.Atoi(pathParams[tag.Get("json")])
				if err != nil {
					return err
				}
				if str.Field(i).CanSet() {
					str.Field(i).SetUint(uint64(v))
				}
			}
		}

		if tag.Get("xtractr") == "query" {
			if !request.URL.Query().Has(tag.Get("json")) {
				return errors.New(fmt.Sprintf("no query parameter found for field: %s", str.Type().Field(i).Name))
			}

			val := request.URL.Query().Get(tag.Get("json"))
			switch str.Field(i).Type().Kind() {
			case reflect.Bool:
				if str.Field(i).CanSet() {
					v, err := strconv.ParseBool(val)
					if err != nil {
						return err
					}
					str.Field(i).SetBool(v)
				}
			case reflect.String:
				if str.Field(i).CanSet() {
					str.Field(i).SetString(val)
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				if str.Field(i).CanSet() {
					str.Field(i).SetInt(int64(v))
				}
			case reflect.Float32, reflect.Float64:
				v, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				if str.Field(i).CanSet() {
					str.Field(i).SetFloat(float64(v))
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				v, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				if str.Field(i).CanSet() {
					str.Field(i).SetUint(uint64(v))
				}
			}
		}
	}

	return nil
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

			m[matchPathParts[i]] = reqPart
		}
	}

	return m
}
