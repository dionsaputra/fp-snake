package logics

import "github.com/dionsaputra/fp-snake/utils"

type Segment struct {
	row int
	col int
}

func NewSegment(row, col int) Segment {
	return Segment{row, col}
}

func (s Segment) Move(direction Direction, dimension Dimension) Segment {
	return NewSegment(
		utils.Mod(s.row+direction.vertical, dimension.height),
		utils.Mod(s.col+direction.horizontal, dimension.width),
	)
}
