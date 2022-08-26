package main

import "fmt"

type color int

const (
	BLACK color = iota
	WHITE
)

type piece struct {
	color color
	isOut bool
	score int
}

type Piece interface {
	openSpots(x, y int, closedSpots [][]int) [][]int
	isEmpty() bool
}

func NewPiece(c color) Piece {
	return &empty{piece{color: c}}
}

type empty struct {
	piece
}

func (e empty) isEmpty() bool {
	return true
}

func (p piece) filterMoves(openSpots [][]int) [][]int {
	var response [][]int
	for _, i := range openSpots {
		if i[0] >= 0 && i[0] <= 7 && i[1] >= 0 && i[1] <= 7 {
			response = append(response, i)
		}
	}
	return response
}

func (e empty) openSpots(x, y int) [][]int {
	return [][]int{}
}

func (e empty) String() string {
	return " "
}

type XCordinate int
type YCordinate int

type spot struct {
	Color      color
	Piece      Piece
	XCordinate int
	YCordinate int
}

type Table [8][8]spot

var initPieceList = []Piece{
	NewRook(BLACK),
	NewKnight(BLACK),
	NewBishop(BLACK),
	NewKing(BLACK),
	NewQueen(BLACK),
	NewBishop(BLACK),
	NewKnight(BLACK),
	NewRook(BLACK),
}

func NewTable() Table {
	var table Table
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			table[i][j] = spot{Color: color(j % 2), XCordinate: j, YCordinate: i, Piece: NewPiece(color(j % 2))}
			if i == 0 || i == 7 {
				table[i][j].Piece = initPieceList[j]
			}
			if i == 1 || i == 6 {
				table[i][j].Piece = NewPawn(color(j % 2))
			}
		}
	}
	return table
}

func (t Table) bOrWPiece(y int) color {
	if y > 4 {
		return WHITE
	}
	return BLACK
}
func (t Table) String() string {
	str := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			str = fmt.Sprintf("%s %s", str, t[i][j].Piece)
		}
		str = fmt.Sprintf("%s\n", str)
	}
	return str
}

func (t Table) Move(xS, yS, xD, yD int) {
	openSpots := t[yS][xS].Piece.openSpots(xS, yS, t.ListClosedSpots())
}
func (t Table) ListClosedSpots() [][]int {
	var response [][]int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if !t[i][j].Piece.isEmpty() {
				response = append(response, []int{i, j})
			}
		}
	}
	return response
}
func main() {

	t := NewTable()
	fmt.Printf("%s", t)
	t.Move(0, 0, 1, 1)
}
