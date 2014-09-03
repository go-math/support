// Package assert provides a number of functions for making assertions in
// tests.
package assert

import (
	"reflect"
	"runtime"
	"testing"
)

// Equal asserts that two objects are equal.
func Equal(actual, expected interface{}, t *testing.T) {
	kind := reflect.TypeOf(actual).Kind()

	if kind != reflect.TypeOf(expected).Kind() {
		goto error
	}

	if kind == reflect.Slice || kind == reflect.Struct {
		if !reflect.DeepEqual(actual, expected) {
			goto error
		}
	} else {
		if actual != expected {
			goto error
		}
	}

	return

error:
	raise(t, "got %v (%T) instead of %v (%T)", actual, actual, expected, expected)
}

// Success asserts that the error is nil.
func Success(err error, t *testing.T) {
	if err != nil {
		raise(t, "got an error '%v'", err)
	}
}

// Success asserts that the error is not nil.
func Failure(err error, t *testing.T) {
	if err == nil {
		raise(t, "expected an error")
	}
}

func raise(t *testing.T, format string, args ...interface{}) {
	if _, file, line, ok := runtime.Caller(2); ok {
		t.Errorf("%s:%d", file, line)
	}
	t.Fatalf(format, args...)
}
