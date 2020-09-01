package grpcutil

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetStatusFromError Get an gRPC status from a domain or custom error
func GetStatusFromError(err error) codes.Code {
	switch {
	case errors.Is(err, exception.AlreadyExists):
		return codes.AlreadyExists
	case errors.Is(err, exception.NotFound):
		return codes.NotFound
	case errors.Is(err, exception.RequiredField) || errors.Is(err, exception.FieldFormat) ||
			errors.Is(err, exception.Invalid):
		return codes.InvalidArgument
	case errors.Is(err, exception.FieldRange):
		return codes.OutOfRange
	default:
		return codes.Internal
	}
}

// RespondError Generate a domain error with gRPC format if available
func RespondError(err error) error {
	if err != nil {
		return status.Error(GetStatusFromError(err), exception.GetDescription(err))
	}

	return nil
}
