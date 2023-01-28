package unmarshal

import (
	"github.com/stretchr/testify/assert"
	"github.com/syke99/xtractr/internal"
	"github.com/syke99/xtractr/internal/pkg/resources"
	"net/http"
	"reflect"
	"testing"
)

func TestUnmarshal_NoError_Regular(t *testing.T) {
	path := resources.PathOne

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.NoError(t, err)
}

func TestUnmarshal_NoError_Sql(t *testing.T) {
	path := resources.PathFive

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.NoError(t, err)
}

func TestUnmarshal_NoError_NestedStruct(t *testing.T) {
	path := resources.PathFour

	params := resources.TestStructFour{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFour, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.NoError(t, err)
}
