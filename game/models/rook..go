package main

type Rook struct {
	piece
}

func NewRook(c color) Piece {
	return Rook{piece{color: c, isOut: false}}
}

func (r Rook) String() string {
	return "R"
}

func (p Rook) openSpots(x, y int) [][]int {
	direction := -1
	var moves [][]int

	if p.color == BLACK {
		direction = 1
	}
	for i := 0; i <= 7; i++ {
		moves = append(moves, []int{x, y + (direction * i)})
		moves = append(moves, []int{x + (direction * i), y})
	}
	return p.piece.filterMoves(moves)
}

func (k Rook) isEmpty() bool {
	return false
}
