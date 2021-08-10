package q14891

import (
	"fmt"
	"math"
	"strconv"
)

const (
	totalGearCnt      int = 4
	totalValuePerGear int = 8
)

var group [totalGearCnt][totalValuePerGear]int
var done [totalGearCnt]bool

func clockwise(gearNum int) {
	var hasLeft bool
	var hasRight bool
	if gearNum == 0 {
		hasRight = true
	} else if gearNum == 3 {
		hasLeft = true
	} else {
		hasLeft = true
		hasRight = true
	}
	done[gearNum] = true
	if hasLeft {
		if done[gearNum-1] == false && group[gearNum-1][2] != group[gearNum][6] {
			done[gearNum-1] = true
			counterclockwise(gearNum - 1)
		}
	}
	if hasRight {
		if done[gearNum+1] == false && group[gearNum+1][6] != group[gearNum][2] {
			done[gearNum+1] = true
			counterclockwise(gearNum + 1)
		}
	}
	tmp := group[gearNum][totalValuePerGear-1]
	for i := 7; i > 0; i-- {
		group[gearNum][i] = group[gearNum][i-1]
	}
	group[gearNum][0] = tmp

}

func counterclockwise(gearNum int) {
	var hasLeft bool
	var hasRight bool
	if gearNum == 0 {
		hasRight = true
	} else if gearNum == 3 {
		hasLeft = true
	} else {
		hasLeft = true
		hasRight = true
	}

	done[gearNum] = true
	if hasLeft {
		if done[gearNum-1] == false && group[gearNum-1][2] != group[gearNum][6] {
			clockwise(gearNum - 1)
		}
	}
	if hasRight {
		if done[gearNum+1] == false && group[gearNum+1][6] != group[gearNum][2] {
			clockwise(gearNum + 1)
		}
	}
	tmp := group[gearNum][0]
	for i := 0; i < totalValuePerGear-1; i++ {
		group[gearNum][i] = group[gearNum][i+1]
	}
	group[gearNum][totalValuePerGear-1] = tmp

}

func Start() {
	for i := range group {
		var str string
		fmt.Scan(&str)
		for j := range group[i] {
			val, _ := strconv.Atoi(string(str[j]))
			group[i][j] = val
		}
	}
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		for j := range done {
			done[j] = false
		}
		var a, b int
		fmt.Scan(&a, &b)
		if b == 1 {
			clockwise(a - 1)
		} else {
			counterclockwise(a - 1)
		}
		for i := 0; i < 4; i++ {
			for j := 0; j < 8; j++ {
				fmt.Printf("%d ", group[i][j])
			}
			fmt.Println()
		}
	}
	answer := 0
	for i := 0; i < totalGearCnt; i++ {
		if group[i][0] == 1 {
			answer += int(math.Pow(2, float64(i)))
		}
	}
	fmt.Println(answer)
}
