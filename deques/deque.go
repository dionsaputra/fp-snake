package deques

import (
	"github.com/dionsaputra/fp-snake/ranges"
	"github.com/dionsaputra/fp-snake/routines"
)

type Deque struct {
	elements []interface{}
}

func DequeOf(elements ...interface{}) *Deque {
	d := Deque{}
	ranges.NewRange(len(elements)).Forward(func(index int) {
		d.PushBack(elements[index])
	})
	return &d
}

func (d *Deque) Size() int {
	return len(d.elements)
}

func (d *Deque) IsEmpty() bool {
	return d.Size() == 0
}

func (d *Deque) PushFront(element interface{}) *Deque {
	d.elements = append([]interface{}{element}, d.elements...)
	return d
}

func (d *Deque) PushBack(element interface{}) *Deque {
	d.elements = append(d.elements, element)
	return d
}

func (d *Deque) PopFront() *Deque {
	routines.DoIf(!d.IsEmpty(), d.doPopFront())
	return d
}

func (d *Deque) doPopFront() func() {
	return func() {
		d.elements = d.elements[1:]
	}
}

func (d *Deque) PopBack() *Deque {
	routines.DoIf(!d.IsEmpty(), d.doPopBack())
	return d
}

func (d *Deque) doPopBack() func() {
	return func() {
		d.elements = d.elements[:d.Size()-1]
	}
}

func (d *Deque) Contains(predicate func(element interface{}) bool) bool {
	return ranges.NewRange(d.Size()).Contains(func(index int) bool {
		return predicate(d.elements[index])
	})
}