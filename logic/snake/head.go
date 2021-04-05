package snake

import "github.com/dionsaputra/fp-snake/logic/geometry"

type Head struct {
	segment   Segment
	direction geometry.Direction
}

func NewHead(segment Segment, direction geometry.Direction) Head {
	return Head{
		segment:   segment,
		direction: direction,
	}
}

func (h Head) SetDirection(direction geometry.Direction) Head {
	if !h.direction.IsOpposite(direction) {
		h.direction = direction
	}
	return h
}

func (h Head) Move(dimension geometry.Dimension) Head {
	h.segment = h.segment.Move(h.direction, dimension)
	return h
}
