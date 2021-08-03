package q2638

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var H, W int
var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, -1, 0, 1}
var visited = [100][100]bool{}
var touch = [100][100]int{}
var board = [100][100]int{}
var cheeseCnt int

func initializing() {
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			visited[i][j] = false
			touch[i][j] = 0
		}
	}
}

func dfs(row, col int) {
	visited[row][col] = true
	nowRow := row
	nowCol := col
	for k := 0; k < 4; k++ {
		nextRow := nowRow + dy[k]
		nextCol := nowCol + dx[k]
		if nextCol >= 0 && nextCol < W && nextRow >= 0 && nextRow < H {
			if board[nextRow][nextCol] == 0 { //빈 칸일때
				if !visited[nextRow][nextCol] {
					dfs(nextRow, nextCol)
				}
			} else { //치즈일때
				touch[nextRow][nextCol]++
				if touch[nextRow][nextCol] >= 2 {
					cheeseCnt--
					board[nextRow][nextCol] = 0
					visited[nextRow][nextCol] = true
					if cheeseCnt == 0 {
						return
					}
				}
			}
		}
	}
}

func Q2638() {
	fmt.Scan(&H, &W)
	answer := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			scanner.Scan()
			board[h][w], _ = strconv.Atoi(scanner.Text())
			if board[h][w] == 1 {
				cheeseCnt++
			}
		}
	}
	for true {
		dfs(0, 0)
		answer++
		initializing()
		if cheeseCnt == 0 {
			break
		}
	}
	fmt.Print(answer)

}
