package unmarshal

import (
	"errors"
	"strings"
)

func DetermineSQL(xtractrTag string) (bool, error) {
	xTags := strings.Split(xtractrTag, ",")

	err := verifyValues(xTags)
	if err != nil ||
		len(xTags) > 2 {
		return false, err
	}

	if len(xTags) == 1 && xTags[0] == "sql" {
		return false, errors.New("")
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
	return errors.New("")
}
