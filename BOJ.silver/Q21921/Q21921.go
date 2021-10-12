package q21921

import (
	"bufio"
	"fmt"
	"os"
)

func Start() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	arr := make([]int, n)
	var max int64
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
		if i < m {
			max += int64(arr[i])
		}
	}
	tmp := max
	cnt := 1
	for i := 1; i < n-m+1; i++ {
		tmp -= int64(arr[i-1])
		tmp += int64(arr[i+m-1])
		if tmp > max {
			max = tmp
			cnt = 1
		} else if tmp == max {
			cnt++
		}
	}
	fmt.Println(max)
	fmt.Println(cnt)

}
