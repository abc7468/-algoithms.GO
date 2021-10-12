package q2110

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var r *bufio.Reader

func Start() {
	r = bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	sort.Ints(arr)
	left := 0
	var ans int
	right := arr[n-1] - arr[0]
	for left <= right {
		mid := (left + right) / 2
		cnt := 1
		start := arr[0]
		for i := 1; i < n; i++ {
			dis := arr[i] - start
			if dis >= mid {
				cnt++
				start = arr[i]
			}
		}
		if cnt >= m {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	fmt.Println(ans)
}
