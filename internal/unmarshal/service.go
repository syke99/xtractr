package unmarshal

import (
	"errors"
	"fmt"
	"github.com/syke99/xtractr/internal/unmarshal/basic"
	"github.com/syke99/xtractr/internal/unmarshal/sql"
	"net/url"
	"reflect"
	"strings"
)

func Unmarshal(queryValues url.Values, str reflect.Value, pathParams map[string]string) error {
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
			ptr, err := recurseIntoNestedStruct(field, queryValues, pathParams)
			if err != nil {
				nestedErrs := fmt.Sprintf("failed to unmarshal fields (%s) in nested struct %s", err.Error(), field.Name)
				fieldErrs = append(fieldErrs, nestedErrs)
				continue
			}

			if elem.Field(i).CanSet() {
				elem.Field(i).Set(reflect.ValueOf(ptr.Elem().Interface()))
			}
			continue
		}

		sqlType, err := DetermineSQL(xtractrTag)
		if err != nil {
			fieldErrs = append(fieldErrs, field.Name)
			continue
		}

		err = unmarshalOnType(sqlType, i, queryValues, xtractrTag, elem, field, tag, pathParams, param)
		if err != nil {
			fieldErrs = append(fieldErrs, fmt.Sprintf("field: %s, error: %s", field.Name, err.Error()))
			continue
		}
	}

	if len(fieldErrs) != 0 {
		err = errors.New(fmt.Sprintf("%s", strings.Join(fieldErrs, ",\n")))
	}

	return err
}

func recurseIntoNestedStruct(field reflect.StructField, queryValues url.Values, pathParams map[string]string) (reflect.Value, error) {
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
	err := Unmarshal(queryValues, ptr, pathParams)

	return ptr, err
}

func unmarshalOnType(sqlType bool, i int, queryValues url.Values, xtractrTag string, elem reflect.Value, field reflect.StructField, tag reflect.StructTag, pathParams map[string]string, param string) error {
	var err error
	switch sqlType {
	case false:
		err = basic.Unmarshal(i, queryValues, xtractrTag, elem, field, tag, pathParams, param)
	case true:
		err = sql.Unmarshal(i, queryValues, xtractrTag, elem, tag, pathParams, param)
	}

	return err
}
