package binarySearch

func BinarySearch(l []int, e int) (int, bool) {
	low := 0
	high := len(l) - 1

	for low <= high {
		mid := int((low + high) / 2)
		guess := l[mid]
		if guess == e {
			return mid, true
		}
		if guess > e {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}
