package httputil

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexandria-oss/common-go/exception"
	"net/http"
)

// GetStatusFromError Get an HTTP status from a custom error
func GetStatusFromError(err error) int {
	switch {
	case errors.Is(err, exception.AlreadyExists):
		return http.StatusConflict
	case errors.Is(err, exception.NotFound):
		return http.StatusNotFound
	case errors.Is(err, exception.RequiredField) || errors.Is(err, exception.FieldRange) ||
			errors.Is(err, exception.FieldFormat) || errors.Is(err, exception.Invalid):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// RespondErrorJSON Generate and send a pre-marshalled error with JSON format if available
func RespondErrorJSON(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		code := GetStatusFromError(err)
		desc := exception.GetDescription(err)

		w.WriteHeader(code)
		errJSON := json.NewEncoder(w).Encode(&Response{
			Message: desc,
			Code:    code,
		})

		if errJSON != nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			_, _ = fmt.Fprintf(w, `%v`, &Response{
				Message: desc,
				Code:    code,
			})
		}
	}
}
