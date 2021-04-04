package logics

type Snake struct {
	head Segment
	tail Tail
}

func NewSnake(head Segment) Snake {
	return Snake{head: head, tail: NewTail()}
}

func (s Snake) Move(direction Direction, dimension Dimension) Snake {
	if !s.tail.IsEmpty() {
		s.tail = s.tail.AddFirst(s.head).DropLast()
	}
	s.head = s.head.Move(direction, dimension)
	return s
}

func (s Snake) Grow(direction Direction, dimension Dimension) Snake {
	s.tail = s.tail.AddFirst(s.head)
	s.head = s.head.Move(direction, dimension)
	return s
}

func (s Snake) Contains(segment Segment) bool {
	return s.head == segment || s.tail.Contains(segment)
}