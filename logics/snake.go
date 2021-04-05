package logics

type Snake struct {
	head Head
	tail Tail
}

func NewSnake(head Head) Snake {
	return Snake{head: head, tail: NewTail()}
}

func (s Snake) Move(dimension Dimension) Snake {
	if !s.tail.IsEmpty() {
		s.tail = s.tail.AddFirst(s.head.segment).DropLast()
	}
	s.head = s.head.Move(dimension)
	return s
}

func (s Snake) Grow(dimension Dimension) Snake {
	s.tail = s.tail.AddFirst(s.head.segment)
	s.head = s.head.Move(dimension)
	return s
}

func (s Snake) Contains(segment Segment) bool {
	return s.head.segment == segment || s.tail.Contains(segment)
}