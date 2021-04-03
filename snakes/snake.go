package snakes

import (
	"github.com/dionsaputra/fp-snake/deques"
	"github.com/dionsaputra/fp-snake/geometry"
)

type Snake struct {
	head geometry.Point
	tail *deques.Deque
}

func SnakeOf(head geometry.Point) *Snake {
	return &Snake{
		head: head,
		tail: deques.DequeOf(),
	}
}

func (s *Snake) Tail() *deques.Deque {
	return s.tail
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

func (s *Snake) Move(direction geometry.Direction) *Snake {
	return SnakeOf(s.head.Translate(direction)).mapIf(s.hasTail(), func(snake *Snake) *Snake {
		snake.tail = s.tail.PushFront(s.head).PopBack()
		return snake
	})
}

func (s *Snake) Contains(point geometry.Point) bool {
	return s.head == point || s.tail.Contains(func(element interface{}) bool {
		return element == point
	})
}
