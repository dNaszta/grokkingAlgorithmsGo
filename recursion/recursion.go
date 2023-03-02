package recursion

import "fmt"

func Countdown(n int) string {
	result := fmt.Sprintf("%d", n)

	if n > 0 {
		result += ".." + Countdown(n-1)
	}

	return result
}

func Fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Fact(n-1)
	}
}

func Sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return a[0] + Sum(a[1:])
}
