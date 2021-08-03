package q2407

import (
	"fmt"
	"strconv"
)

//큰수의 합 문제

const (
	max = 101
)

var answer string
var N, M string
var combi [max][max]string

func getInt(val byte) int {
	value, _ := strconv.Atoi(string(val))
	return value
}

func calc(n, m string) string {
	var answer string
	if len(n) < len(m) {
		n, m = m, n
	}
	//	fmt.Printf("n : %s m : %s\n", n, m)

	long := len(n)
	short := len(m)
	var tmp int
	olim := false

	for i := range n {
		index := i + 1
		var longVal int
		var shortVal int
		if short-index < 0 {
			shortVal = 0
		} else {
			shortVal = getInt(m[short-index])
		}
		longVal = getInt(n[long-index])
		tmp = shortVal + longVal
		if olim {
			tmp++
		}
		if tmp >= 10 {
			olim = true
			answer = strconv.Itoa(tmp%10) + answer

		} else {
			answer = strconv.Itoa(tmp) + answer
			olim = false
		}
	}
	if olim {
		answer = "1" + answer
	}
	return answer
}

func combination(n, m int) string {
	if n == m || m == 0 {
		combi[n][m] = "1"
		return "1"
	}
	if combi[n-1][m-1] == "" {
		combination(n-1, m-1)
	}
	if combi[n-1][m] == "" {
		combination(n-1, m)
	}
	combi[n][m] = calc(combi[n-1][m-1], combi[n-1][m])
	return combi[n][m]
}

func Start() {
	fmt.Scan(&N, &M)
	//	combination(N, M)
	fmt.Println(calc(N, M))
}
