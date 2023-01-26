package unmarshal

import (
	"errors"
	"fmt"
	"github.com/syke99/xtractr/internal/unmarshal/basic"
	"github.com/syke99/xtractr/internal/unmarshal/sql"
	"net/http"
	"reflect"
	"strings"
)

func Unmarshal(request *http.Request, str reflect.Value, pathParams map[string]string) error {
	var err error

	elem := str.Elem()

	fieldErrs := make([]string, 0)

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
				nestedErrs := fmt.Sprintf("failed to unmarshal fields (%s) in nested struct %s", err.Error(), field.Name)
				fieldErrs = append(fieldErrs, nestedErrs)
				continue
			}

			elem.Field(i).Set(reflect.ValueOf(ptr.Elem().Interface()))
			continue
		}

		sqlType, err := DetermineSQL(xtractrTag)
		if err != nil {
			fieldErrs = append(fieldErrs, field.Name)
			continue
		}

		switch sqlType {
		case false:
			err = basic.Unmarshal(i, request, xtractrTag, elem, field, tag, pathParams, param)
			if err != nil {
				fieldErrs = append(fieldErrs, field.Name)
				continue
			}
		case true:
			err = sql.Unmarshal(i, request, xtractrTag, elem, tag, pathParams, param)
			if err != nil {
				fieldErrs = append(fieldErrs, field.Name)
				continue
			}
		}
	}

	if len(fieldErrs) != 0 {
		err = errors.New(fmt.Sprintf("%s", strings.Join(fieldErrs, ", ")))
	}

	return err
}
