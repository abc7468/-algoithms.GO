package q12952

import "fmt"

var board [][]bool
var colCheck []bool

func initBoard(n int) {
	board = make([][]bool, n)
	colCheck = make([]bool, n)
	for i := 0; i < len(board); i++ {
		board[i] = make([]bool, n)
	}

}

var answer int

func goBackTracking(size, row int) {
	if row == size {
		fmt.Println(board)
		answer++
		return
	}
	for i := 0; i < size; i++ {
		if !colCheck[i] && checkBoard(size, row, i) {
			board[row][i] = true
			colCheck[i] = true
			// fmt.Printf("row : %d\ncol : %d\n\n", row, i)
			goBackTracking(size, row+1)
			colCheck[i] = false
			board[row][i] = false
		}
	}
}

func checkBoard(size, row, col int) bool {
	if !diagnolRightDown(size, row, col) {
		return false

	}
	if !diagnolLeftDown(size, row, col) {
		return false
	}
	return true
}

func diagnolRightDown(size, row, col int) bool {
	for row >= 0 && col >= 0 {
		if board[row][col] {
			return false
		}
		row--
		col--
	}

	return true
}
func diagnolLeftDown(size, row, col int) bool {
	for row >= 0 && col < size {

		if board[row][col] {
			return false
		}
		row--
		col++
	}
	return true
}

func solution(n int) int {
	initBoard(n)
	goBackTracking(n, 0)
	return answer
}

func Start() {
	fmt.Print(solution(4))
}
