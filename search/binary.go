package search

// BinarySearchInt - binary search in slice
func BinarySearchInt(s []int, x int) int {
	left := 0
	right := len(s)

	for left != right {
		var median = (left + right) / 2
		switch {
		case x == s[median]:
			return median
		case x < s[median]:
			right = median
		default:
			left = median + 1
		}

	}

	return -1
}