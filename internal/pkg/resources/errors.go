package resources

import "errors"

var MissingPattern error = errors.New("no pattern to match request path against provided")
var MissingIncomingRequest error = errors.New("no incoming *http.Request provided")
var InvalidDst error = errors.New("dst provided is not a valid pointer to a struct")
var PathParseErr error = errors.New("error parsing path for parameters")
