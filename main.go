package main

import "fmt"

func isUnique(scores [][]int, standard int, studentNum int) bool {
	size := len(scores)
	for i := 0; i < size; i++ {
		if studentNum == i {
			continue
		}
		if standard == scores[i][studentNum] {
			return false
		}
	}
	return true
}
func isSmall(scores [][]int, standard int, studentNum int) bool {
	size := len(scores)
	for i := 0; i < size; i++ {
		if studentNum == i {
			continue
		}
		if standard > scores[i][studentNum] {
			return false
		}
	}
	return true
}
func isBig(scores [][]int, standard int, studentNum int) bool {
	size := len(scores)
	for i := 0; i < size; i++ {
		if studentNum == i {
			continue
		}
		if standard < scores[i][studentNum] {
			return false
		}
	}
	return true
}

func solution(scores [][]int) string {
	answer := ""
	size := len(scores)
	studentScores := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				unique := isUnique(scores, scores[i][j], j)
				small := isSmall(scores, scores[i][j], j)
				big := isBig(scores, scores[i][j], j)
				if unique && small {
					continue
				} else if unique && big {
					continue
				}
			}
			studentScores[j] = append(studentScores[j], scores[i][j])
		}
	}
	for i := 0; i < size; i++ {
		tmp := 0
		for j := 0; j < len(studentScores[i]); j++ {
			tmp += studentScores[i][j]
		}
		total := tmp / len(studentScores[i])
		if total >= 90 {
			answer += "A"
		} else if total >= 80 && total < 90 {
			answer += "B"

		} else if total >= 70 && total < 80 {
			answer += "C"

		} else if total >= 50 && total < 70 {
			answer += "D"

		} else {
			answer += "E"

		}
	}

	return answer
}

func main() {
	var a [][]int
	a = [][]int{{100, 90, 98, 88, 65}, {50, 45, 99, 85, 77}, {47, 88, 95, 80, 67}, {61, 57, 100, 80, 65}, {24, 90, 94, 75, 65}}
	fmt.Println(solution(a))

}
