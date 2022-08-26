package main

type Pawn struct {
	piece
}

func (p Pawn) isEmpty() bool {
	return false
}

func NewPawn(c color) Piece {
	return Pawn{piece{color: c, isOut: false}}
}

func (p Pawn) String() string {
	return "P"
}

func (p Pawn) openSpots(x, y int) [][]int {
	step := 2
	direction := -1
	var moves [][]int
	if x == 1 || x == 6 {
		step += 1
	}
	if p.color == BLACK {
		direction = 1
	}
	for i := 1; i <= step; i++ {
		moves = append(moves, []int{x, y + (direction * i)})
	}
	return p.piece.filterMoves(moves)
}
