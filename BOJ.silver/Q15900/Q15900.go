package q15900

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var graph [][]int
var visited []int

type queue []int

func (q *queue) push(val int) {
	*q = append(*q, val)
}

func (q *queue) pop() int {
	older := *q
	val := older[0]
	*q = older[1:]
	return val
}
func (q queue) isEmpty() bool {
	if len(q) == 0 {
		return true
	}
	return false
}

func bfs() int64 {
	var q queue
	var cnt int64
	cnt = 0
	q.push(1)
	visited[1] = 0
	for !q.isEmpty() {
		now := q.pop()
		for i := 0; i < len(graph[now]); i++ {
			for _, v := range graph[now] {
				if visited[v] == -1 {
					visited[v] = visited[now] + 1
					if len(graph[v]) == 1 {
						cnt += int64(visited[v])
						continue
					}
					q.push(v)

				}
			}
		}
	}
	return cnt
}

func strToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func Start() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	r.Scan()
	n := strToInt(r.Text())
	graph = make([][]int, n+1)
	visited = make([]int, n+1)
	for i, _ := range visited {
		visited[i] = -1
	}
	for i := 0; i < n-1; i++ {
		var a, b int
		r.Scan()
		a = strToInt(r.Text())
		r.Scan()
		b = strToInt(r.Text())
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	tmp := bfs()
	if tmp%2 == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}
