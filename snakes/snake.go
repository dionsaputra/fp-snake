package snakes

import (
	"github.com/dionsaputra/fp-snake/deques"
	"github.com/dionsaputra/fp-snake/points"
)

type Snake struct {
	head points.Point
	tail *deques.Deque
}

func SnakeOf(head points.Point) *Snake {
	return &Snake{
		head: head,
		tail: deques.DequeOf(),
	}
}

func (s *Snake) Tail() *deques.Deque {
	return s.tail
}

func (s *Snake) WithTail(points ...points.Point) *Snake {
	s.tail = deques.DequeOf(points)
	return s
}

func (s *Snake) hasTail() bool {
	return !s.tail.IsEmpty()
}

func (s *Snake) mapIf(condition bool, transform func(snake *Snake) *Snake) *Snake {
	if condition {
		return transform(s)
	}
	return s
}

func (s *Snake) Move(dx, dy int) *Snake {
	return SnakeOf(s.head.Translate(dx, dy)).mapIf(s.hasTail(), s.followHead())
}

func (s *Snake) followHead() func(*Snake) *Snake {
	return func(snake *Snake) *Snake {
		snake.tail = s.tail.PushFront(s.head).PopBack()
		return snake
	}
}

func (s *Snake) Contains(point points.Point) bool {
	return s.head == point || s.tail.Contains(func(element interface{}) bool {
		return element == point
	})
}
