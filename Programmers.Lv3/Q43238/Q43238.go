package q43238

import (
	"fmt"
	"sort"
)

func binarySearch(n int, times []int) int64 {
	sort.Ints(times)
	left := 1
	right := n * times[len(times)-1]
	answer := 0
	for left < right {
		middle := ((left + right) / 2)
		tmp := 0
		for _, val := range times {
			tmp += middle / val
		}
		if tmp < n {
			left = middle + 1
		} else if tmp >= n {
			answer = middle

			right = middle - 1
		} else {
			break
		}

	}
	return int64(answer)
}

func solution(n int, times []int) int64 {
	return binarySearch(n, times)
}

func Start() {
	times := []int{7, 10}
	fmt.Println(solution(6, times))
}
