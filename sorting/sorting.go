package sorting

import (
	"math/rand"
)

// GenIntSlice - generate int slice
func GenIntSlice(sz, nonUniq int) []int {
	uniqs := make(map[int]int)

	s := make([]int, sz)
	for i := 0; i < sz; i++ {
		var (
			v  int
			ex bool
		)

		for !ex {
			v = rand.Intn(sz)
			u := uniqs[v]
			uniqs[v] += 1
			ex = nonUniq >= u
		}

		s[i] = v
	}

	return s
}

func swapIntItems(s []int, i, j int) {
	var t = s[j]
	s[j] = s[i]
	s[i] = t
}
