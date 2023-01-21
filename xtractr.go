package xtractr

import (
	"net/http"
	"reflect"
	"strings"
)

func Extract(request *http.Request, dst any) error {
	str := reflect.ValueOf(dst).Elem()

	path := str.FieldByName("XtractrPath")

	// TODO: after getting the path to match, determine if there are any params to grab

	for i := 0; i < str.NumField(); i++ {
		if strings.Contains(string(str.Type().Field(i).Tag), "json:\"-\"") {
			continue
		}

	}

	println(path)

	return nil
}
