package types

import (
	"fmt"
	"testing"
)

func TestOk(t *testing.T) {
	result := Ok[int](10)
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

func TestOkFail(t *testing.T) {
	result := Err[int](fmt.Errorf("error"))
	if _, ok := result.Ok(); ok {
		t.Error("result is not err")
	}
}

func TestErr(t *testing.T) {
	result := Err[int](fmt.Errorf("error"))
	if err, ok := result.Err(); ok {
		if err.Error() != "error" {
			t.Error("result error is not error")
		}
	}
}

func TestErrFail(t *testing.T) {
	result := Ok[int](10)
	if _, ok := result.Err(); ok {
		t.Error("result is not ok")
	}
}

func TestUnwrapOn(t *testing.T) {
	result := Ok[int](10)
	x := result.Unwrap()
	if x != 10 {
		t.Error("result value is not 10")
	}
}

func TestUnwrapOnErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The result did not panic")
		}
	}()

	result := Err[int](fmt.Errorf("error"))
	result.Unwrap()
}
