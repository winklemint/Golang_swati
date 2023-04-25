package maths

import (
	"testing"
)

func AddTest(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
func SubTest(t *testing.T) {
	got := Sub(4, 7)
	want := -3
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func MultiTest(t *testing.T) {
	got := Multiply(5, 5)
	want := 25
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func DivTest(t *testing.T) {
	got := Div(99.11, 11.22)
	want := 9.36
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
