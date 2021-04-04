package logics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSnake(t *testing.T) {
	type args struct {
		head Segment
	}
	tests := []struct {
		name string
		args args
		want Snake
	}{
		{
			name: "new snake",
			args: args{NewSegment(10, 20)},
			want: Snake{head: NewSegment(10, 20)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewSnake(tt.args.head))
		})
	}
}

func TestSnake_Move(t *testing.T) {
	type fields struct {
		head Segment
		tail Tail
	}
	type args struct {
		direction Direction
		dimension Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Snake
	}{
		{
			name:   "just head",
			fields: fields{head: NewSegment(10, 20)},
			args:   args{NewDirection(-1, 0), NewDimension(20, 30)},
			want:   NewSnake(NewSegment(9, 20)),
		},
		{
			name: "one tail",
			fields: fields{
				head: NewSegment(10, 20),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{NewDirection(-1, 0), NewDimension(20, 30)},
			want: Snake{
				head: NewSegment(9, 20),
				tail: NewTail(NewSegment(10, 20)),
			},
		},
		{
			name: "two tail",
			fields: fields{
				head: NewSegment(10, 20),
				tail: NewTail(
					NewSegment(11, 20),
					NewSegment(12, 20),
				),
			},
			args: args{NewDirection(-1, 0), NewDimension(20, 30)},
			want: Snake{
				head: NewSegment(9, 20),
				tail: NewTail(
					NewSegment(10, 20),
					NewSegment(11, 20),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Snake{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			assert.Equal(t, tt.want, s.Move(tt.args.direction, tt.args.dimension))
		})
	}
}

func TestSnake_Grow(t *testing.T) {
	type fields struct {
		head Segment
		tail Tail
	}
	type args struct {
		direction Direction
		dimension Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Snake
	}{
		{
			name:   "just head",
			fields: fields{head: NewSegment(10, 20)},
			args:   args{NewDirection(-1, 0), NewDimension(20, 30)},
			want: Snake{
				head: NewSegment(9, 20),
				tail: NewTail(NewSegment(10, 20)),
			},
		},
		{
			name: "one tail",
			fields: fields{
				head: NewSegment(10, 20),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{NewDirection(-1, 0), NewDimension(20, 30)},
			want: Snake{
				head: NewSegment(9, 20),
				tail: NewTail(
					NewSegment(10, 20),
					NewSegment(11, 20),
				),
			},
		},
		{
			name: "two tail",
			fields: fields{
				head: NewSegment(10, 20),
				tail: NewTail(
					NewSegment(11, 20),
					NewSegment(12, 20),
				),
			},
			args: args{NewDirection(-1, 0), NewDimension(20, 30)},
			want: Snake{
				head: NewSegment(9, 20),
				tail: NewTail(
					NewSegment(10, 20),
					NewSegment(11, 20),
					NewSegment(12, 20),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Snake{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			assert.Equal(t, tt.want, s.Grow(tt.args.direction, tt.args.dimension))
		})
	}
}

func TestSnake_Contains(t *testing.T) {
	type fields struct {
		head Segment
		tail Tail
	}
	type args struct {
		segment Segment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "equals with head",
			fields: fields{
				head: NewSegment(10, 20),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{NewSegment(10, 20)},
			want: true,
		},
		{
			name: "contains in tail",
			fields: fields{
				head: NewSegment(10, 20),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{NewSegment(11, 20)},
			want: true,
		},
		{
			name: "not contains",
			fields: fields{
				head: NewSegment(10, 20),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{NewSegment(12, 20)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Snake{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			assert.Equal(t, tt.want, s.Contains(tt.args.segment))
		})
	}
}