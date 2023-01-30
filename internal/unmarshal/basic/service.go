package basic

import (
	"github.com/syke99/xtractr/internal/unmarshal/common"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

func Unmarshal(i int, queryValues url.Values, xtractrTag string, elem reflect.Value, field reflect.StructField, tag reflect.StructTag, pathParams map[string]string, param string) error {

	if xtractrTag == "query" &&
		elem.Field(i).CanSet() {

		vals, ok := queryValues[param]
		if !ok {
			return nil
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			switch elem.Field(i).Interface().(type) {
			case time.Time:
				t, err := common.FormatTime(vals[0], tag.Get("xtractr-time"))
				if err != nil {
					return err
				}
				elem.Field(i).Set(reflect.ValueOf(t))
			default:
				return nil
			}
		case reflect.Bool:
			b := false
			if queryValues.Has(param) &&
				queryValues.Get(param) != "" {
				v, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}

				b = v
			}

			if queryValues.Has(param) &&
				queryValues.Get(param) == "" {
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
				return err
			}
			elem.Field(i).SetInt(v)
		case reflect.Float32, reflect.Float64:
			bz := 32
			if field.Type.Kind() == reflect.Float64 {
				bz = 64
			}
			v, err := strconv.ParseFloat(vals[0], bz)
			if err != nil {
				return err
			}
			elem.Field(i).SetFloat(v)
		case reflect.Complex64, reflect.Complex128:
			bz := 64
			if field.Type.Kind() == reflect.Complex128 {
				bz = 128
			}
			v, err := strconv.ParseComplex(vals[0], bz)
			if err != nil {
				return err
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
				return err
			}
			elem.Field(i).SetUint(v)
		case reflect.Array, reflect.Slice:
			elem.Field(i).Set(reflect.ValueOf(vals))
		}
	}

	if xtractrTag == "path" &&
		elem.Field(i).CanSet() {

		j, ok := pathParams[param]
		if !ok {
			if !ok {
				return nil
			}
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			switch elem.Field(i).Interface().(type) {
			case time.Time:
				t, err := common.FormatTime(j, tag.Get("xtractr-time"))
				if err != nil {
					return err
				}
				elem.Field(i).Set(reflect.ValueOf(t))
			default:
				return nil
			}
		case reflect.Bool:
			v, err := strconv.ParseBool(j)
			if err != nil {
				return err
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
				return err
			}
			elem.Field(i).SetInt(v)
		case reflect.Float32, reflect.Float64:
			bz := 32
			if field.Type.Kind() == reflect.Float64 {
				bz = 64
			}
			v, err := strconv.ParseFloat(j, bz)
			if err != nil {
				return err
			}
			elem.Field(i).SetFloat(v)
		case reflect.Complex64, reflect.Complex128:
			bz := 64
			if field.Type.Kind() == reflect.Complex128 {
				bz = 128
			}
			v, err := strconv.ParseComplex(j, bz)
			if err != nil {
				return err
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
				return err
			}
			elem.Field(i).SetUint(v)
		}
	}

	return nil
}
