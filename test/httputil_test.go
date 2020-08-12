package test

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/common-go/httputil"
	"testing"
)

func TestGetStatusFromError(t *testing.T) {
	code := httputil.GetStatusFromError(exception.NewCustomError(exception.AlreadyExists, "foo"))
	t.Log(code)

	if code != 409 {
		t.Fatal("invalid http code, expected 409")
	}

	code = httputil.GetStatusFromError(exception.NewCustomError(exception.NotFound, "bar"))
	t.Log(code)

	if code != 404 {
		t.Fatal("invalid http code, expected 404")
	}

	code = httputil.GetStatusFromError(exception.NewCustomError(errors.New("example error")))
	t.Log(code)

	if code != 500 {
		t.Fatal("invalid http code, expected 500")
	}
}
