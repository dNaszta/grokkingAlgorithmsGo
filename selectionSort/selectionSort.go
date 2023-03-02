package selectionSort

func Remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func FindSmallestIndex(array []int) int {
	smallestIndex := 0
	smallest := array[smallestIndex]

	for i, element := range array {
		if element < smallest {
			smallestIndex = i
			smallest = element
		}
	}
	return smallestIndex
}

func SelectionSort(array []int) []int {
	result := make([]int, len(array))
	arrayLength := len(array)
	for i := 0; i < arrayLength; i++ {
		smallest := FindSmallestIndex(array)
		result[i] = array[smallest]
		array = Remove(array, smallest)
	}
	return result
}
