package selectionSort

import "testing"

func TestFindSmallestIndex(t *testing.T) {
	t.Run("Find smallest item from an unordered list", func(t *testing.T) {
		wanted := 6
		l := []int{7, 8, 9, 3, 4, 2, 1, 5}
		got := FindSmallestIndex(l)
		if got != wanted {
			t.Errorf("Found incorrect index: %d - wanted: %d", got, wanted)
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("Remove indexed element from an unordered list", func(t *testing.T) {
		l := []int{7, 8, 9, 3, 4, 2, 1, 5}
		unwantedElementValue := 9
		unwantedElementIndex := 2

		l = Remove(l, unwantedElementIndex)

		for _, e := range l {
			if e == unwantedElementValue {
				t.Errorf("Remove is incorrect, %d found in %v list", unwantedElementValue, l)
			}
		}
	})
}

func TestSelectionSort(t *testing.T) {
	t.Run("Sorting a list with the hard way", func(t *testing.T) {
		l := []int{7, 8, 9, 3, 4, 2, 1, 5}
		wanted := []int{1, 2, 3, 4, 5, 7, 8, 9}
		got := SelectionSort(l)

		for i := 0; i < len(wanted); i++ {
			if wanted[i] != got[i] {
				t.Errorf("Lists mismatch, on element %d, got: %d, wanted: %d", i, got[i], wanted[i])
			}
		}
	})
}
