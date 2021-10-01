package q86052

import (
	"sort"
)

var rowSize int
var colSize int
var globalMap []string
var check map[light]struct{}

type position struct {
	row, col int
}

func (p *position) getRightPosition() {
	if p.row == -1 {
		p.row = rowSize - 1
	}
	if p.row == rowSize {
		p.row = 0
	}
	if p.col == -1 {
		p.col = colSize - 1
	}
	if p.col == colSize {
		p.col = 0
	}
}

type light struct {
	pos position
	dir int // 0 위, 1 오른쪽, 2 아래, 3 왼쪽
}

func (l light) getNextLight() light {
	nextRow := l.pos.row
	nextCol := l.pos.col
	nextDir := l.dir
	switch l.dir {
	case 0:
		nextRow -= 1
	case 1:
		nextCol += 1
	case 2:
		nextRow += 1
	case 3:
		nextCol -= 1
	}

	pos := position{
		row: nextRow,
		col: nextCol,
	}
	pos.getRightPosition()
	switch globalMap[pos.row][pos.col] {
	case 'L':
		nextDir--
	case 'R':
		nextDir++
	case 'S':

	}

	if nextDir == -1 {
		nextDir = 3
	}
	if nextDir == 4 {
		nextDir = 0
	}

	return light{
		pos,
		nextDir,
	}

}

var footPrints [][]light
var idx int

func findSameLight(footPrints []light, indexMap map[light]int, aLight light) int {
	if _, ok := indexMap[aLight]; ok {
		return indexMap[aLight]
	}
	check[aLight] = struct{}{}

	indexMap[aLight] = idx
	idx++
	return -1
}

func solution(grid []string) []int {
	var answer []int
	globalMap = grid
	rowSize = len(grid)
	colSize = len(grid[0])
	check = make(map[light]struct{})
	index := 0
	for r := 0; r < rowSize; r++ {
		for c := 0; c < colSize; c++ {
			for d := 0; d < 4; d++ {

				if _, ok := check[light{pos: position{row: r, col: c}, dir: d}]; ok {
					continue
				}
				footPrints = append(footPrints, make([]light, 0))
				footPrints[index] = append(footPrints[index], light{pos: position{row: r, col: c}, dir: d})
				idx = 0
				indexMap := make(map[light]int)
				for {
					nextLight := footPrints[index][len(footPrints[index])-1].getNextLight()
					val := findSameLight(footPrints[index], indexMap, nextLight)
					if val == -1 {
						footPrints[index] = append(footPrints[index], nextLight)
					} else {
						answer = append(answer, len(footPrints[index])-1)
						check[nextLight] = struct{}{}
						break
					}
				}
				index++
			}
		}
	}

	sort.Ints(answer)
	return answer
}

func Start() {
	grid := []string{"SL", "LR"}
	solution(grid)
}
