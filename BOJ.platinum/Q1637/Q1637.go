package q1637

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var arr [][]int
var n int

func strToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Start() {
	r := bufio.NewReader(os.Stdin)
	// w := bufio.NewWriter(os.Stdout)
	fmt.Fscan(r, &n)
	arr = make([][]int, n)
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)
		arr[i] = append(arr[i], a)
		arr[i] = append(arr[i], b)
		arr[i] = append(arr[i], c)
	}
	left := 0
	right := math.MaxInt32
	var cnt int64
	for left+1 <= right {
		mid := (left + right) / 2
		cnt = getCnt(mid)
		if cnt%2 == 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// if getCnt(left)%2 == 0 {
	// 	fmt.Println("NOTHING")
	// 	return
	// }
	fmt.Println(left, getCnt(left))
	fmt.Println(getCnt(left - 1))
}
func getCnt(val int) (tmp int64) {
	tmp = 0
	for i := 0; i < n; i++ {
		start := arr[i][0]
		if start > val {
			continue
		}
		end := min(arr[i][1], val)
		tmp += int64((end-start)/arr[i][2]) + 1
	}
	return tmp
}
