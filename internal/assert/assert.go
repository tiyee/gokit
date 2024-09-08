// Use of this source code is governed by MIT license

// Package assert is for internal use.
package assert

import (
	"cmp"
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"testing"
)

type IAssert interface {
	CaseName() string
	T() *testing.T
}

// Assert is a simple implementation of assertion, only for internal usage
type Assert struct {
	t    *testing.T
	Name string
}

func (a *Assert) CaseName() string {
	return a.Name
}

func (a *Assert) T() *testing.T {
	return a.t
}

// NewAssert return instance of Assert
func NewAssert(t *testing.T, caseName string) *Assert {
	return &Assert{t: t, Name: caseName}
}

// Equal check if expected is equal with actual
func Equal[T comparable](a IAssert, expected, actual T) {
	if expected != actual {
		makeTestFailed(a.T(), a.CaseName(), expected, actual)
	}
}

// ShouldBeFalse  check if expected is true
func ShouldBeFalse(a IAssert, actual bool) {
	if actual != false {
		makeTestFailed(a.T(), a.CaseName(), false, actual)
	}
}

// NotEqual check if expected is not equal with actual
func NotEqual[T comparable](a IAssert, expected, actual T) {
	if expected == actual {
		makeTestFailed(a.T(), a.CaseName(), expected, actual)
	}
}

// Greater check if expected is greate than actual
func GT[T cmp.Ordered](a IAssert, expected, actual T) {
	if expected <= actual {
		expectedInfo := fmt.Sprintf("> %v", expected)
		makeTestFailed(a.T(), a.CaseName(), expectedInfo, actual)
	}
}

// GreaterOrEqual check if expected is greate than or equal with actual
func GTE[T cmp.Ordered](a IAssert, expected, actual T) {
	if expected < actual {
		expectedInfo := fmt.Sprintf(">= %v", expected)
		makeTestFailed(a.T(), a.CaseName(), expectedInfo, actual)
	}
}

// LT check if expected is less than actual
func LT[T cmp.Ordered](a IAssert, expected, actual T) {
	if expected >= actual {
		expectedInfo := fmt.Sprintf("< %v", expected)
		makeTestFailed(a.T(), a.CaseName(), expectedInfo, actual)
	}
}

// LTE check if expected is less than or equal with actual
func LTE[T cmp.Ordered](a IAssert, expected, actual T) {

	if expected > actual {
		expectedInfo := fmt.Sprintf("<= %v", expected)
		makeTestFailed(a.T(), a.CaseName(), expectedInfo, actual)
	}
}

// IsNil check if value is nil
func IsNil(a IAssert, v any) {
	if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		return
	}

	makeTestFailed(a.T(), a.CaseName(), nil, v)
}

// IsNotNil check if value is not nil
func IsNotNil(a IAssert, v any) {
	if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		makeTestFailed(a.T(), a.CaseName(), "not nil", v)
	}
}
func ContainsKey[V any](a IAssert, k string, m map[string]V) {
	if _, exist := m[k]; !exist {
		makeTestFailed(a.T(), a.CaseName(), "not ContainsKey", k)
	}
}
func Contains[T comparable](a IAssert, k T, arr []T) {
	exist := false
	for _, v := range arr {
		if v == k {
			exist = true
			break
		}
	}
	if !exist {
		makeTestFailed(a.T(), a.CaseName(), "not Contains", k)
	}
}

// PanicRunFunc define
type PanicRunFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func runPanicFunc(f PanicRunFunc) (didPanic bool, message any, stack string) {
	didPanic = true
	defer func() {
		message = recover()
		if didPanic {
			stack = string(debug.Stack())
		}
	}()

	// call the target function
	f()
	didPanic = false

	return
}

// Panics asserts that the code inside the specified func panics.
func Panics(a IAssert, fn PanicRunFunc) {
	if hasPanic, panicVal, _ := runPanicFunc(fn); !hasPanic {
		makeTestFailed(a.T(), a.CaseName(), "not Panic", fmt.Sprintf("func '%#v' should panic\n\tPanic value:\t%#v", fn, panicVal))
	}

}

// makeTestFailed make test failed and log error info
func makeTestFailed(t *testing.T, caseName string, expected, actual any) {
	_, file, line, _ := runtime.Caller(2)
	errInfo := fmt.Sprintf("Case %v failed. file: %v, line: %v, expected: %v, actual: %v.", caseName, file, line, expected, actual)
	t.Error(errInfo)
	t.FailNow()
}
