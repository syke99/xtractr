package xtractr

import (
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
