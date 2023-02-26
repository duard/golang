package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("fixed 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Add(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
