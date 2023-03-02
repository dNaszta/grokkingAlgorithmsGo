package quickSort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("QuickSort it cannot 'sort' empty array", func(t *testing.T) {
		var a []int
		s := QuickSort(a)

		if !reflect.DeepEqual(a, s) {
			t.Errorf("It cannot be sorted: %v", a)
		}

		a = []int{10}
		s = QuickSort(a)
		if !reflect.DeepEqual(a, s) {
			t.Errorf("It cannot be sorted: %v", a)
		}
	})

	t.Run("QuickSort it can 'sort' short array", func(t *testing.T) {
		a := []int{5, 3, 2, 4, 1}
		wanted := []int{1, 2, 3, 4, 5}

		s := QuickSort(a)
		fmt.Println(s, wanted)
	})
}
