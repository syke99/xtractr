package internal

import (
	"errors"
	"fmt"
	"github.com/syke99/xtractr/internal/unmarshal"
	"net/http"
	"reflect"
	"strings"
)

func SanitizePaths(pattern string, reqPath string) (string, string) {
	if pattern[:1] == "/" {
		pattern = pattern[1:]
	}

	if pattern[len(pattern)-1:] == "/" {
		pattern = pattern[:len(pattern)-1]
	}

	if reqPath[:1] == "/" {
		reqPath = reqPath[1:]
	}

	if reqPath[len(reqPath)-1:] == "/" {
		reqPath = reqPath[:len(reqPath)-1]
	}

	return pattern, reqPath
}

func GetMatchedPathParams(toMatch string, requested string) map[string]string {

	matchPathParts := strings.Split(toMatch, "/")

	reqPathParts := strings.Split(requested, "/")

	if len(matchPathParts) != len(reqPathParts) {
		return nil
	}

	m := make(map[string]string)

	for i, reqPart := range reqPathParts {
		if matchPathParts[i][:1] == "{" &&
			matchPathParts[i][len(matchPathParts[i])-1:] == "}" {

			strippedMatch := matchPathParts[i][1 : len(matchPathParts[i])-1]

			m[strippedMatch] = reqPart
		}
	}

	return m
}

func Unmarshal(request *http.Request, str reflect.Value, pathParams map[string]string) error {
	err := unmarshal.Unmarshal(request, str, pathParams)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed to unmarshal fields (%s)", err.Error()))
	}
	return err
}
