package resources

import "errors"

var InvalidDst error = errors.New("dst provided is not a valid pointer to a struct")
var PathParseErr error = errors.New("error parsing path for parameters")
