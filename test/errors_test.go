package test

import (
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestNewCustomError(t *testing.T) {
	custom := exception.NewCustomError(exception.AlreadyExists, "foo")
	t.Log(custom)

	t.Log(exception.GetDescription(custom))
}
