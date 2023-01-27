package xtractr

import (
	"github.com/stretchr/testify/assert"
	"github.com/syke99/xtractr/internal/pkg/resources"
	"net/http"
	"testing"
	"time"
)

func TestExtractParams_WrongDst(t *testing.T) {
	dummyPattern := resources.DummyPattern

	req, _ := http.NewRequest("", "", nil)

	dst := "hello"
	dstPtr := &dst

	dstErr := ExtractParams(dummyPattern, req, dst)
	dstPtrErr := ExtractParams(dummyPattern, req, dstPtr)

	assert.Equal(t, resources.InvalidDst, dstErr)
	assert.Equal(t, resources.InvalidDst, dstPtrErr)
}

func TestExtractParams_NoPattern(t *testing.T) {
	dst := resources.DummyDst

	req, _ := http.NewRequest("", "", nil)

	err := ExtractParams("", req, &dst)

	assert.Equal(t, resources.MissingPattern, err)
}

func TestExtractParams_NilRequest(t *testing.T) {
	dummyPattern := resources.DummyPattern

	dst := resources.DummyDst

	err := ExtractParams(dummyPattern, nil, &dst)

	assert.Equal(t, resources.MissingIncomingRequest, err)
}

func TestExtractParams_PathParamParseError(t *testing.T) {
	dst := resources.DummyDst

	req, _ := http.NewRequest("", "/hello", nil)

	err := ExtractParams("/hello/world", req, &dst)

	assert.Equal(t, resources.PathParseErr, err)
}

func TestExtractParams_FirstStruct(t *testing.T) {
	path := resources.PathOne

	params := resources.TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(resources.TestPathOne, request, &params)

	assert.NoError(t, err)
	assert.Equal(t, true, params.FieldOne)
	assert.Equal(t, "goodbye", params.FieldTwo)
	assert.Equal(t, 1, params.FieldThree)
	assert.Equal(t, int8(2), params.FieldFour)
	assert.Equal(t, int16(3), params.FieldFive)
	assert.Equal(t, int32(4), params.FieldSix)
	assert.Equal(t, int64(5), params.FieldSeven)
	assert.Equal(t, true, params.FieldSeventeen)
	assert.Equal(t, "hello", params.FieldEighteen)
	assert.Equal(t, 1, params.FieldNineteen)
	assert.Equal(t, int8(2), params.FieldTwenty)
	assert.Equal(t, int16(3), params.FieldTwentyOne)
	assert.Equal(t, int32(4), params.FieldTwentyTwo)
	assert.Equal(t, int64(5), params.FieldTwentyThree)
}

func TestExtractParams_SecondStuct(t *testing.T) {
	path := resources.PathTwo

	params := resources.TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(resources.TestPathTwo, request, &params)

	var complexThirteen complex64 = 0 + 13i

	var complexFourteen complex128 = 0 + 14i

	strSlc := make([]string, 2)

	strSlc[0] = "hello"
	strSlc[1] = "world"

	assert.NoError(t, err)
	assert.Equal(t, float32(11.0), params.FieldThirteen)
	assert.Equal(t, 12.1, params.FieldFourteen)
	assert.Equal(t, complexThirteen, params.FieldFifteen)
	assert.Equal(t, complexFourteen, params.FieldSixteen)
	assert.Equal(t, float32(11.0), params.FieldTwentyNine)
	assert.Equal(t, 12.1, params.FieldThirty)
	assert.Equal(t, complexThirteen, params.FieldThirtyOne)
	assert.Equal(t, complexFourteen, params.FieldThirtyTwo)
	for i, word := range params.FieldThirtyThree {
		assert.Equal(t, strSlc[i], word)
	}
}

func TestExtractParams_ThirdStruct(t *testing.T) {
	path := resources.PathThree

	params := resources.TestStructThree{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(resources.TestPathThree, request, &params)

	assert.NoError(t, err)
	assert.Equal(t, uint(8), params.FieldEight)
	assert.Equal(t, uint8(9), params.FieldNine)
	assert.Equal(t, uint16(10), params.FieldTen)
	assert.Equal(t, uint32(11), params.FieldEleven)
	assert.Equal(t, uint64(12), params.FieldTwelve)
	assert.Equal(t, uint(24), params.FieldTwentyFour)
	assert.Equal(t, uint8(25), params.FieldTwentyFive)
	assert.Equal(t, uint16(26), params.FieldTwentySix)
	assert.Equal(t, uint32(27), params.FieldTwentySeven)
	assert.Equal(t, uint64(28), params.FieldTwentyEight)
}

func TestExtractParams_ForthStruct(t *testing.T) {
	path := resources.PathFour

	params := resources.TestStructFour{}

	tm := time.Date(2020, time.Month(12), 02, 0, 0, 0, 0, time.UTC)

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(resources.TestPathFour, request, &params)

	assert.NoError(t, err)
	assert.Equal(t, "one", params.Nested.One)
	assert.Equal(t, tm, params.Time)
}

func TestExtractParams_SQL(t *testing.T) {
	path := resources.PathFive

	params := resources.SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(resources.TestPathFive, request, &params)

	assert.NoError(t, err)
	// sql.NullTime
	t13 := time.Date(2022, time.Month(12), 01, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, t13, params.FieldThirteen.Time)
	assert.Equal(t, true, params.FieldThirteen.Valid)
	t14 := time.Date(2020, time.Month(12), 04, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, t14, params.FieldFourteen.Time)
	assert.Equal(t, true, params.FieldFourteen.Valid)
	// sql.NullBool
	assert.Equal(t, false, params.FieldOne.Bool)
	assert.Equal(t, true, params.FieldOne.Valid)
	assert.Equal(t, true, params.FieldTwo.Bool)
	assert.Equal(t, true, params.FieldTwo.Valid)
	// sql.NullString
	assert.Equal(t, "", params.FieldThree.String)
	assert.Equal(t, false, params.FieldThree.Valid)
	assert.Equal(t, "hello", params.FieldFour.String)
	assert.Equal(t, true, params.FieldFour.Valid)
	// sql.NullInt16
	assert.Equal(t, int16(1), params.FieldFive.Int16)
	assert.Equal(t, true, params.FieldFive.Valid)
	assert.Equal(t, int16(5), params.FieldSix.Int16)
	assert.Equal(t, true, params.FieldSix.Valid)
	// sql.NullInt32
	assert.Equal(t, int32(2), params.FieldSeven.Int32)
	assert.Equal(t, true, params.FieldSeven.Valid)
	assert.Equal(t, int32(6), params.FieldEight.Int32)
	assert.Equal(t, true, params.FieldEight.Valid)
	// sql.NullInt64
	assert.Equal(t, int64(3), params.FieldNine.Int64)
	assert.Equal(t, true, params.FieldNine.Valid)
	assert.Equal(t, int64(7), params.FieldTen.Int64)
	assert.Equal(t, true, params.FieldTen.Valid)
	// sql.NullFloat64
	assert.Equal(t, float64(4.0), params.FieldEleven.Float64)
	assert.Equal(t, true, params.FieldEleven.Valid)
	assert.Equal(t, float64(8.1), params.FieldTwelve.Float64)
	assert.Equal(t, true, params.FieldTwelve.Valid)
}
