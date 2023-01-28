package resources

import "errors"

var MissingPattern error = errors.New("no pattern to match request path against provided")
var MissingIncomingRequest error = errors.New("no incoming *http.Request provided")
var InvalidDst error = errors.New("dst provided is not a valid pointer to a struct")
var PathParseErr error = errors.New("error parsing path for parameters")

// Basic
//
// path
var ParsePathBasicBoolError error = errors.New("field: FieldOne, error: strconv.ParseBool: parsing \"hello\": invalid syntax")
var ParsePathBasicIntError error = errors.New("field: FieldThree, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParsePathBasicInt8Error error = errors.New("field: FieldFour, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParsePathBasicInt16Error error = errors.New("field: FieldFive, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParsePathBasicInt32Error error = errors.New("field: FieldSix, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParsePathBasicInt64Error error = errors.New("field: FieldSeven, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParsePathBasicUintError error = errors.New("field: FieldEight, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParsePathBasicUint8Error error = errors.New("field: FieldNine, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParsePathBasicUint16Error error = errors.New("field: FieldTen, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParsePathBasicUint32Error error = errors.New("field: FieldEleven, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParsePathBasicUint64Error error = errors.New("field: FieldTwelve, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParsePathBasicFloat32Error error = errors.New("field: FieldThirteen, error: strconv.ParseFloat: parsing \"hello\": invalid syntax")
var ParsePathBasicFloat64Error error = errors.New("field: FieldFourteen, error: strconv.ParseFloat: parsing \"hello\": invalid syntax")
var ParsePathBasicComplex64Error error = errors.New("field: FieldFifteen, error: strconv.ParseComplex: parsing \"hello\": invalid syntax")
var ParsePathBasicComplex128Error error = errors.New("field: FieldSixteen, error: strconv.ParseComplex: parsing \"hello\": invalid syntax")

// query
var ParseQueryBasicIntError error = errors.New("field: FieldNineteen, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParseQueryBasicInt8Error error = errors.New("field: FieldTwenty, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParseQueryBasicInt16Error error = errors.New("field: FieldTwentyOne, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParseQueryBasicInt32Error error = errors.New("field: FieldTwentyTwo, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParseQueryBasicInt64Error error = errors.New("field: FieldTwentyThree, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParseQueryBasicUintError error = errors.New("field: FieldTwentyFour, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParseQueryBasicUint8Error error = errors.New("field: FieldTwentyFive, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParseQueryBasicUint16Error error = errors.New("field: FieldTwentySix, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParseQueryBasicUint32Error error = errors.New("field: FieldTwentySeven, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParseQueryBasicUint64Error error = errors.New("field: FieldTwentyEight, error: strconv.ParseUint: parsing \"hello\": invalid syntax")
var ParseQueryBasicFloat32Error error = errors.New("field: FieldTwentyNine, error: strconv.ParseFloat: parsing \"hello\": invalid syntax")
var ParseQueryBasicFloat64Error error = errors.New("field: FieldThirty, error: strconv.ParseFloat: parsing \"hello\": invalid syntax")
var ParseQueryBasicComplex64Error error = errors.New("field: FieldThirtyOne, error: strconv.ParseComplex: parsing \"hello\": invalid syntax")
var ParseQueryBasicComplex128Error error = errors.New("field: FieldThirtyTwo, error: strconv.ParseComplex: parsing \"hello\": invalid syntax")

// Sql
//
// path
var ParsePathSqlBoolError error = errors.New("field: FieldOne, error: strconv.ParseBool: parsing \"hello\": invalid syntax")
var ParsePathSqlInt16Error error = errors.New("field: FieldFive, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParsePathSqlInt32Error error = errors.New("field: FieldSeven, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParsePathSqlInt64Error error = errors.New("field: FieldNine, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParsePathSqlFloat64Error error = errors.New("field: FieldEleven, error: strconv.ParseFloat: parsing \"hello\": invalid syntax")

// query
var ParseQuerySqlBoolError error = errors.New("field: FieldTwo, error: strconv.ParseBool: parsing \"hello\": invalid syntax")
var ParseQuerySqlInt16Error error = errors.New("field: FieldSix, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParseQuerySqlInt32Error error = errors.New("field: FieldEight, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParseQuerySqlInt64Error error = errors.New("field: FieldTen, error: strconv.ParseInt: parsing \"hello\": invalid syntax")
var ParseQuerySqlFloat64Error error = errors.New("field: FieldTwelve, error: strconv.ParseFloat: parsing \"hello\": invalid syntax")
