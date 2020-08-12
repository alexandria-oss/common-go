package test

import (
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestNewCustomError(t *testing.T) {
	custom := exception.NewCustomError(exception.AlreadyExists, "foo")
	t.Log(custom)

	desc := exception.GetDescription(custom)
	t.Log(desc)

	if desc != "foo already exists" {
		t.Fatal("invalid error description, expected 'foo already exists'")
	}

	custom = exception.NewCustomError(exception.NotFound, "bar")
	t.Log(custom)

	desc = exception.GetDescription(custom)
	t.Log(desc)

	if desc != "bar not found" {
		t.Fatal("invalid error description, expected 'bar not found'")
	}

	custom = exception.NewCustomError(exception.FieldRange, "foo", "1", "255")
	t.Log(custom)

	desc = exception.GetDescription(custom)
	t.Log(desc)

	if desc != "field foo is out of range [1, 255)" {
		t.Fatal("invalid error description, expected 'field foo is out of range [1, 255)'")
	}
}
