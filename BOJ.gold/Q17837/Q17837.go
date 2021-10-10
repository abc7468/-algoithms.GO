package q17837

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var size int

type Field struct {
	status int
	pieces []*Piece
}

func (f Field) findPieceIdx(p *Piece) (idx int) {
	for i := 0; i < len(f.pieces); i++ {
		if f.pieces[i] == p {
			idx = i
			break
		}
	}
	return idx
}
func (f *Field) movePiece(p *Piece) {

	nextR, nextC := p.getNextPos()
	if !isRange(nextR, nextC) || mmap[nextR][nextC].status == 2 {
		p.toggleDir()
		nextR, nextC = p.getNextPos()
		if !isRange(nextR, nextC) || mmap[nextR][nextC].status == 2 {
			return
		}
	}

	pieceIdx := f.findPieceIdx(p)
	mass := f.pieces[pieceIdx:]

	f.pieces = f.pieces[:pieceIdx]

	if mmap[nextR][nextC].status == 1 {
		for i := 0; i < len(mass)/2; i++ {
			mass[i], mass[len(mass)-i-1] = mass[len(mass)-i-1], mass[i]
		}
	}

	for _, v := range mass {
		v.pos.row = nextR
		v.pos.col = nextC
	}

	mmap[nextR][nextC].pieces = append(mmap[nextR][nextC].pieces, mass...)

}

func isRange(row, col int) bool {
	return row >= 0 && row < size && col >= 0 && col < size
}

type Pos struct {
	row, col int
}

type Piece struct {
	pos Pos
	dir int
}

func (p *Piece) toggleDir() {
	switch p.dir {
	case 1:
		p.dir = 2
	case 2:
		p.dir = 1
	case 3:
		p.dir = 4
	case 4:
		p.dir = 3
	}
}
func (p Piece) getNextPos() (int, int) {
	var nextR int
	var nextC int
	switch p.dir {
	case 1:
		nextR = p.pos.row
		nextC = p.pos.col + 1
	case 2:
		nextR = p.pos.row
		nextC = p.pos.col - 1
	case 3:
		nextR = p.pos.row - 1
		nextC = p.pos.col
	case 4:
		nextR = p.pos.row + 1
		nextC = p.pos.col
	}
	return nextR, nextC
}

var mmap [][]Field
var totalPieces []*Piece

func StrtoInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func Start() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	size = StrtoInt(sc.Text())
	sc.Scan()
	k := StrtoInt(sc.Text())
	mmap = make([][]Field, size)
	totalPieces = make([]*Piece, k)
	for i := 0; i < size; i++ {
		mmap[i] = make([]Field, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sc.Scan()
			mmap[i][j].status = StrtoInt(sc.Text())
		}
	}
	for i := 0; i < k; i++ {
		p := Piece{}
		sc.Scan()
		p.pos.row = StrtoInt(sc.Text()) - 1
		sc.Scan()
		p.pos.col = StrtoInt(sc.Text()) - 1
		sc.Scan()
		p.dir = StrtoInt(sc.Text())
		totalPieces[i] = &p
		mmap[p.pos.row][p.pos.col].pieces = append(mmap[p.pos.row][p.pos.col].pieces, &p)
	}

	for cnt := 1; cnt <= 1000; cnt++ {
		for _, v := range totalPieces {
			mmap[v.pos.row][v.pos.col].movePiece(v)
			if len(mmap[v.pos.row][v.pos.col].pieces) >= 4 {

				fmt.Println(cnt)
				return
			}
		}

	}
	fmt.Println("-1")
}
