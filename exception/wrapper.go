package exception

import (
	"fmt"
	"strings"
)

// NewCustomError Generate and wrap a domain error with its own description
func NewCustomError(err error, fields ...string) error {
	format := "%w#%s"
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
