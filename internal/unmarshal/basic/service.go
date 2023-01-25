package basic

import (
	"errors"
	"github.com/syke99/xtractr/internal/models"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Unmarshal(i int, request *http.Request, xtractrTag string, elem reflect.Value, field reflect.StructField, tag reflect.StructTag, pathParams map[string]string, jsonTag string) error {

	if xtractrTag == "query" &&
		elem.Field(i).CanSet() {

		vals, ok := request.URL.Query()[jsonTag]
		if !ok {
			return errors.New("parameter not found in query")
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			switch elem.Field(i).Interface().(type) {
			case time.Time:
				var t time.Time
				var err error

				format := tag.Get("xtractr-time")

				layout := ""

				if format == "" || format == "ISO8601" {
					if format == "ISO8601" {
						var year int
						var month time.Month
						var day int
						var er error

						tParts := strings.Split(vals[0], "-")

						year, er = strconv.Atoi(tParts[0])
						if er != nil {
							return er
						}

						m := 0
						m, er = strconv.Atoi(tParts[1])
						if er != nil {
							return er
						}

						month = time.Month(m)

						day, er = strconv.Atoi(tParts[2])
						if er != nil {
							return er
						}

						t = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
					} else {
						layout = time.Layout
					}
				}

				if format != "" && format != "ISO8601" {
					if f, ok := models.TimeLayouts()[format]; ok {
						layout = f
					}

					t, err = time.Parse(layout, vals[0])
					if err != nil {
						return err
					}
				}
				elem.Field(i).Set(reflect.ValueOf(t))
			default:
				return nil
			}
		case reflect.Bool:
			b := false
			if request.URL.Query().Has(jsonTag) &&
				request.URL.Query().Get(jsonTag) != "" {
				v, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
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
		}
	}

	if xtractrTag == "path" &&
		elem.Field(i).CanSet() {

		j, ok := pathParams[jsonTag]
		if !ok {
			if !ok {
				return errors.New("parameter not found in path")
			}
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			switch elem.Field(i).Interface().(type) {
			case time.Time:
				var t time.Time
				var err error

				format := tag.Get("xtractr-time")

				layout := ""

				if format == "" || format == "ISO8601" {
					if format == "ISO8601" {
						var year int
						var month time.Month
						var day int
						var er error

						tParts := strings.Split(j, "-")

						year, er = strconv.Atoi(tParts[0])
						if er != nil {
							return err
						}

						m := 0
						m, er = strconv.Atoi(tParts[1])
						if er != nil {
							return err
						}

						month = time.Month(m)

						day, er = strconv.Atoi(tParts[2])
						if er != nil {
							return err
						}

						t = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
					} else {
						layout = time.Layout
					}
				}

				if format != "" && format != "ISO8601" {
					if f, ok := models.TimeLayouts()[format]; ok {
						layout = f
					}

					t, err = time.Parse(layout, j)
					if err != nil {
						return err
					}
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
		case reflect.Array, reflect.Slice:
			elem.Set(reflect.ValueOf(j))
		}
	}

	return nil
}
