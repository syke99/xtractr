package unmarshal

import (
	"errors"
	"strings"
)

func DetermineSQL(xtractrTag string) (bool, error) {
	xTags := strings.Split(xtractrTag, ",")

	if len(xTags) > 2 {
		return false, errors.New("incorrect xtractr value(s) provided")
	}

	err := verifyValues(xTags)
	if err != nil {
		return false, err
	}

	if len(xTags) == 1 && xTags[0] == "sql" {
		return false, errors.New("sql type specified without location of parameter")
	}

	if len(xTags) == 1 {
		return false, nil
	}

	for _, word := range xTags {
		if word == "sql" {
			return true, nil
		}
	}

	return false, nil
}

func verifyValues(xTags []string) error {
	for _, word := range xTags {
		if word == "query" ||
			word == "path" ||
			word == "sql" {
			return nil
		}
	}
	return errors.New("incorrect xtractr value(s) provided")
}
