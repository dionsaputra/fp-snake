package logics

type Head struct {
	segment   Segment
	direction Direction
}

func NewHead(segment Segment, direction Direction) Head {
	return Head{
		segment:   segment,
		direction: direction,
	}
}

func (h Head) SetDirection(direction Direction) Head {
	if !h.direction.IsOpposite(direction) {
		h.direction = direction
	}
	return h
}

func (h Head) Move(dimension Dimension) Head {
	h.segment = h.segment.Move(h.direction, dimension)
	return h
}
