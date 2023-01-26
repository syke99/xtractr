package xtractr

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

type TestStruct struct {
	FieldOne         bool   `xtractr:"path" xtractr-param:"fieldOne"`
	FieldTwo         string `xtractr:"path" xtractr-param:"fieldTwo"`
	FieldThree       int    `xtractr:"path" xtractr-param:"fieldThree"`
	FieldFour        int8   `xtractr:"path" xtractr-param:"fieldFour"`
	FieldFive        int16  `xtractr:"path" xtractr-param:"fieldFive"`
	FieldSix         int32  `xtractr:"path" xtractr-param:"fieldSix"`
	FieldSeven       int64  `xtractr:"path" xtractr-param:"fieldSeven"`
	FieldSeventeen   bool   `xtractr:"query" xtractr-param:"fieldSeventeen"`
	FieldEighteen    string `xtractr:"query" xtractr-param:"fieldEighteen"`
	FieldNineteen    int    `xtractr:"query" xtractr-param:"fieldNineteen"`
	FieldTwenty      int8   `xtractr:"query" xtractr-param:"fieldTwenty"`
	FieldTwentyOne   int16  `xtractr:"query" xtractr-param:"fieldTwentyOne"`
	FieldTwentyTwo   int32  `xtractr:"query" xtractr-param:"fieldTwentyTwo"`
	FieldTwentyThree int64  `xtractr:"query" xtractr-param:"fieldTwentyThree"`
}

const testPath = "/{fieldOne}/{fieldTwo}/{fieldThree}/{fieldFour}/{fieldFive}/{fieldSix}/{fieldSeven}"

func TestExtractParams_FirstStruct(t *testing.T) {
	path := "/true/goodbye/1/2/3/4/5/?fieldSeventeen&fieldEighteen=hello&fieldNineteen=1&fieldTwenty=2&fieldTwentyOne=3&fieldTwentyTwo=4&fieldTwentyThree=5"

	params := TestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(testPath, request, &params)

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

type TestStructTwo struct {
	FieldThirteen    float32    `xtractr:"path" xtractr-param:"fieldThirteen"`
	FieldFourteen    float64    `xtractr:"path" xtractr-param:"fieldFourteen"`
	FieldFifteen     complex64  `xtractr:"path" xtractr-param:"fieldFifteen"`
	FieldSixteen     complex128 `xtractr:"path" xtractr-param:"fieldSixteen"`
	FieldTwentyNine  float32    `xtractr:"query" xtractr-param:"fieldTwentyNine"`
	FieldThirty      float64    `xtractr:"query" xtractr-param:"fieldThirty"`
	FieldThirtyOne   complex64  `xtractr:"query" xtractr-param:"fieldThirtyOne"`
	FieldThirtyTwo   complex128 `xtractr:"query" xtractr-param:"fieldThirtyTwo"`
	FieldThirtyThree []string   `xtractr:"query" xtractr-param:"fieldThirtyThree"`
}

const testPathTwo = "/{fieldThirteen}/{fieldFourteen}/{fieldFifteen}/{fieldSixteen}"

func TestExtractParams_SecondStuct(t *testing.T) {
	path := "11.0/12.1/13i/14i?fieldTwentyNine=11.0&fieldThirty=12.1&fieldThirtyOne=13i&fieldThirtyTwo=14i&fieldThirtyThree=hello&fieldThirtyThree=world"

	params := TestStructTwo{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(testPathTwo, request, &params)

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

// TODO: implement tests for unsigned numbers
//type TestStructThree struct {
//	FieldEight       uint   `xtractr:"path" xtractr-param:"fieldEight"`
//	FieldNine        uint8  `xtractr:"path" xtractr-param:"fieldNine"`
//	FieldTen         uint16 `xtractr:"path" xtractr-param:"fieldTen"`
//	FieldEleven      uint32 `xtractr:"path" xtractr-param:"fieldEleven"`
//	FieldTwelve      uint64 `xtractr:"path" xtractr-param:"fieldTwelve"`
//	FieldTwentyFour  uint   `xtractr:"query" xtractr-param:"fieldTwentyFour"`
//	FieldTwentyFive  uint8  `xtractr:"query" xtractr-param:"fieldTwentyFive"`
//	FieldTwentySix   uint16 `xtractr:"query" xtractr-param:"fieldTwentySix"`
//	FieldTwentySeven uint32 `xtractr:"query" xtractr-param:"fieldTwentySeven"`
//	FieldTwentyEight uint64 `xtractr:"query" xtractr-param:"fieldTwentyEight"`
//}

type TestStructFour struct {
	Nested TestStructFive `xtractr:"struct"`
	Time   time.Time      `xtractr:"query" xtractr-param:"time" xtractr-time:"ISO8601"`
}

type TestStructFive struct {
	One string `xtractr:"path" xtractr-param:"thisOne"`
}

const testPathFour = "/{thisOne}"

func TestExtractParams_ForthStruct(t *testing.T) {
	path := "/one?time=2020-12-02"

	params := TestStructFour{}

	tm := time.Date(2020, time.Month(12), 02, 0, 0, 0, 0, time.UTC)

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(testPathFour, request, &params)

	assert.NoError(t, err)
	assert.Equal(t, "one", params.Nested.One)
	assert.Equal(t, tm, params.Time)
}

const testPathFive = "/{fieldOne}/{fieldThree}/{fieldFive}/{fieldSeven}/{fieldNine}/{fieldEleven}/{fieldThirteen}"

type SQLTestStruct struct {
	FieldOne      sql.NullBool    `xtractr:"path,sql" xtractr-param:"fieldOne"`
	FieldTwo      sql.NullBool    `xtractr:"query,sql" xtractr-param:"fieldTwo"`
	FieldThree    sql.NullString  `xtractr:"path,sql" xtractr-param:"fieldThree"`
	FieldFour     sql.NullString  `xtractr:"query,sql" xtractr-param:"fieldFour"`
	FieldFive     sql.NullInt16   `xtractr:"path,sql" xtractr-param:"fieldFive"`
	FieldSix      sql.NullInt16   `xtractr:"query,sql" xtractr-param:"fieldSix"`
	FieldSeven    sql.NullInt32   `xtractr:"path,sql" xtractr-param:"fieldSeven"`
	FieldEight    sql.NullInt32   `xtractr:"query,sql" xtractr-param:"fieldEight"`
	FieldNine     sql.NullInt64   `xtractr:"path,sql" xtractr-param:"fieldNine"`
	FieldTen      sql.NullInt64   `xtractr:"query,sql" xtractr-param:"fieldTen"`
	FieldEleven   sql.NullFloat64 `xtractr:"path,sql" xtractr-param:"fieldEleven"`
	FieldTwelve   sql.NullFloat64 `xtractr:"query,sql" xtractr-param:"fieldTwelve"`
	FieldThirteen sql.NullTime    `xtractr:"path,sql" xtractr-param:"fieldThirteen" xtractr-time:"ISO8601"`
	FieldFourteen sql.NullTime    `xtractr:"query,sql" xtractr-param:"fieldFourteen" xtractr-time:"ISO8601"`
}

func TestExtractParams_SQL(t *testing.T) {
	path := "/false//1/2/3/4.0/2022-12-01?fieldFourteen=2020-12-04&fieldTwo&fieldFour=hello&fieldSix=5&fieldEight=6&fieldTen=7&fieldTwelve=8.1"

	params := SQLTestStruct{}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	err := ExtractParams(testPathFive, request, &params)

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
