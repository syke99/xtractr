package unmarshal

import (
	"github.com/syke99/xtractr/internal/unmarshal/basic"
	"github.com/syke99/xtractr/internal/unmarshal/sql"
	"net/http"
	"reflect"
)

func Unmarshal(request *http.Request, str reflect.Value, pathParams map[string]string) {
	elem := str.Elem()

	for i := 0; i < elem.Type().NumField(); i++ {

		field := elem.Type().Field(i)
		tag := field.Tag

		jsonTag := tag.Get("json")

		xtractrTag, ok := tag.Lookup("xtractr")
		if !ok {
			return
		}

		if xtractrTag == "-" {
			continue
		}

		if xtractrTag == "struct" {
			num := field.Type.NumField()
			sF := make([]reflect.StructField, num)
			for i := 0; i < num; i++ {
				fl := field.Type.Field(i)
				f := reflect.StructField{
					Name: fl.Name,
					Type: fl.Type,
					Tag:  fl.Tag,
				}
				sF[i] = f
			}

			ptr := reflect.New(reflect.StructOf(sF))
			Unmarshal(request, ptr, pathParams)

			elem.Field(i).Set(reflect.ValueOf(ptr.Elem().Interface()))
			continue
		}

		sqlType, err := DetermineSQL(xtractrTag)
		if err != nil {
			return
		}

		switch sqlType {
		case false:
			basic.Unmarshal(i, request, xtractrTag, elem, field, tag, pathParams, jsonTag)
		case true:
			sql.Unmarshal(i, request, xtractrTag, elem, field, tag, pathParams, jsonTag)
		}

	}
}
