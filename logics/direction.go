package logics

type Direction struct {
	deltaRow int
	deltaCol int
}

func NewDirection(deltaRow, deltaCol int) Direction {
	return Direction{deltaRow, deltaCol}
}
