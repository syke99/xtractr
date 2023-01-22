package sql

import (
	"net/http"
	"reflect"
)

func Unmarshal(i int, request *http.Request, xtractrTag string, elem reflect.Value, field reflect.StructField, tag reflect.StructTag, pathParams map[string]string, jsonTag string) {
	//
	//if xtractrTag == "query" &&
	//	elem.Field(i).CanSet() {
	//
	//	vals := request.URL.Query()[jsonTag]
	//
	//	switch field.Type.Kind() {
	//	// TODO: fix time fields and nested structs
	//	//case reflect.Struct:
	//	//	switch elem.Field(i).Interface().(type) {
	//	//	case time.Time:
	//	//		t, err := time.Parse(tag.Get("xtractr-time"), vals[0])
	//	//		if err != nil {
	//	//			println(err.Error())
	//	//			return
	//	//		}
	//	//		elem.Field(i).Set(reflect.ValueOf(t))
	//	//	case sql.NullTime:
	//	//		t, err := time.Parse(tag.Get("xtractr-time"), vals[0])
	//	//		if err != nil {
	//	//			return
	//	//		}
	//	//
	//	//		s := sql.NullTime{
	//	//			Time:  t,
	//	//			Valid: true,
	//	//		}
	//	//		elem.Field(i).Set(reflect.ValueOf(s))
	//	//	default:
	//	//		continue
	//	//	}
	//	case reflect.Bool:
	//		b := false
	//		if request.URL.Query().Has(jsonTag) &&
	//			request.URL.Query().Get(jsonTag) != "" {
	//			v, err := strconv.ParseBool(vals[0])
	//			if err != nil {
	//				return
	//			}
	//
	//			b = v
	//		}
	//
	//		if request.URL.Query().Has(jsonTag) &&
	//			request.URL.Query().Get(jsonTag) == "" {
	//			b = true
	//		}
	//
	//		elem.Field(i).SetBool(b)
	//	case reflect.String:
	//		elem.Field(i).SetString(vals[0])
	//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//		bz := 0
	//		switch field.Type.Kind() {
	//		case reflect.Int8:
	//			bz = 8
	//		case reflect.Int16:
	//			bz = 16
	//		case reflect.Int32:
	//			bz = 32
	//		case reflect.Int64:
	//			bz = 64
	//		}
	//		v, err := strconv.ParseInt(vals[0], 10, bz)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetInt(v)
	//	case reflect.Float32, reflect.Float64:
	//		bz := 32
	//		if field.Type.Kind() == reflect.Float64 {
	//			bz = 64
	//		}
	//		v, err := strconv.ParseFloat(vals[0], bz)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetFloat(v)
	//	case reflect.Complex64, reflect.Complex128:
	//		bz := 64
	//		if field.Type.Kind() == reflect.Complex128 {
	//			bz = 128
	//		}
	//		v, err := strconv.ParseComplex(vals[0], bz)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetComplex(v)
	//	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	//		bz := 0
	//		switch field.Type.Kind() {
	//		case reflect.Uint8:
	//			bz = 8
	//		case reflect.Uint16:
	//			bz = 16
	//		case reflect.Uint32:
	//			bz = 32
	//		case reflect.Uint64:
	//			bz = 64
	//		}
	//		v, err := strconv.ParseUint(vals[0], 10, bz)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetUint(v)
	//	}
	//}
	//
	//if xtractrTag == "path" &&
	//	elem.Field(i).CanSet() {
	//
	//	j := pathParams[jsonTag]
	//
	//	switch field.Type.Kind() {
	//	case reflect.Struct:
	//		switch elem.Field(i).Interface().(type) {
	//		case time.Time:
	//			time.Parse(tag.Get("xtractr-time"), j)
	//		default:
	//			return
	//		}
	//	case reflect.Bool:
	//		v, err := strconv.ParseBool(j)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetBool(v)
	//	case reflect.String:
	//		elem.Field(i).SetString(j)
	//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//		bz := 0
	//		switch field.Type.Kind() {
	//		case reflect.Int8:
	//			bz = 8
	//		case reflect.Int16:
	//			bz = 16
	//		case reflect.Int32:
	//			bz = 32
	//		case reflect.Int64:
	//			bz = 64
	//		}
	//		v, err := strconv.ParseInt(j, 10, bz)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetInt(v)
	//	case reflect.Float32, reflect.Float64:
	//		bz := 32
	//		if field.Type.Kind() == reflect.Float64 {
	//			bz = 64
	//		}
	//		v, err := strconv.ParseFloat(j, bz)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetFloat(v)
	//	case reflect.Complex64, reflect.Complex128:
	//		bz := 64
	//		if field.Type.Kind() == reflect.Complex128 {
	//			bz = 128
	//		}
	//		v, err := strconv.ParseComplex(j, bz)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetComplex(v)
	//	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	//		bz := 0
	//		switch field.Type.Kind() {
	//		case reflect.Uint8:
	//			bz = 8
	//		case reflect.Uint16:
	//			bz = 16
	//		case reflect.Uint32:
	//			bz = 32
	//		case reflect.Uint64:
	//			bz = 64
	//		}
	//		v, err := strconv.ParseUint(j, 10, bz)
	//		if err != nil {
	//			return
	//		}
	//		elem.Field(i).SetUint(v)
	//	case reflect.Array, reflect.Slice:
	//		elem.Set(reflect.ValueOf(j))
	//	}
	//}
}
