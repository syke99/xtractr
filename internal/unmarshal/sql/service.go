package sql

import (
	"database/sql"
	"github.com/syke99/xtractr/internal/models"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Unmarshal(i int, request *http.Request, xtractrTag string, elem reflect.Value, field reflect.StructField, tag reflect.StructTag, pathParams map[string]string, jsonTag string) {

	xtractrTag = strings.Split(xtractrTag, ",")[0]

	if xtractrTag == "query" &&
		elem.Field(i).CanSet() {

		vals, ok := request.URL.Query()[jsonTag]
		if !ok {
			return
		}

		switch elem.Field(i).Interface().(type) {
		case sql.NullBool:
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
			nb := sql.NullBool{
				Bool:  b,
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(nb))
		case sql.NullString:
			ns := sql.NullString{
				String: vals[0],
				Valid:  true,
			}
			elem.Field(i).Set(reflect.ValueOf(ns))
		case sql.NullInt16:
			v, err := strconv.ParseInt(vals[0], 10, 16)
			if err != nil {
				return
			}
			ni16 := sql.NullInt16{
				Int16: int16(v),
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(ni16))
		case sql.NullInt32:
			v, err := strconv.ParseInt(vals[0], 10, 32)
			if err != nil {
				return
			}
			ni32 := sql.NullInt32{
				Int32: int32(v),
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(ni32))
		case sql.NullInt64:
			v, err := strconv.ParseInt(vals[0], 10, 64)
			if err != nil {
				return
			}
			ni64 := sql.NullInt64{
				Int64: v,
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(ni64))
		case sql.NullFloat64:
			v, err := strconv.ParseFloat(vals[0], 64)
			if err != nil {
				return
			}
			nf64 := sql.NullFloat64{
				Float64: v,
				Valid:   true,
			}
			elem.Field(i).Set(reflect.ValueOf(nf64))
		case sql.NullTime:
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
						return
					}

					m := 0
					m, er = strconv.Atoi(tParts[1])
					if er != nil {
						return
					}

					month = time.Month(m)

					day, er = strconv.Atoi(tParts[2])
					if er != nil {
						return
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
					return
				}
			}
			s := sql.NullTime{
				Time:  t,
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(s))
		default:
			return
		}
	}

	if xtractrTag == "path" &&
		elem.Field(i).CanSet() {

		j, ok := pathParams[jsonTag]
		if !ok {
			return
		}

		switch elem.Field(i).Interface().(type) {
		case sql.NullBool:
			b := false
			if request.URL.Query().Has(jsonTag) &&
				request.URL.Query().Get(jsonTag) != "" {
				v, err := strconv.ParseBool(j)
				if err != nil {
					return
				}

				b = v
			}

			if request.URL.Query().Has(jsonTag) &&
				request.URL.Query().Get(jsonTag) == "" {
				b = true
			}
			nb := sql.NullBool{
				Bool:  b,
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(nb))
		case sql.NullString:
			ns := sql.NullString{
				String: j,
				Valid:  false,
			}
			elem.Field(i).Set(reflect.ValueOf(ns))
		case sql.NullInt16:
			v, err := strconv.ParseInt(j, 10, 16)
			if err != nil {
				return
			}
			ni16 := sql.NullInt16{
				Int16: int16(v),
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(ni16))
		case sql.NullInt32:
			v, err := strconv.ParseInt(j, 10, 32)
			if err != nil {
				return
			}
			ni32 := sql.NullInt32{
				Int32: int32(v),
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(ni32))
		case sql.NullInt64:
			v, err := strconv.ParseInt(j, 10, 64)
			if err != nil {
				return
			}
			ni64 := sql.NullInt64{
				Int64: v,
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(ni64))
		case sql.NullFloat64:
			v, err := strconv.ParseFloat(j, 64)
			if err != nil {
				return
			}
			nf64 := sql.NullFloat64{
				Float64: v,
				Valid:   true,
			}
			elem.Field(i).Set(reflect.ValueOf(nf64))
		case sql.NullTime:
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
						return
					}

					m := 0
					m, er = strconv.Atoi(tParts[1])
					if er != nil {
						return
					}

					month = time.Month(m)

					day, er = strconv.Atoi(tParts[2])
					if er != nil {
						return
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
					return
				}
			}
			s := sql.NullTime{
				Time:  t,
				Valid: true,
			}
			elem.Field(i).Set(reflect.ValueOf(s))
		default:
			return
		}
	}
}
