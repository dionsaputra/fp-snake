package deques

import (
	"github.com/dionsaputra/fp-snake/fp"
)

type Deque struct {
	elements []interface{}
}

func DequeOf(elements ...interface{}) *Deque {
	d := Deque{}
	fp.RangeOf(len(elements)).Forward(func(index int) {
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
	fp.DoIf(!d.IsEmpty(), func() {
		d.elements = d.elements[1:]
	})
	return d
}

func (d *Deque) PopBack() *Deque {
	fp.DoIf(!d.IsEmpty(), func() {
		d.elements = d.elements[:d.Size()-1]
	})
	return d
}

func (d *Deque) Contains(predicate func(element interface{}) bool) bool {
	return fp.RangeOf(d.Size()).Contains(func(index int) bool {
		return predicate(d.elements[index])
	})
}

func (d *Deque) First() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.elements[0]
}
