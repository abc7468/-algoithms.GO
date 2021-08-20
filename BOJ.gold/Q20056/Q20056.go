package q20056

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type field struct {
	fireBall []*fireBall
	stack    []*fireBall
}

type fireBall struct {
	mass  int
	speed int
	dir   int
}

var N, M, K int
var dr = []int{-1, -1, 0, 1, 1, 1, 0, -1}
var dc = []int{0, 1, 1, 1, 0, -1, -1, -1}
var mmap [][]field

func goFireBall() {
	for row := 1; row <= N; row++ {
		for col := 1; col <= N; col++ {
			if len(mmap[row][col].fireBall) == 0 {

			} else if len(mmap[row][col].fireBall) == 1 {
				nextRow := row + dr[mmap[row][col].fireBall[0].dir]*(mmap[row][col].fireBall[0].speed%N)
				nextCol := col + dc[mmap[row][col].fireBall[0].dir]*(mmap[row][col].fireBall[0].speed%N)
				fmt.Printf("beforeRow: %d\nbeforeCol: %d\n", nextRow, nextCol)
				nextRow, nextCol = checkRange(nextRow, nextCol)
				mmap[nextRow][nextCol].stack = append(mmap[nextRow][nextCol].stack, mmap[row][col].fireBall[0])
			} else {
				for i := 0; i < len(mmap[row][col].fireBall); i++ {
					nextRow := row + dr[mmap[row][col].fireBall[i].dir]*(mmap[row][col].fireBall[i].speed%N)
					nextCol := col + dc[mmap[row][col].fireBall[i].dir]*(mmap[row][col].fireBall[i].speed%N)
					nextRow, nextCol = checkRange(nextRow, nextCol)
					mmap[nextRow][nextCol].stack = append(mmap[nextRow][nextCol].stack, mmap[row][col].fireBall[i])
				}
			}
			mmap[row][col].fireBall = nil
		}
	}
}
func mix() {
	for row := 1; row <= N; row++ {
		for col := 1; col <= N; col++ {
			if len(mmap[row][col].stack) == 0 {

			} else if len(mmap[row][col].stack) == 1 {
				mmap[row][col].fireBall = append(mmap[row][col].fireBall, mmap[row][col].stack[0])
			} else {
				madeBalls := makeFireBalls(mmap[row][col].stack)
				if madeBalls != nil {
					mmap[row][col].fireBall = append(mmap[row][col].fireBall, madeBalls...)
				}
			}
			mmap[row][col].stack = nil
		}
	}
}

func makeFireBalls(fireBalls []*fireBall) []*fireBall {
	var madeFireBalls [4]*fireBall
	var amountMass int
	var amountSpeed int
	isSameDir := fireBalls[0].dir % 2
	isSame := true
	fireBallsLen := len(fireBalls)
	for i := 0; i < fireBallsLen; i++ {
		amountMass += fireBalls[i].mass
		amountSpeed += fireBalls[i].speed
		if isSameDir != fireBalls[i].dir%2 {
			isSame = false
		}
	}
	dir := [4]int{0, 2, 4, 6}
	mass := amountMass / 5
	speed := amountSpeed / fireBallsLen
	if !isSame {
		for i := 0; i < 4; i++ {
			dir[i]++
		}
	}
	if mass == 0 {
		return nil
	}
	for i := 0; i < 4; i++ {
		madeFireBalls[i] = &fireBall{mass, speed, dir[i]}
	}
	return madeFireBalls[:]
}

func checkRange(row, col int) (int, int) {
	if row > N {
		row = row - N
	}
	if row <= 0 {
		row = row + N
	}
	if col > N {
		col = col - N
	}
	if col <= 0 {
		col = col + N
	}
	fmt.Printf("row : %d\ncol : %d\n", row, col)
	return row, col
}

func strToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func Start() {
	fmt.Scan(&N, &M, &K)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	mmap = make([][]field, N+1)
	for i := 1; i <= N; i++ {
		mmap[i] = make([]field, N+1)
	}
	for i := 0; i < M; i++ {
		var input [5]int
		for j := 0; j < 5; j++ {
			scanner.Scan()
			input[j] = strToInt(scanner.Text())
		}
		r, c, m, s, d := input[0], input[1], input[2], input[3], input[4]
		mmap[r][c].fireBall = append(mmap[r][c].fireBall, &fireBall{m, s, d})
	}
	for i := 0; i < K; i++ {
		goFireBall()
		mix()
	}
	var answer int
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			for k := 0; k < len(mmap[i][j].fireBall); k++ {
				answer += mmap[i][j].fireBall[k].mass
			}
		}
	}
	fmt.Println(answer)
}
