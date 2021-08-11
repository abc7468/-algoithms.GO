package q16234

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type country struct {
	x          int
	y          int
	population int
}

type pos struct {
	x int
	y int
}

var countryStatus [][]*country
var visited [][]bool
var openCountryGroup [][]*country

var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, -1, 0, 1}

var N, L, R int
var queue []pos

var groupNum int

func bfsForOpenCountry(groupNum int) {
	for len(queue) != 0 {
		nowX := queue[0].x
		nowY := queue[0].y
		queue = queue[1:]

		for k := 0; k < 4; k++ {
			nextX := nowX + dx[k]
			nextY := nowY + dy[k]
			if posRange(nextX, nextY) && visited[nextX][nextY] == false {
				if isNeedOpen(countryStatus[nowX][nowY].population, countryStatus[nextX][nextY].population) {
					visited[nextX][nextY] = true
					queue = append(queue, pos{nextX, nextY})
					openCountryGroup[groupNum] = append(openCountryGroup[groupNum], countryStatus[nextX][nextY])
				}
			}
		}
	}
}

func isNeedOpen(first, second int) bool {
	tmp := int(math.Abs(float64(first) - float64(second)))
	if tmp >= L && tmp <= R {
		return true
	}
	return false
}

func populationMove() {
	for i := 0; i < len(openCountryGroup); i++ {
		sum := 0
		for _, c := range openCountryGroup[i] {
			sum += c.population
		}
		val := int(sum / len(openCountryGroup[i]))
		for _, c := range openCountryGroup[i] {
			c.population = val
		}
	}
}

func initAlgo() {
	visited = make([][]bool, N)
	groupNum = 0
	for i := range countryStatus {
		visited[i] = make([]bool, N)
	}
	for i := range countryStatus {
		for j := range countryStatus {
			visited[i][j] = false
		}
	}
	openCountryGroup = make([][]*country, 0)
}

func posRange(x, y int) bool {
	if x >= 0 && x < N && y >= 0 && y < N {
		return true
	}
	return false
}

func stringToInt(a string) int {
	val, _ := strconv.Atoi(a)
	return val
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fmt.Scan(&N, &L, &R)
	countryStatus = make([][]*country, N)
	for i := range countryStatus {
		countryStatus[i] = make([]*country, N)
	}
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			scanner.Scan()
			countryStatus[x][y] = &country{x, y, stringToInt(scanner.Text())}
		}
	}
	day := 0
	for {
		initAlgo()
		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				if visited[x][y] == false {
					queue = append(queue, pos{x, y})
					visited[x][y] = true
					openCountryGroup = append(openCountryGroup, make([]*country, 0))
					openCountryGroup[groupNum] = append(openCountryGroup[groupNum], countryStatus[x][y])
					bfsForOpenCountry(groupNum)
					groupNum++
				}
			}
		}
		populationMove()
		if len(openCountryGroup) == N*N {
			break
		}
		day++
	}
	fmt.Println(day)
}
