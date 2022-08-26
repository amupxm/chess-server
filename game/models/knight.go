package main

type Knight struct {
	piece
}

func (k Knight) isEmpty() bool {
	return false
}

func NewKnight(c color) Piece {
	return Knight{piece{color: c, isOut: false}}
}

func (k Knight) String() string {
	return "N"
}

func (e Knight) openSpots(x, y int) [][]int {
	return [][]int{}
}
