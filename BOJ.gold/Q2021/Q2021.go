package q2021

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var N, L int
var start, end int

const max int = 987654321

type ns struct{}

var stationBelongLine []map[int]ns
var line [][]int
var visited []bool
var answer = max
var isFind bool = false

// 시작
func getAnswer(hos []int, cnt int) {
	for _, ho := range hos {
		visited[ho] = true
		if findEnd(ho) {
			answer = cnt
			isFind = true
			return
		}
	}
	var tmp = map[int]ns{}
	var nextLines []int
	for _, ho := range hos {
		for key := range getLinesToGo(tmp, ho) {
			tmp = add(tmp, key)
		}
	}
	for key := range tmp {
		nextLines = append(nextLines, key)
	}

	if len(nextLines) == 0 || isFind {
		return
	}
	getAnswer(nextLines, cnt+1)
}

// 해당 라인에 도착지가 있는지
func findEnd(ho int) bool {
	for _, val := range line[ho] {
		if end == val {
			return true
		}
	}
	return false
}

// 다음 턴에 갈 수 있는 모든 라인
func getLinesToGo(lines map[int]ns, ho int) map[int]ns {
	for _, station := range line[ho] {
		for key := range stationBelongLine[station] {
			if !visited[key] {
				lines = add(lines, key)
			}
		}
	}
	return lines
}

func add(station map[int]ns, line int) map[int]ns {
	station[line] = ns{}
	return station
}

func stringToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fmt.Scan(&N, &L)
	line = make([][]int, L+1)
	visited = make([]bool, L+1)
	stationBelongLine = make([]map[int]ns, N+1)
	for i := range stationBelongLine {
		stationBelongLine[i] = map[int]ns{}
	}
	for i := 1; i <= L; i++ {
		for scanner.Scan() {
			inputVal := stringToInt(scanner.Text())
			if inputVal == -1 {
				break
			}
			line[i] = append(line[i], inputVal)
			stationBelongLine[inputVal] = add(stationBelongLine[inputVal], i)
		}
	}
	fmt.Scan(&start, &end)
	var startLines []int
	for key := range stationBelongLine[start] {
		startLines = append(startLines, key)
	}
	getAnswer(startLines, 0)
	if answer == max {
		fmt.Println("-1")
		return
	}
	fmt.Println(answer)
}
