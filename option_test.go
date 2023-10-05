package types

import "testing"

func TestOptionSome(t *testing.T) {
	some := Some[int](10)
	var x int
	if y, ok := some.Some(); ok {
		x = y
	} else {
		t.Error("option is none")
	}
	if x != 10 {
		t.Error("option value is not 10")
	}
}

func TestOptionNone(t *testing.T) {
	none := None[int]()
	if none.IsSome() {
		t.Error("option is some")
	}
}

func TestOptionTake(t *testing.T) {
	some := Some[int](10)
	var y int
	if x, ok := some.Take(); ok {
		y = x
	}
	if y != 10 {
		t.Error("option value is not 10")
	}
	if some.IsSome() {
		t.Error("taken option is still some")
	}
}

func TestOptionUnwrap(t *testing.T) {
	some := Some[int](10)
	y := some.Unwrap()
	if y != 10 {
		t.Error("option value is not 10")
	}
}

func TestOptionUnwrapPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("should panic")
		}
	}()
	none := None[int]()
	none.Unwrap()
}
