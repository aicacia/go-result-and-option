package types

import "testing"

func TestSome(t *testing.T) {
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

func TestNone(t *testing.T) {
	none := None[int]()
	if none.IsSome() {
		t.Error("option is some")
	}
}

func TestTake(t *testing.T) {
	some := Some[int](10)
	var y int
	if x, ok := some.Take(); ok {
		y = x
	}
	if y != 10 {
		t.Error("option value is not 10")
	}
}
