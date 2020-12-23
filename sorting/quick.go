package sorting

// QuickSortInt - quick sort slice
func QuickSortInt(s []int) {
	if len(s) <= 1 {
		return
	}

	pi := len(s) - 1
	left := 0
	right := pi

	for {
		for ; left < len(s)-1 && s[left] <= s[pi]; left++ {
		}

		for ; right >= 0 && s[right] > s[pi]; right-- {
		}

		if left >= right {
			break
		}

		swapIntItems(s, left, right)
	}

	swapIntItems(s, pi, left)
	QuickSortInt(s[:left])
	QuickSortInt(s[left:])

	return
}
