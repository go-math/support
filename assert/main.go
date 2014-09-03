// Package assert provides a number of functions for making assertions in
// tests.
package assert

import (
	"reflect"
	"testing"
)

// Equal asserts that two objects are equal
func Equal(actual, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Fatalf("got %v (%T) instead of %v (%T)", actual, actual, expected, expected)
	}
}

// DeepEqual asserts that two objects with nontrivial types, such as structs
// and slices, are equal.
func DeepEqual(actual, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got %v (%T) instead of %v (%T)", actual, actual, expected, expected)
	}
}

// Success asserts that the error is nil.
func Success(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("got an error '%v'", err)
	}
}

// Success asserts that the error is not nil.
func Failure(err error, t *testing.T) {
	if err == nil {
		t.Fatalf("expected an error")
	}
}
