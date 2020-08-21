package exception

import (
	"fmt"
	"strings"
)

const format = "%w#%s"

// NewAlreadyExists Create and wrap an AlreadyExists exception
//
// Usage
//
// err := NewAlreadyExists("foo") -> "resource already exists#foo already exists"
//
// GetDescription(err) -> "foo already exists"
func NewAlreadyExists(field string) error {
	return fmt.Errorf(format, AlreadyExists, fmt.Sprintf(AlreadyExistsDescription, field))
}

// NewNotFound Create and wrap an NotFound exception
//
// Usage
//
// err := NewNotFound("foo") -> "resource not found#foo not found"
//
// GetDescription(err) -> "foo not found"
func NewNotFound(field string) error {
	return fmt.Errorf(format, NotFound, fmt.Sprintf(NotFoundDescription, field))
}

// NewRequiredField Create and wrap an RequiredField exception
//
// Usage
//
// err := NewRequiredField("foo") -> "missing required field#missing required field: foo"
//
// GetDescription(err) -> "missing required field: foo"
func NewRequiredField(field string) error {
	return fmt.Errorf(format, RequiredField, fmt.Sprintf(RequiredFieldDescription, field))
}

// NewFieldFormat Create and wrap an FieldFormat exception
//
// Usage
//
// err := NewFieldFormat("foo", "integer") -> "field has an invalid format#field foo has an invalid format, expected integer"
//
// GetDescription(err) -> "field foo has an invalid format, expected integer"
func NewFieldFormat(field, reqFormat string) error {
	return fmt.Errorf(format, FieldFormat, fmt.Sprintf(FieldFormatDescription, field, reqFormat))
}

// NewFieldRange Create and wrap an FieldRange exception
//
// Usage
//
// err := NewFieldRange("foo", "1", "256") -> "field has an invalid range#field foo is out of range [1, 256)"
//
// GetDescription(err) -> "field foo is out of range [1, 256)"
func NewFieldRange(field, x, y string) error {
	return fmt.Errorf(format, FieldRange, fmt.Sprintf(FieldRangeDescription, field, x, y))
}

// WrapCustomError Wrap a custom domain error with a description
//
// Usage
//
// err := WrapCustomError(ErrCustom, "error description example with field %s", "foo") ->
// "custom error#error description example with field foo"
//
// GetDescription(err) ->
// error description example with field foo
func WrapCustomError(err error, description string, field interface{}) error {
	return fmt.Errorf(format, err, fmt.Sprintf(description, field))
}

// NewCustomError Generate and wrap a domain error with its own description
//
// Deprecated
func NewCustomError(err error, fields ...string) error {
	switch err {
	case AlreadyExists:
		errParam := verifyParamsSize(1, fields)
		if errParam != nil {
			return errParam
		}

		return fmt.Errorf(format, err, fmt.Sprintf(AlreadyExistsDescription, fields[0]))
	case NotFound:
		errParam := verifyParamsSize(1, fields)
		if errParam != nil {
			return errParam
		}

		return fmt.Errorf(format, err, fmt.Sprintf(NotFoundDescription, fields[0]))
	case RequiredField:
		errParam := verifyParamsSize(1, fields)
		if errParam != nil {
			return errParam
		}

		return fmt.Errorf(format, err, fmt.Sprintf(RequiredFieldDescription, fields[0]))
	case FieldFormat:
		errParam := verifyParamsSize(2, fields)
		if errParam != nil {
			return errParam
		}

		return fmt.Errorf(format, err, fmt.Sprintf(FieldFormatDescription, fields[0], fields[1]))
	case FieldRange:
		errParam := verifyParamsSize(3, fields)
		if errParam != nil {
			return errParam
		}

		return fmt.Errorf(format, err, fmt.Sprintf(FieldRangeDescription, fields[0], fields[1], fields[2]))
	default:
		return err
	}
}

// GetDescription Get a detailed description from a custom error
func GetDescription(err error) string {
	s := strings.Split(err.Error(), "#")

	if len(s)>1 {
		return strings.Join(s[1:], "")
	}

	return err.Error()
}
