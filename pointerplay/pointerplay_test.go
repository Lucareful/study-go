package pointerplay_test

import (
	"pointerplay"
	"testing"
)

func TestDouble2(t *testing.T) {
	t.Parallel()
	var x int = 12
	want := 24
	pointerplay.Double(&x)
	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()
	x := pointerplay.MyInt(12)
	want := pointerplay.MyInt(24)
	p := &x
	p.Double()
	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}
}
