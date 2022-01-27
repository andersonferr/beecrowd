package main

import (
	"fmt"
	"sort"
)

func main() {
	for c := 1; true; c++ {
		var n, q int

		fmt.Scanf("%d %d\n", &n, &q)
		if n == 0 && q == 0 {
			return
		}

		fmt.Printf("CASE# %d:\n", c)

		marmores := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scanf("%d\n", &marmores[i])
		}

		sort.Ints(marmores)

		var query int
		for i := 0; i < q; i++ {
			fmt.Scanf("%d\n", &query)
			idx := find(query, marmores)
			if idx >= 0 {
				fmt.Printf("%d found at %d\n", query, idx+1)
			} else {
				fmt.Printf("%d not found\n", query)
			}
		}
	}
}

func find(n int, a []int) int {
	min := 0
	max := len(a) - 1
	for min <= max {
		mid := (max + min) / 2

		switch {
		case a[mid] > n:
			max = mid - 1

		case a[mid] < n:
			min = mid + 1

		case a[mid] == n:
			if max == min {
				return min
			} else {
				max = mid
			}
		}
	}

	return -1
}
