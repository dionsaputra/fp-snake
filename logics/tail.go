package logics

type Tail struct {
	segments []Segment
}

func NewTail(segments ...Segment) Tail {
	return Tail{segments}
}

func (t Tail) AddFirst(segment Segment) Tail {
	t.segments = append([]Segment{segment}, t.segments...)
	return t
}

func (t Tail) DropLast() Tail {
	size := len(t.segments)
	if size > 0 {
		t.segments = t.segments[:size-1]
	}
	return t
}

func (t Tail) IsEmpty() bool {
	return len(t.segments) == 0
}

func (t Tail) Contains(segment Segment) bool {
	for _, s := range t.segments {
		if s == segment {
			return true
		}
	}
	return false
}
