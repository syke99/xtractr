package resources

import (
	"database/sql"
	"time"
)

const DummyPattern = "/{hello}"

var DummyDst = struct{}{}

const TestPathOne = "/{fieldOne}/{fieldTwo}/{fieldThree}/{fieldFour}/{fieldFive}/{fieldSix}/{fieldSeven}"
const PathOne = "/true/goodbye/1/2/3/4/5/?fieldSeventeen&fieldEighteen=hello&fieldNineteen=1&fieldTwenty=2&fieldTwentyOne=3&fieldTwentyTwo=4&fieldTwentyThree=5"

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

const TestPathTwo = "/{fieldThirteen}/{fieldFourteen}/{fieldFifteen}/{fieldSixteen}"
const PathTwo = "/11.0/12.1/13i/14i?fieldTwentyNine=11.0&fieldThirty=12.1&fieldThirtyOne=13i&fieldThirtyTwo=14i&fieldThirtyThree=hello&fieldThirtyThree=world"

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

const TestPathThree = "/{fieldEight}/{fieldNine}/{fieldTen}/{fieldEleven}/{fieldTwelve}"
const PathThree = "/8/9/10/11/12?fieldTwentyFour=24&fieldTwentyFive=25&fieldTwentySix=26&fieldTwentySeven=27&fieldTwentyEight=28"

type TestStructThree struct {
	FieldEight       uint   `xtractr:"path" xtractr-param:"fieldEight"`
	FieldNine        uint8  `xtractr:"path" xtractr-param:"fieldNine"`
	FieldTen         uint16 `xtractr:"path" xtractr-param:"fieldTen"`
	FieldEleven      uint32 `xtractr:"path" xtractr-param:"fieldEleven"`
	FieldTwelve      uint64 `xtractr:"path" xtractr-param:"fieldTwelve"`
	FieldTwentyFour  uint   `xtractr:"query" xtractr-param:"fieldTwentyFour"`
	FieldTwentyFive  uint8  `xtractr:"query" xtractr-param:"fieldTwentyFive"`
	FieldTwentySix   uint16 `xtractr:"query" xtractr-param:"fieldTwentySix"`
	FieldTwentySeven uint32 `xtractr:"query" xtractr-param:"fieldTwentySeven"`
	FieldTwentyEight uint64 `xtractr:"query" xtractr-param:"fieldTwentyEight"`
}

const TestPathFour = "/{thisOne}"
const PathFour = "/one?time=2020-12-02"

type TestStructFour struct {
	Nested TestStructFive `xtractr:"struct"`
	Time   time.Time      `xtractr:"query" xtractr-param:"time" xtractr-time:"ISO8601"`
}

type TestStructFive struct {
	One string `xtractr:"path" xtractr-param:"thisOne"`
}

const TestPathFive = "/{fieldOne}/{fieldThree}/{fieldFive}/{fieldSeven}/{fieldNine}/{fieldEleven}/{fieldThirteen}"
const PathFive = "/false//1/2/3/4.0/2022-12-01?fieldFourteen=2020-12-04&fieldTwo&fieldFour=hello&fieldSix=5&fieldEight=6&fieldTen=7&fieldTwelve=8.1"

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
