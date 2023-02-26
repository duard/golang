package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitcode := m.Run()
	os.Exit(exitcode)
}

func TestRunMain(t *testing.T) {
	main()
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
func TestAdd(t *testing.T) {
	t.Run("fixed 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Add(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Another size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Add(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
