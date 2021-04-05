package snake

import (
	"github.com/dionsaputra/fp-snake/logic/geometry"
	"github.com/dionsaputra/fp-snake/util"
)

type Segment struct {
	row int
	col int
}

func NewSegment(row, col int) Segment {
	return Segment{row, col}
}

func (s Segment) Move(direction geometry.Direction, dimension geometry.Dimension) Segment {
	return NewSegment(
		util.Mod(s.row+direction.Vertical, dimension.Height),
		util.Mod(s.col+direction.Horizontal, dimension.Width),
	)
}
