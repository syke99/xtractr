package unmarshal

import (
	"github.com/syke99/xtractr/internal/unmarshal/basic"
	"github.com/syke99/xtractr/internal/unmarshal/sql"
	"net/http"
	"reflect"
)

func Unmarshal(request *http.Request, str reflect.Value, pathParams map[string]string) error {
	var err error

	elem := str.Elem()

	for i := 0; i < elem.Type().NumField(); i++ {

		field := elem.Type().Field(i)
		tag := field.Tag

		param := tag.Get("xtractr-param")

		xtractrTag, ok := tag.Lookup("xtractr")
		if !ok {
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
			err := Unmarshal(request, ptr, pathParams)
			if err != nil {
				continue
			}

			elem.Field(i).Set(reflect.ValueOf(ptr.Elem().Interface()))
			continue
		}

		sqlType, err := DetermineSQL(xtractrTag)
		if err != nil {
			return err
		}

		switch sqlType {
		case false:
			err = basic.Unmarshal(i, request, xtractrTag, elem, field, tag, pathParams, param)
		case true:
			err = sql.Unmarshal(i, request, xtractrTag, elem, tag, pathParams, param)
		}

	}

	return err
}
