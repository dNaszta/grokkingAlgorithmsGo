package recursion

import (
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("Countdown from 3", func(t *testing.T) {
		wanted := "3..2..1..0"
		got := Countdown(3)
		if got != wanted {
			t.Errorf("Result missmatch, wanted: %s, got: %s", wanted, got)
		}
	})
}

func TestFact(t *testing.T) {
	t.Run("5!", func(t *testing.T) {
		wanted := 120
		got := Fact(5)

		if got != wanted {
			t.Errorf("Result missmatch, wanted: %d, got: %d", wanted, got)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("Sum with empty array response 0", func(t *testing.T) {
		a := []int{}
		got := Sum(a)
		wanted := 0
		if got != wanted {
			t.Errorf("Sum with empty array is not 0")
		}
	})

	t.Run("Sum with response correctly with 1 element", func(t *testing.T) {
		a := []int{10}
		got := Sum(a)
		wanted := 10
		if got != wanted {
			t.Errorf("Sum error, wanted: %d, got: %d", wanted, got)
		}
	})

	t.Run("Sum with response correctly with [1,2,3] elements", func(t *testing.T) {
		a := []int{1, 2, 3}
		got := Sum(a)
		wanted := 6
		if got != wanted {
			t.Errorf("Sum error, wanted: %d, got: %d", wanted, got)
		}
	})
}
