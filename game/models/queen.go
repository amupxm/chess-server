package main

type Queen struct {
	piece
}

func NewQueen(c color) Piece {
	return Queen{piece{color: c, isOut: false}}
}

func (q Queen) String() string {
	return "Q"
}

func (q Queen) openSpots(x, y int) [][]int {
	return [][]int{}
}

func (k Queen) isEmpty() bool {
	return false
}
