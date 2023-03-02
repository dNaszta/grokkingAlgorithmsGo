package quickSort

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := a[0]
	var less []int
	var more []int
	for i := 1; i < len(a); i++ {
		if pivot > a[i] {
			less = append(less, a[i])
		} else {
			more = append(more, a[i])
		}
	}

	less = QuickSort(less)
	more = QuickSort(more)

	result := append(less, pivot)
	result = append(result, more...)
	return result
}
