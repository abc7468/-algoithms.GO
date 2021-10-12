package q2056

import (
	"bufio"
	"fmt"
	"os"
)

type queue []int

var dp []int
var parents []int
var child [][]int
var mmap map[int][]int

func (q *queue) push(val int) {
	*q = append(*q, val)
}

func (q *queue) pop() int {
	if len(*q) == 0 {
		return -1
	}
	old := *q
	val := old[0]
	*q = old[1:]
	return val
}

func (q queue) isEmpty() bool {
	if len(q) == 0 {
		return true
	}
	return false
}

func Start() {
	r := bufio.NewReader(os.Stdin)
	var n int
	var q queue
	fmt.Fscan(r, &n)
	dp = make([]int, n)
	parents = make([]int, n)
	child = make([][]int, n)
	mmap = map[int][]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &dp[i])
		fmt.Fscan(r, &parents[i])
		if parents[i] == 0 {
			q.push(i)
		}
		for j := 0; j < parents[i]; j++ {
			var tmp int
			fmt.Fscan(r, &tmp)
			mmap[i] = append(mmap[i], tmp-1)
			child[tmp-1] = append(child[tmp-1], i)
		}
	}
	for !q.isEmpty() {

		now := q.pop()
		for _, v := range child[now] {
			parents[v]--
			if parents[v] == 0 {
				q.push(v)
			}
		}
		if len(mmap[now]) == 0 {
			continue
		}
		max := 0
		for _, v := range mmap[now] {
			if max < dp[v] {
				max = dp[v]
			}
		}
		dp[now] += max

	}
	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	fmt.Println(max)
}
