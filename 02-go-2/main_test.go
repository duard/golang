package main

import "testing"

func TestAdd(t *testing.T) {

	x, y := 2, 3
	want := 5

	got := Add(x, y)

	if got != want {

		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSub(t *testing.T) {

	x, y := 5, 3
	want := 2

	got := Sub(x, y)

	if got != want {

		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDiv(t *testing.T) {

	x, y := 7., 2.
	want := 3.5

	got := Div(x, y)

	if got != want {

		t.Errorf("got %f, want %f", got, want)
	}
}

func TestMul(t *testing.T) {

	x, y := 6, 5
	want := 30

	got := Mul(x, y)

	if got != want {

		t.Errorf("got %d, want %d", got, want)
	}
}
