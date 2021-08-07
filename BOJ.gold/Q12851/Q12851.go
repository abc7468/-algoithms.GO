package q12851

import (
	"fmt"
)

var visited = []bool{}
var N, M int
var max int

type status struct {
	location int
	step     int
}

type queue struct {
	slice []status
}

func (q *queue) pop() status {
	val := q.slice[0]
	q.slice = q.slice[1:]
	return val
}

func (q *queue) push(val status) {
	q.slice = append(q.slice, val)
}

func (q *queue) isEmpty() bool {
	if len(q.slice) == 0 {
		return true
	} else {
		return false
	}
}

func Start() {
	fmt.Scan(&N, &M)
	maxVal := M
	if N > M {
		maxVal = N
	}
	maxVal = maxVal*2 + 1
	visited = make([]bool, maxVal)
	firstStatus := status{N, 0}
	var q queue
	q.slice = append(q.slice, firstStatus)
	var minStep int
	var cnt int
	for !q.isEmpty() {
		nowStatus := q.pop()
		nowLocation := nowStatus.location
		nowStep := nowStatus.step
		visited[nowLocation] = true
		if nowLocation == M && cnt == 0 {
			cnt++
			minStep = nowStep
		} else if nowLocation == M && nowStep == minStep {
			cnt++
		}
		if nowLocation+1 < maxVal && !visited[nowLocation+1] {
			q.push(status{nowLocation + 1, nowStep + 1})
		}
		if nowLocation-1 >= 0 && !visited[nowLocation-1] {
			q.push(status{nowLocation - 1, nowStep + 1})
		}
		if nowLocation*2 < maxVal && !visited[nowLocation*2] {
			q.push(status{nowLocation * 2, nowStep + 1})
		}
	}
	fmt.Println(minStep)
	fmt.Println(cnt)
}
