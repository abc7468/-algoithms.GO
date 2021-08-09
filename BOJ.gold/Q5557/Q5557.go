package q5557

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	maxVal int = 21
)

var N int
var inputVal []int
var dp [][]int64
var wantAnswer int

func Start() {
	fmt.Scan(&N)
	dp = make([][]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int64, maxVal)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < N; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		inputVal = append(inputVal, val)
	}
	wantAnswer = inputVal[N-1]
	inputVal = inputVal[:N-1]
	dp[0][inputVal[0]] = 1
	for i := 1; i < N-1; i++ {
		for j := 0; j < maxVal; j++ {
			if dp[i-1][j] != 0 {
				tmp := j + inputVal[i]
				if tmp <= 20 && tmp >= 0 {
					dp[i][tmp] += dp[i-1][j]
				}
				tmp = j - inputVal[i]
				if tmp <= 20 && tmp >= 0 {
					dp[i][tmp] += dp[i-1][j]
				}
			}
		}
	}
	fmt.Println(dp[N-2][wantAnswer])
}
