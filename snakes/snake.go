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

func (s *Snake) WithTail(tail *deques.Deque) *Snake {
	s.tail = tail
	return s
}

func (s *Snake) mapIfHasTail(transform func(snake *Snake) *Snake) *Snake {
	if !s.tail.IsEmpty() {
		return transform(s)
	}
	return s
}

func (s *Snake) Move(direction geometry.Direction) *Snake {
	return SnakeOf(s.head.Translate(direction)).
		WithTail(s.tail).
		mapIfHasTail(func(snake *Snake) *Snake {
			snake.tail = s.tail.PushFront(s.head).PopBack()
			return snake
		})
}

func (s *Snake) Eat() *Snake {
	return SnakeOf(s.head.Translate(s.direction())).
		WithTail(s.tail).
		mapIfHasTail(func(snake *Snake) *Snake {
			snake.tail = s.tail.PushFront(s.head)
			return snake
		})
}

func (s *Snake) direction() geometry.Direction {
	return s.head.GetDirection(s.tail.First().(geometry.Point))
}

func (s *Snake) Contains(point geometry.Point) bool {
	return s.head == point || s.tail.Contains(func(element interface{}) bool {
		return element == point
	})
}
