package unmarshal

import (
	"github.com/stretchr/testify/assert"
	"github.com/syke99/xtractr/internal"
	"github.com/syke99/xtractr/internal/pkg/resources"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestUnmarshal_NoError_Basic(t *testing.T) {
	path := resources.PathOne

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.NoError(t, err)
}

func TestUnmarshal_ErrorBool_Basic_Path(t *testing.T) {
	path := resources.PathOne

	pathParts := strings.Split(path, "/")

	pathParts[1] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicBoolError, err)
}

func TestUnmarshal_ErrorInt_Basic_Path(t *testing.T) {
	path := resources.PathOne

	pathParts := strings.Split(path, "/")

	pathParts[3] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicIntError, err)
}

func TestUnmarshal_ErrorInt8_Basic_Path(t *testing.T) {
	path := resources.PathOne

	pathParts := strings.Split(path, "/")

	pathParts[4] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicInt8Error, err)
}

func TestUnmarshal_ErrorInt16_Basic_Path(t *testing.T) {
	path := resources.PathOne

	pathParts := strings.Split(path, "/")

	pathParts[5] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicInt16Error, err)
}

func TestUnmarshal_ErrorInt32_Basic_Path(t *testing.T) {
	path := resources.PathOne

	pathParts := strings.Split(path, "/")

	pathParts[6] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicInt32Error, err)
}

func TestUnmarshal_ErrorInt64_Basic_Path(t *testing.T) {
	path := resources.PathOne

	pathParts := strings.Split(path, "/")

	pathParts[7] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicInt64Error, err)
}

func TestUnmarshal_ErrorUint_Basic_Path(t *testing.T) {
	path := resources.PathThree

	pathParts := strings.Split(path, "/")

	pathParts[1] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicUintError, err)
}

func TestUnmarshal_ErrorUint8_Basic_Path(t *testing.T) {
	path := resources.PathThree

	pathParts := strings.Split(path, "/")

	pathParts[2] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicUint8Error, err)
}

func TestUnmarshal_ErrorUint16_Basic_Path(t *testing.T) {
	path := resources.PathThree

	pathParts := strings.Split(path, "/")

	pathParts[3] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicUint16Error, err)
}

func TestUnmarshal_ErrorUint32_Basic_Path(t *testing.T) {
	path := resources.PathThree

	pathParts := strings.Split(path, "/")

	pathParts[4] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicUint32Error, err)
}

func TestUnmarshal_ErrorUint64_Basic_Path(t *testing.T) {
	path := resources.PathThree

	pathParts := strings.Split(path, "/")

	splitFifth := strings.Split(pathParts[5], "?")

	splitFifth[0] = "hello"

	pathParts[5] = strings.Join(splitFifth, "?")

	path = strings.Join(pathParts, "/")

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicUint64Error, err)
}

func TestUnmarshal_ErrorFloat32_Basic_Path(t *testing.T) {
	path := resources.PathTwo

	pathParts := strings.Split(path, "/")

	pathParts[1] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathTwo, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicFloat32Error, err)
}

func TestUnmarshal_ErrorFloat64_Basic_Path(t *testing.T) {
	path := resources.PathTwo

	pathParts := strings.Split(path, "/")

	pathParts[2] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathTwo, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicFloat64Error, err)
}

func TestUnmarshal_ErrorComplex64_Basic_Path(t *testing.T) {
	path := resources.PathTwo

	pathParts := strings.Split(path, "/")

	pathParts[3] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathTwo, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicComplex64Error, err)
}

func TestUnmarshal_ErrorComplex128_Basic_Path(t *testing.T) {
	path := resources.PathTwo

	pathParts := strings.Split(path, "/")

	splitFourth := strings.Split(pathParts[4], "?")

	splitFourth[0] = "hello"

	pathParts[4] = strings.Join(splitFourth, "?")

	path = strings.Join(pathParts, "/")

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathTwo, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathBasicComplex128Error, err)
}

func TestUnmarshal_ErrorInt_Basic_Query(t *testing.T) {
	path := resources.PathOne

	path = strings.Replace(path, "fieldNineteen=1", "fieldNineteen=hello", -1)

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicIntError, err)
}

func TestUnmarshal_ErrorInt8_Basic_Query(t *testing.T) {
	path := resources.PathOne

	path = strings.Replace(path, "fieldTwenty=2", "fieldTwenty=hello", -1)

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicInt8Error, err)
}

func TestUnmarshal_ErrorInt16_Basic_Query(t *testing.T) {
	path := resources.PathOne

	path = strings.Replace(path, "fieldTwentyOne=3", "fieldTwentyOne=hello", -1)

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicInt16Error, err)
}

func TestUnmarshal_ErrorInt32_Basic_Query(t *testing.T) {
	path := resources.PathOne

	path = strings.Replace(path, "fieldTwentyTwo=4", "fieldTwentyTwo=hello", -1)

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicInt32Error, err)
}

func TestUnmarshal_ErrorInt64_Basic_Query(t *testing.T) {
	path := resources.PathOne

	path = strings.Replace(path, "fieldTwentyThree=5", "fieldTwentyThree=hello", -1)

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathOne, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicInt64Error, err)
}

func TestUnmarshal_ErrorUin_Basic_Query(t *testing.T) {
	path := resources.PathThree

	path = strings.Replace(path, "fieldTwentyFour=24", "fieldTwentyFour=hello", -1)

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicUintError, err)
}

func TestUnmarshal_ErrorUint8_Basic_Query(t *testing.T) {
	path := resources.PathThree

	path = strings.Replace(path, "fieldTwentyFive=25", "fieldTwentyFive=hello", -1)

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicUint8Error, err)
}

func TestUnmarshal_ErrorUint16_Basic_Query(t *testing.T) {
	path := resources.PathThree

	path = strings.Replace(path, "fieldTwentySix=26", "fieldTwentySix=hello", -1)

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicUint16Error, err)
}

func TestUnmarshal_ErrorUint32_Basic_Query(t *testing.T) {
	path := resources.PathThree

	path = strings.Replace(path, "fieldTwentySeven=27", "fieldTwentySeven=hello", -1)

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicUint32Error, err)
}

func TestUnmarshal_ErrorUint64_Basic_Query(t *testing.T) {
	path := resources.PathThree

	path = strings.Replace(path, "fieldTwentyEight=28", "fieldTwentyEight=hello", -1)

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathThree, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicUint64Error, err)
}

func TestUnmarshal_ErrorFloat32_Basic_Query(t *testing.T) {
	path := resources.PathTwo

	path = strings.Replace(path, "fieldTwentyNine=11.0", "fieldTwentyNine=hello", -1)

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathTwo, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicFloat32Error, err)
}

func TestUnmarshal_ErrorFloat64_Basic_Query(t *testing.T) {
	path := resources.PathTwo

	path = strings.Replace(path, "fieldThirty=12.1", "fieldThirty=hello", -1)

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathTwo, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicFloat64Error, err)
}

func TestUnmarshal_ErrorComplex64_Basic_Query(t *testing.T) {
	path := resources.PathTwo

	path = strings.Replace(path, "fieldThirtyOne=13i", "fieldThirtyOne=hello", -1)

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathTwo, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicComplex64Error, err)
}

func TestUnmarshal_ErrorComplex128_Basic_Query(t *testing.T) {
	path := resources.PathTwo

	path = strings.Replace(path, "fieldThirtyTwo=14i", "fieldThirtyTwo=hello", -1)

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathTwo, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQueryBasicComplex128Error, err)
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

func TestUnmarshal_ErrorBool_Sql_Path(t *testing.T) {
	path := resources.PathFive

	pathParts := strings.Split(path, "/")

	pathParts[1] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathSqlBoolError, err)
}

func TestUnmarshal_ErrorInt16_Sql_Path(t *testing.T) {
	path := resources.PathFive

	pathParts := strings.Split(path, "/")

	pathParts[3] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathSqlInt16Error, err)
}

func TestUnmarshal_ErrorInt32_Sql_Path(t *testing.T) {
	path := resources.PathFive

	pathParts := strings.Split(path, "/")

	pathParts[4] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathSqlInt32Error, err)
}

func TestUnmarshal_ErrorInt64_Sql_Path(t *testing.T) {
	path := resources.PathFive

	pathParts := strings.Split(path, "/")

	pathParts[5] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathSqlInt64Error, err)
}

func TestUnmarshal_ErrorFloat64_Sql_Path(t *testing.T) {
	path := resources.PathFive

	pathParts := strings.Split(path, "/")

	pathParts[6] = "hello"

	path = strings.Join(pathParts, "/")

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParsePathSqlFloat64Error, err)
}

func TestUnmarshal_ErrorBool_Sql_Query(t *testing.T) {
	path := resources.PathFive

	path = strings.Replace(path, "fieldTwo", "fieldTwo=hello", -1)

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQuerySqlBoolError, err)
}

func TestUnmarshal_ErrorInt16_Sql_Query(t *testing.T) {
	path := resources.PathFive

	path = strings.Replace(path, "fieldSix=5", "fieldSix=hello", -1)

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQuerySqlInt16Error, err)
}

func TestUnmarshal_ErrorInt32_Sql_Query(t *testing.T) {
	path := resources.PathFive

	path = strings.Replace(path, "fieldEight=6", "fieldEight=hello", -1)

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQuerySqlInt32Error, err)
}

func TestUnmarshal_ErrorInt64_Sql_Query(t *testing.T) {
	path := resources.PathFive

	path = strings.Replace(path, "fieldTen=7", "fieldTen=hello", -1)

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQuerySqlInt64Error, err)
}

func TestUnmarshal_ErrorFloat64_Sql_Query(t *testing.T) {
	path := resources.PathFive

	path = strings.Replace(path, "fieldTwelve=8.1", "fieldTwelve=hello", -1)

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFive, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	err := Unmarshal(request.URL.Query(), reflect.ValueOf(&params), pathParams)

	assert.Error(t, err)
	assert.Equal(t, resources.ParseQuerySqlFloat64Error, err)
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

func TestRecurseIntoNestedStruct(t *testing.T) {
	path := resources.PathFour

	params := resources.TestStructFour{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	tPath, rPath := internal.SanitizePaths(resources.TestPathFour, request.URL.EscapedPath())

	pathParams := internal.GetMatchedPathParams(tPath, rPath)

	elem := reflect.ValueOf(&params).Elem()

	ptr, err := recurseIntoNestedStruct(elem.Type().Field(0), request.URL.Query(), pathParams)

	elem.Field(0).Set(reflect.ValueOf(ptr.Elem().Interface()))

	assert.NoError(t, err)
	assert.Equal(t, "one", params.Nested.One)
}
