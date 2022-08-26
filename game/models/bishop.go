package main

type Bishop struct {
	piece
}

func NewBishop(c color) Piece {
	return Bishop{piece{color: c, isOut: false}}
}

func (b Bishop) String() string {
	return "B"
}

func (e Bishop) openSpots(x, y int) [][]int {
	return [][]int{}
}

func (k Bishop) isEmpty() bool {
	return false
}
