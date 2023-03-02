package binarySearch

import (
	"math/rand"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	t.Run("Basic example", func(t *testing.T) {
		l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
		rand.Seed(time.Now().Unix())
		wantedPosition := rand.Int() % len(l)
		wantedElement := l[wantedPosition]

		got, _ := BinarySearch(l, wantedElement)
		if wantedPosition != got {
			t.Errorf("List %v, Wanted element: %d, Wanted position: %d, Got: %d", l, wantedElement,
				wantedPosition, got)
		}
	})
}
