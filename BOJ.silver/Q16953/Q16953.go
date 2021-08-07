package q16953

import (
	"fmt"
	"strconv"
)

var input string
var isEnd bool

func stringToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func dfsForFindAnswer(val string) {
	if val == input {
		isEnd = true
	}
	if isEnd == true {
		return
	}

	value := stringToInt(val)

	if value > stringToInt(input) {
		return
	}
	mulVal := fmt.Sprint(value * 2)
	val = val + "1"
	dfsForFindAnswer(mulVal)
	dfsForFindAnswer(val)

}

func Start() {
	fmt.Scan(&input)
	fmt.Printf("%s", input)
}
