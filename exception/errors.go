package exception

import "errors"

// AlreadyExists Resource already exists
var AlreadyExists = errors.New("resource already exists")
// NotFound Resource was not found
var NotFound = errors.New("resource not found")
// RequiredField Aggregate field is missing
var RequiredField = errors.New("missing required field")
// FieldFormat Aggregate field format is invalid
var FieldFormat = errors.New("field has an invalid format")
// FieldRange Aggregate field range is invalid
var FieldRange = errors.New("field has an invalid range")
// Invalid Aggregate/Entity/... is invalid
var Invalid = errors.New("invalid field")

/* Description(s) */

var AlreadyExistsDescription = "%v already exists"
var NotFoundDescription = "%v not found"
var RequiredFieldDescription = "missing required field: %v"
var FieldFormatDescription = "field %v has an invalid format, expected %v"
var FieldRangeDescription = "field %v is out of range [%v, %v)"
