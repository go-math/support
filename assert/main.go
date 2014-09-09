// Package assert provides a number of functions for making assertions in
// tests.
package assert

import (
	"math"
	"reflect"
	"runtime"
	"testing"
)

const (
	epsilon = 1e-8
)

// Equal asserts that two objects are equal.
func Equal(actual, expected interface{}, t *testing.T) {
	kind := reflect.TypeOf(actual).Kind()

	if kind != reflect.TypeOf(expected).Kind() {
		goto error
	}

	if kind == reflect.Slice || kind == reflect.Struct || kind == reflect.Ptr {
		if !reflect.DeepEqual(actual, expected) {
			goto error
		}
	} else if actual != expected {
		goto error
	}

	return

error:
	raise(t, "got %v (%T) instead of %v (%T)", actual, actual, expected, expected)
}

// AlmostEqual asserts that the absolute difference between any corresponding
// pair of elements of two float64 slices is not more than a predefined small
// constant, which is 1e-8.
func AlmostEqual(actual, expected []float64, t *testing.T) {
	if len(actual) != len(expected) {
		goto error
	}

	for i := range actual {
		if math.Abs(actual[i]-expected[i]) > epsilon {
			goto error
		}
	}

	return

error:
	raise(t, "got %v instead of %v", actual, expected)
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
