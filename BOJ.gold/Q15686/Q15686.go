package q15686

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	inf int = 987654321
)

type position struct {
	x, y int
}

var chicken []position
var home []position
var candi []int
var N, M int
var min = inf

// 치킨가게 조합 구한다.
// home을 반복문 돌며 치킨과 거리 구한다.
// 최솟값 리턴

func stringToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
func combination(start int) {
	if len(candi) == M {
		var disSum int
		//for문
		for i := range home {
			minChoice := inf
			for _, j := range candi {
				tmp := getDistance(home[i], chicken[j])
				if minChoice > tmp {
					minChoice = tmp
				}
			}
			disSum += minChoice
		}
		if min > disSum {
			min = disSum
		}
		return
	}

	for i := start; i < len(chicken); i++ {
		candi = append(candi, i)
		combination(i + 1)
		candi = candi[:len(candi)-1]
	}
}

func getDistance(homePosition, chickenPosition position) int {
	x := int(math.Abs(float64(homePosition.x - chickenPosition.x)))
	y := int(math.Abs(float64(homePosition.y - chickenPosition.y)))
	return x + y
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fmt.Scan(&N, &M)
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			scanner.Scan()
			input := stringToInt(scanner.Text())
			if input == 1 {
				home = append(home, position{col, row})
			} else if input == 2 {
				chicken = append(chicken, position{col, row})
			}
		}
	}
	combination(0)
	fmt.Println(min)
}
