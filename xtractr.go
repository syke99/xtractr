package xtractr

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
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

	pathParams := getMatchedPathParams(matchPath, reqPath)
	if pathParams == nil {
		return
	}

	unmarshal(request, str, pathParams, true)
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

func unmarshal(request *http.Request, str reflect.Value, pathParams map[string]string, parent bool) {
	//var elem reflect.Value
	//
	//println(str.Kind())
	//
	//if str.Kind() == reflect.Pointer &&
	//	reflect.TypeOf(str.Elem().Interface()).Kind() == reflect.Struct &&
	//	!parent {
	//	s := reflect.ValueOf(str.Elem())
	//
	//	elem = s.Elem()
	//} else {
	//	elem = str.Elem()
	//}

	elem := str.Elem()

	for i := 0; i < elem.Type().NumField(); i++ {
		field := elem.Type().Field(i)
		tag := field.Tag

		jsonTag := tag.Get("json")

		xtractrTag := tag.Get("xtractr")

		if xtractrTag == "-" {
			continue
		}

		if xtractrTag == "struct" {
			s := elem.Field(i).Interface()

			ptr := reflect.PointerTo(reflect.TypeOf(s))

			unmarshal(request, reflect.ValueOf(ptr), pathParams, false)
		}

		if xtractrTag == "query" &&
			elem.Field(i).CanSet() {

			vals := request.URL.Query()[jsonTag]

			switch field.Type.Kind() {
			case reflect.Struct:
				switch elem.Field(i).Interface().(type) {
				case time.Time:
					time.Parse(tag.Get("xtractr-time"), vals[0])
				default:
					continue
				}
			case reflect.Bool:
				b := false
				if request.URL.Query().Has(jsonTag) &&
					request.URL.Query().Get(jsonTag) != "" {
					v, err := strconv.ParseBool(vals[0])
					if err != nil {
						return
					}

					b = v
				}

				if request.URL.Query().Has(jsonTag) &&
					request.URL.Query().Get(jsonTag) == "" {
					b = true
				}

				elem.Field(i).SetBool(b)
			case reflect.String:
				elem.Field(i).SetString(vals[0])
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				bz := 0
				switch field.Type.Kind() {
				case reflect.Int8:
					bz = 8
				case reflect.Int16:
					bz = 16
				case reflect.Int32:
					bz = 32
				case reflect.Int64:
					bz = 64
				}
				v, err := strconv.ParseInt(vals[0], 10, bz)
				if err != nil {
					return
				}
				elem.Field(i).SetInt(v)
			case reflect.Float32, reflect.Float64:
				bz := 32
				if field.Type.Kind() == reflect.Float64 {
					bz = 64
				}
				v, err := strconv.ParseFloat(vals[0], bz)
				if err != nil {
					return
				}
				elem.Field(i).SetFloat(v)
			case reflect.Complex64, reflect.Complex128:
				bz := 64
				if field.Type.Kind() == reflect.Complex128 {
					bz = 128
				}
				v, err := strconv.ParseComplex(vals[0], bz)
				if err != nil {
					return
				}
				elem.Field(i).SetComplex(v)
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
					return
				}
				elem.Field(i).SetUint(v)
			}
		}

		if xtractrTag == "path" &&
			elem.Field(i).CanSet() {

			j := pathParams[jsonTag]

			switch field.Type.Kind() {
			case reflect.Struct:
				switch elem.Field(i).Interface().(type) {
				case time.Time:
					time.Parse(tag.Get("xtractr-time"), j)
				default:
					continue
				}
			case reflect.Bool:
				v, err := strconv.ParseBool(j)
				if err != nil {
					return
				}
				elem.Field(i).SetBool(v)
			case reflect.String:
				elem.Field(i).SetString(j)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				bz := 0
				switch field.Type.Kind() {
				case reflect.Int8:
					bz = 8
				case reflect.Int16:
					bz = 16
				case reflect.Int32:
					bz = 32
				case reflect.Int64:
					bz = 64
				}
				v, err := strconv.ParseInt(j, 10, bz)
				if err != nil {
					return
				}
				elem.Field(i).SetInt(v)
			case reflect.Float32, reflect.Float64:
				bz := 32
				if field.Type.Kind() == reflect.Float64 {
					bz = 64
				}
				v, err := strconv.ParseFloat(j, bz)
				if err != nil {
					return
				}
				elem.Field(i).SetFloat(v)
			case reflect.Complex64, reflect.Complex128:
				bz := 64
				if field.Type.Kind() == reflect.Complex128 {
					bz = 128
				}
				v, err := strconv.ParseComplex(j, bz)
				if err != nil {
					return
				}
				elem.Field(i).SetComplex(v)
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
				v, err := strconv.ParseUint(j, 10, bz)
				if err != nil {
					return
				}
				elem.Field(i).SetUint(v)
			case reflect.Array, reflect.Slice:
				elem.Set(reflect.ValueOf(j))
			}
		}
	}
}
