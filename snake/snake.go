package snake

import (
	"github.com/dionsaputra/fp-snake/math"
)

type (
	Snake struct {
		Head Head
		Tail Tail
	}
	Segment struct {
		Row int
		Col int
	}
	Head struct {
		Segment   Segment
		Direction math.Direction
	}
	Tail struct {
		Segments []Segment
	}
)

func (s Snake) Move(dimension math.Dimension) Snake {
	if len(s.Tail.Segments) > 0 {
		s.Tail = s.Tail.AddFirst(s.Head.Segment).DropLast()
	}
	s.Head = s.Head.Move(dimension)
	return s
}

func (s Snake) Grow(dimension math.Dimension) Snake {
	s.Tail = s.Tail.AddFirst(s.Head.Segment)
	s.Head = s.Head.Move(dimension)
	return s
}

func (s Snake) Contains(segment Segment) bool {
	return s.Head.Segment == segment || s.Tail.Contains(segment)
}

func (s Segment) Move(direction math.Direction, dimension math.Dimension) Segment {
	s.Row = math.Mod(s.Row+direction.Vertical, dimension.Height);
	s.Col = math.Mod(s.Col+direction.Horizontal, dimension.Width)
	return s
}

func (h Head) SetDirection(direction math.Direction) Head {
	if !h.Direction.IsOpposite(direction) {
		h.Direction = direction
	}
	return h
}

func (h Head) Move(dimension math.Dimension) Head {
	h.Segment = h.Segment.Move(h.Direction, dimension)
	return h
}

func (t Tail) AddFirst(segment Segment) Tail {
	t.Segments = append([]Segment{segment}, t.Segments...)
	return t
}

func (t Tail) DropLast() Tail {
	size := len(t.Segments)
	if size > 0 {
		t.Segments = t.Segments[:size-1]
	}
	return t
}

func (t Tail) Contains(segment Segment) bool {
	for _, s := range t.Segments {
		if s == segment {
			return true
		}
	}
	return false
}
