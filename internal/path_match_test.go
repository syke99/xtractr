package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSanitizePaths_BeginningSlashPattern(t *testing.T) {
	pattern, _ := SanitizePaths("/pattern", "/reqPath")

	assert.Equal(t, "pattern", pattern)
}

func TestSanitizePaths_EndingSlashPattern(t *testing.T) {
	pattern, _ := SanitizePaths("pattern/", "/reqPath")

	assert.Equal(t, "pattern", pattern)
}

func TestSanitizePaths_BeginningSlashReqPath(t *testing.T) {
	_, reqPath := SanitizePaths("/pattern", "/reqPath")

	assert.Equal(t, "reqPath", reqPath)
}

func TestSanitizePaths_EndingSlashReqPath(t *testing.T) {
	_, reqPath := SanitizePaths("/pattern", "reqPath/")

	assert.Equal(t, "reqPath", reqPath)
}

func TestGetMatchedPathParams(t *testing.T) {
	toMatch := "/{first}/{second}"

	reqPath := "/one/two"

	toMatch, reqPath = SanitizePaths(toMatch, reqPath)

	p := GetMatchedPathParams(toMatch, reqPath)

	assert.Equal(t, "one", p["first"])
	assert.Equal(t, "two", p["second"])
}

func TestGetMatchedPathParams_MisMatchedLengthPattern(t *testing.T) {
	toMatch := "/{first}/{second}"

	reqPath := "/one"

	toMatch, reqPath = SanitizePaths(toMatch, reqPath)

	p := GetMatchedPathParams(toMatch, reqPath)

	assert.Nil(t, p)
}

func TestGetMatchedPathParams_MisMatchedLengthReqPath(t *testing.T) {
	toMatch := "/{first}"

	reqPath := "/one/two"

	toMatch, reqPath = SanitizePaths(toMatch, reqPath)

	p := GetMatchedPathParams(toMatch, reqPath)

	assert.Nil(t, p)
}
