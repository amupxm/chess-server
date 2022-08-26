package main

type King struct {
	piece
}

func NewKing(c color) Piece {
	return King{piece{color: c, isOut: false}}
}

func (k King) String() string {
	return "K"
}

func (e King) openSpots(x, y int) [][]int {
	return [][]int{}
}
func (k King) isEmpty() bool {
	return false
}
