package q1516

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func StrToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

type Obj struct {
	in    []int
	time  int
	index int
}

type queue []int

func (q *queue) push(objNum int) {
	*q = append(*q, objNum)
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
func (q *queue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	var q queue
	n := StrToInt(scanner.Text())

	var objects []Obj = make([]Obj, n+1)
	var need []int = make([]int, n+1)
	var dp []int = make([]int, n+1)
	var chain [][]int = make([][]int, n+1)

	for i := 1; i <= n; i++ {
		o := Obj{}
		cnt := 0

		for {
			scanner.Scan()
			input := StrToInt(scanner.Text())
			if cnt == 0 {
				o.index = i
				o.time = input
				cnt++
				continue
			}
			if input == -1 {
				objects[i] = o
				break
			}
			o.in = append(o.in, input)
			chain[input] = append(chain[input], i)
			need[i]++
		}
	}
	for i := 1; i <= n; i++ {
		dp[i] = objects[i].time
		if len(objects[i].in) == 0 {
			q.push(i)
		}
	}

	for !q.isEmpty() {
		idx := q.pop()
		for _, v := range chain[idx] {
			need[v]--
			if need[v] == 0 {
				q.push(v)
			}
		}
		if len(objects[idx].in) == 0 {
			dp[idx] = objects[idx].time
			continue
		}
		max := 0
		for _, v := range objects[idx].in {
			if max < dp[v]+dp[idx] {
				max = dp[v] + dp[idx]
			}
		}
		dp[idx] = max
	}
	for i := 1; i <= n; i++ {
		fmt.Println(dp[i])
	}
}
