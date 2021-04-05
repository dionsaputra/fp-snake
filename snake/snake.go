package snake

import (
	"github.com/dionsaputra/fp-snake/math"
)

type (
	Snake struct {
		head Head
		tail Tail
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

func NewSnake(head Head) Snake {
	return Snake{head: head}
}

func (s Snake) Move(dimension math.Dimension) Snake {
	if !s.tail.IsEmpty() {
		s.tail = s.tail.AddFirst(s.head.Segment).DropLast()
	}
	s.head = s.head.Move(dimension)
	return s
}

func (s Snake) Grow(dimension math.Dimension) Snake {
	s.tail = s.tail.AddFirst(s.head.Segment)
	s.head = s.head.Move(dimension)
	return s
}

func (s Snake) Contains(segment Segment) bool {
	return s.head.Segment == segment || s.tail.Contains(segment)
}

func (s Segment) Move(direction math.Direction, dimension math.Dimension) Segment {
	return Segment{
		math.Mod(s.Row+direction.Vertical, dimension.Height),
		math.Mod(s.Col+direction.Horizontal, dimension.Width),
	}
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

func (t Tail) IsEmpty() bool {
	return len(t.Segments) == 0
}

func (t Tail) Contains(segment Segment) bool {
	for _, s := range t.Segments {
		if s == segment {
			return true
		}
	}
	return false
}