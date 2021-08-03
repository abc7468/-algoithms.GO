package q12851

import (
	"fmt"
)

var dp = []int{}
var step = []int{}
var visited = []bool{}
var N, M int
var max int

func bfs(node int) {
	if node-1 >= 0 {

	}
	if node+1 < max {

	}
	if node%2 == 0 {
		// heap.Pop
	}
}

func Start() {
	fmt.Scan(&N, &M)
	fmt.Print(N, M)
	if N > M {
		max = N
	} else {
		max = M
	}
	max = max*2 + 1
	dp = make([]int, max)
	step = make([]int, max)
	visited = make([]bool, max)
}
