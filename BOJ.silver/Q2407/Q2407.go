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
var N, M int
var combi [max][max]string

func getInt(val byte) int {
	value, _ := strconv.Atoi(string(val))
	return value
}

func isOlim(val int) bool {
	if val >= 10 {
		return true
	}
	return false
}

func calc(n, m string) string {
	var answer string
	if len(n) < len(m) {
		n, m = m, n
	}
	long := len(n)
	short := len(m)
	var tmp int
	olim := false

	for i := 1; i <= len(n); i++ {
		var longVal int
		var shortVal int
		if short-i < 0 {
			shortVal = 0
		} else {
			shortVal = getInt(m[short-i])
		}
		longVal = getInt(n[long-i])
		tmp = shortVal + longVal
		if olim {
			tmp++
		}
		olim = isOlim(tmp)
		answer = strconv.Itoa(tmp%10) + answer
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
	combination(N, M)
	fmt.Println(combi[N][M])
}
