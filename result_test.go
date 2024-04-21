package types

import (
	"fmt"
	"testing"
)

var errTestError = fmt.Errorf("error")

func TestResultOk(t *testing.T) {
	result := Ok(10)
	var x int
	if y, ok := result.Ok(); ok {
		x = y
	} else {
		t.Error("result is not ok")
	}
	if x != 10 {
		t.Error("result value is not 10")
	}
}

func TestResultOkFail(t *testing.T) {
	result := Err[int](errTestError)
	if _, ok := result.Ok(); ok {
		t.Error("result is not err")
	}
}

func TestResultErr(t *testing.T) {
	result := Err[int](errTestError)
	if err, ok := result.Err(); ok {
		if err.Error() != "error" {
			t.Error("result error is not error")
		}
	}
}

func TestResultErrFail(t *testing.T) {
	result := Ok(10)
	if _, ok := result.Err(); ok {
		t.Error("result is not ok")
	}
}

func TestResultUnwrap(t *testing.T) {
	result := Ok(10)
	x := result.Unwrap()
	if x != 10 {
		t.Error("result value is not 10")
	}
}

func TestResultUnwrapPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The result did not panic")
		}
	}()

	result := Err[int](errTestError)
	result.Unwrap()
}

func TestResultUnwrapErr(t *testing.T) {
	result := Err[int](errTestError)
	err := result.UnwrapErr()
	if err != errTestError {
		t.Error("result error is not an error")
	}
}

func TestResultUnwrapErrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The result did not panic")
		}
	}()

	result := Ok(10)
	result.UnwrapErr()
}

func TestResultOptionSome(t *testing.T) {
	result := Ok(10)
	option := result.Option()
	if option.IsNone() {
		t.Error("result is some")
	}
}

func TestResultOptionNone(t *testing.T) {
	result := Err[int](errTestError)
	option := result.Option()
	if option.IsSome() {
		t.Error("result is not none")
	}
}
