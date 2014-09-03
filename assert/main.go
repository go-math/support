// Package assert provides a number of functions for making assertions in
// tests.
package assert

import (
	"reflect"
	"runtime"
	"testing"
)

// Equal asserts that two objects are equal
func Equal(actual, expected interface{}, t *testing.T) {
	if actual != expected {
		raise(t, "got %v (%T) instead of %v (%T)", actual, actual, expected, expected)
	}
}

// DeepEqual asserts that two objects with nontrivial types, such as structs
// and slices, are equal.
func DeepEqual(actual, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		raise(t, "got %v (%T) instead of %v (%T)", actual, actual, expected, expected)
	}
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
