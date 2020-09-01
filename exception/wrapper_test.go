package exception

import (
	"errors"
	"testing"
)

func TestNewAlreadyExists(t *testing.T) {
	if err := NewAlreadyExists("foo"); err != nil {
		t.Log(GetDescription(err))

		msg := "foo already exists"
		if GetDescription(err) != msg {
			t.Fatalf("invalid error description, expected '%s'", msg)
		}
	}
}

func TestNewNotFound(t *testing.T) {
	err := NewNotFound("foo")
	t.Log(GetDescription(err))

	msg := "foo not found"
	if GetDescription(err) != msg {
		t.Fatalf("invalid error description, expected '%s'", msg)
	}
}

func TestNewRequiredField(t *testing.T) {
	err := NewRequiredField("foo")
	t.Log(GetDescription(err))

	msg := "missing required field: foo"
	if GetDescription(err) != msg {
		t.Fatalf("invalid error description, expected '%s'", msg)
	}
}

func TestNewFieldFormat(t *testing.T) {
	err := NewFieldFormat("foo", "integer 32 bits")
	t.Log(GetDescription(err))

	msg := "field foo has an invalid format, expected integer 32 bits"
	if GetDescription(err) != msg {
		t.Fatalf("invalid error description, expected '%s'", msg)
	}
}

func TestNewFieldRange(t *testing.T) {
	err := NewFieldRange("foo", "1", "256")
	t.Log(GetDescription(err))

	msg := "field foo is out of range [1, 256)"
	if GetDescription(err) != msg {
		t.Fatalf("invalid error description, expected '%s'", msg)
	}
}

func TestNewNetworkCall(t *testing.T) {
	err := NewNetworkCall("cassandra", "192.168.1.64:9042")
	t.Log(GetDescription(err))

	msg := "remote call to cassandra with address 192.168.1.64:9042 has failed"
	if GetDescription(err) != msg {
		t.Fatalf("invalid error description, expected '%s'", msg)
	}
}

func TestWrapCustomError(t *testing.T) {
	domainErr := errors.New("custom error")
	err := WrapCustomError(domainErr, "error description with field %v", "foo")
	t.Log(GetDescription(err))

	msg := "error description with field foo"
	if GetDescription(err) != msg {
		t.Fatalf("invalid error description, exptected '%s'", msg)
	}
}
