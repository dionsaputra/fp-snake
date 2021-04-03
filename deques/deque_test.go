package deques

import (
	"reflect"
	"testing"
)

func TestNewDeque(t *testing.T) {
	type args struct {
		elements []interface{}
	}
	tests := []struct {
		name string
		args args
		want *Deque
	}{
		{
			name: "empty elements",
			args: args{elements: []interface{}{}},
			want: &Deque{},
		},
		{
			name: "not empty elements",
			args: args{elements: []interface{}{1, 2, 3}},
			want: &Deque{elements: []interface{}{1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DequeOf(tt.args.elements...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DequeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_IsEmpty(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "empty",
			fields: fields{elements: []interface{}{}},
			want:   true,
		},
		{
			name:   "not empty",
			fields: fields{elements: []interface{}{1, 2, 3}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deque{
				elements: tt.fields.elements,
			}
			if got := d.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PopBack(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Deque
	}{
		{
			name:   "empty",
			fields: fields{elements: []interface{}{}},
			want:   &Deque{elements: []interface{}{}},
		},
		{
			name:   "not empty",
			fields: fields{elements: []interface{}{1, 2, 3}},
			want:   DequeOf(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deque{
				elements: tt.fields.elements,
			}
			if got := d.PopBack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PopFront(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Deque
	}{
		{
			name:   "empty",
			fields: fields{elements: []interface{}{}},
			want:   &Deque{elements: []interface{}{}},
		},
		{
			name:   "not empty",
			fields: fields{elements: []interface{}{1, 2, 3}},
			want:   DequeOf(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deque{
				elements: tt.fields.elements,
			}
			if got := d.PopFront(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PushBack(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	type args struct {
		element interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Deque
	}{
		{
			name:   "push back",
			fields: fields{elements: []interface{}{1, 2}},
			args:   args{element: 3},
			want:   DequeOf(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deque{
				elements: tt.fields.elements,
			}
			if got := d.PushBack(tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PushFront(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	type args struct {
		element interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Deque
	}{
		{
			name:   "push back",
			fields: fields{elements: []interface{}{1, 2}},
			args:   args{element: 3},
			want:   DequeOf(3, 1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deque{
				elements: tt.fields.elements,
			}
			if got := d.PushFront(tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_Size(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "empty",
			fields: fields{elements: []interface{}{}},
			want:   0,
		},
		{
			name:   "not empty",
			fields: fields{elements: []interface{}{1, 2, 3}},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deque{
				elements: tt.fields.elements,
			}
			if got := d.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_Contains(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	type args struct {
		predicate func(element interface{}) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "contains",
			fields: fields{elements: []interface{}{1, 2, 3}},
			args: args{predicate: func(element interface{}) bool {
				return element == 3
			}},
			want: true,
		},
		{
			name:   "not contains",
			fields: fields{elements: []interface{}{1,2,3}},
			args:   args{predicate: func(element interface{}) bool {
				return element == 4
			}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deque{
				elements: tt.fields.elements,
			}
			if got := d.Contains(tt.args.predicate); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}