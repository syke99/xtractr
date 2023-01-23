package xtractr

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// TODO: implement tests for unsigned numbers

type TestStruct struct {
	Xtractr          string `xtractr:"-"`
	FieldOne         bool   `json:"fieldOne" xtractr:"path"`
	FieldTwo         string `json:"fieldTwo" xtractr:"path"`
	FieldThree       int    `json:"fieldThree" xtractr:"path"`
	FieldFour        int8   `json:"fieldFour" xtractr:"path"`
	FieldFive        int16  `json:"fieldFive" xtractr:"path"`
	FieldSix         int32  `json:"fieldSix" xtractr:"path"`
	FieldSeven       int64  `json:"fieldSeven" xtractr:"path"`
	FieldSeventeen   bool   `json:"fieldSeventeen" xtractr:"query"`
	FieldEighteen    string `json:"fieldEighteen" xtractr:"query"`
	FieldNineteen    int    `json:"fieldNineteen" xtractr:"query"`
	FieldTwenty      int8   `json:"fieldTwenty" xtractr:"query"`
	FieldTwentyOne   int16  `json:"fieldTwentyOne" xtractr:"query"`
	FieldTwentyTwo   int32  `json:"fieldTwentyTwo" xtractr:"query"`
	FieldTwentyThree int64  `json:"fieldTwentyThree" xtractr:"query"`
}

type TestStructTwo struct {
	Xtractr          string     `xtractr:"-"`
	FieldThirteen    float32    `json:"fieldThirteen" xtractr:"path"`
	FieldFourteen    float64    `json:"fieldFourteen" xtractr:"path"`
	FieldFifteen     complex64  `json:"fieldFifteen" xtractr:"path"`
	FieldSixteen     complex128 `json:"fieldSixteen" xtractr:"path"`
	FieldTwentyNine  float32    `json:"fieldTwentyNine" xtractr:"query"`
	FieldThirty      float64    `json:"fieldThirty" xtractr:"query"`
	FieldThirtyOne   complex64  `json:"fieldThirtyOne" xtractr:"query"`
	FieldThirtyTwo   complex128 `json:"fieldThirtyTwo" xtractr:"query"`
	FieldThirtyThree []string   `json:"fieldThirtyThree" xtractr:"query"`
}

type TestStructThree struct {
	Xtractr          string `xtractr:"-"`
	FieldEight       uint   `json:"fieldEight" xtractr:"path"`
	FieldNine        uint8  `json:"fieldNine" xtractr:"path"`
	FieldTen         uint16 `json:"fieldTen" xtractr:"path"`
	FieldEleven      uint32 `json:"fieldEleven" xtractr:"path"`
	FieldTwelve      uint64 `json:"fieldTwelve" xtractr:"path"`
	FieldTwentyFour  uint   `json:"fieldTwentyFour" xtractr:"query"`
	FieldTwentyFive  uint8  `json:"fieldTwentyFive" xtractr:"query"`
	FieldTwentySix   uint16 `json:"fieldTwentySix" xtractr:"query"`
	FieldTwentySeven uint32 `json:"fieldTwentySeven" xtractr:"query"`
	FieldTwentyEight uint64 `json:"fieldTwentyEight" xtractr:"query"`
}

const testPath = "/{fieldOne}/{fieldTwo}/{fieldThree}/{fieldFour}/{fieldFive}/{fieldSix}/{fieldSeven}"
const testPathTwo = "/{fieldThirteen}/{fieldFourteen}/{fieldFifteen}/{fieldSixteen}"
const testPathFour = "/{thisOne}"

func TestExtractParams_FirstStruct(t *testing.T) {
	path := "/true/goodbye/1/2/3/4/5/?fieldSeventeen&fieldEighteen=hello&fieldNineteen=1&fieldTwenty=2&fieldTwentyOne=3&fieldTwentyTwo=4&fieldTwentyThree=5"

	params := TestStruct{
		Xtractr: testPath,
	}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	ExtractParams(request, &params)

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
	path := "11.0/12.1/13i/14i?fieldTwentyNine=11.0&fieldThirty=12.1&fieldThirtyOne=13i&fieldThirtyTwo=14i&fieldThirtyThree=hello&fieldThirtyThree=world"

	params := TestStructTwo{
		Xtractr: testPathTwo,
	}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	ExtractParams(request, &params)

	var complexThirteen complex64 = 0 + 13i

	var complexFourteen complex128 = 0 + 14i

	strSlc := make([]string, 2)

	strSlc[0] = "hello"
	strSlc[1] = "world"

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

// TODO: fix time nested struct fields
type TestStructFour struct {
	Xtractr string         `xtractr:"-"`
	Nested  TestStructFive `xtractr:"struct"`
	//Time    time.Time      `json:"time" xtractr:"query" xtractr-time:"2006-12-01"`
}

type TestStructFive struct {
	One string `json:"thisOne" xtractr:"path"`
}

func TestExtractParams_ForthStruct(t *testing.T) {
	path := "/one?time=2020-12-02"

	params := TestStructFour{
		Xtractr: testPathFour,
	}

	//timeFormat, _ := time.Parse(strings.Split(time.RFC3339, "T")[0], "2020-08-20")

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	ExtractParams(request, &params)

	assert.Equal(t, "one", params.Nested.One)
	//assert.Equal(t, timeFormat, params.Time)
}

const testPathFive = "/{fieldOne}/{fieldThree}/{fieldFive}/{fieldSeven}/{fieldNine}/{fieldEleven}"

type SQLTestStruct struct {
	Xtractr     string          `xtractr:"-"`
	FieldOne    sql.NullBool    `json:"fieldOne" xtractr:"path,sql"`
	FieldTwo    sql.NullBool    `json:"fieldTwo" xtractr:"query,sql"`
	FieldThree  sql.NullString  `json:"fieldThree" xtractr:"path,sql"`
	FieldFour   sql.NullString  `json:"fieldFour" xtractr:"query,sql"`
	FieldFive   sql.NullInt16   `json:"fieldFive" xtractr:"path,sql"`
	FieldSix    sql.NullInt16   `json:"fieldSix" xtractr:"query,sql"`
	FieldSeven  sql.NullInt32   `json:"fieldSeven" xtractr:"path,sql"`
	FieldEight  sql.NullInt32   `json:"fieldEight" xtractr:"query,sql"`
	FieldNine   sql.NullInt64   `json:"fieldNine" xtractr:"path,sql"`
	FieldTen    sql.NullInt64   `json:"fieldTen" xtractr:"query,sql"`
	FieldEleven sql.NullFloat64 `json:"fieldEleven" xtractr:"path,sql"`
	FieldTwelve sql.NullFloat64 `json:"fieldTwelve" xtractr:"query,sql"`
}

func TestExtractParams_SQL(t *testing.T) {
	path := "/false//1/2/3/4.0?fieldTwo&fieldFour=hello&fieldSix=5&fieldEight=6&fieldTen=7&fieldTwelve=8.1"

	params := SQLTestStruct{
		Xtractr: testPathFive,
	}

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	ExtractParams(request, &params)

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
