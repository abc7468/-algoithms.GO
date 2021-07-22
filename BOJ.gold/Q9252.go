package bojgold

import (
	"fmt"
)

func main() {
	var str1 string
	var str2 string
	// fmt.Scanln(&str1)
	// fmt.Scanln(&str2)
	str1 = "CCCBBBACA"
	str2 = "AAACCCABA"
	var answer string
	var dp [1001][1001]int
	var tmp [1001][1001]int
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				tmp[i][j] = 5
			} else {
				if dp[i][j-1] >= dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
					tmp[i][j] = 3
				} else {
					dp[i][j] = dp[i-1][j]
					tmp[i][j] = 6
				}
			}
		}
	}
	val := dp[len(str1)][len(str2)]
	if val == 0 {
		fmt.Print(val)
		return
	}
	y := len(str1)
	x := len(str2)
	for true {
		switch tmp[y][x] {
		case 5:
			x--
			y--
			val--
			answer += string(str1[y])
		case 3:
			x--

		case 6:
			y--
		}
		if val == 0 {
			break
		}
	}
	fmt.Println(dp[len(str1)][len(str2)])

	for i := len(answer) - 1; i >= 0; i-- {
		fmt.Print(string(answer[i]))
	}
}
