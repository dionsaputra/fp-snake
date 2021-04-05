package snake

import (
	"github.com/dionsaputra/fp-snake/logic/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSnake(t *testing.T) {
	type args struct {
		head Head
	}
	tests := []struct {
		name string
		args args
		want Snake
	}{
		{
			name: "new snake",
			args: args{NewHead(NewSegment(10, 20), geometry.Left())},
			want: Snake{head: NewHead(NewSegment(10, 20), geometry.Left())},
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
		head Head
		tail Tail
	}
	type args struct {
		dimension geometry.Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Snake
	}{
		{
			name:   "just head",
			fields: fields{head: NewHead(NewSegment(10, 20), geometry.Up())},
			args:   args{geometry.NewDimension(20, 30)},
			want:   NewSnake(NewHead(NewSegment(9, 20), geometry.Up())),
		},
		{
			name: "one tail",
			fields: fields{
				head: NewHead(NewSegment(10, 20), geometry.Up()),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{geometry.NewDimension(20, 30)},
			want: Snake{
				head: NewHead(NewSegment(9, 20), geometry.Up()),
				tail: NewTail(NewSegment(10, 20)),
			},
		},
		{
			name: "two tail",
			fields: fields{
				head: NewHead(NewSegment(10, 20), geometry.Up()),
				tail: NewTail(
					NewSegment(11, 20),
					NewSegment(12, 20),
				),
			},
			args: args{geometry.NewDimension(20, 30)},
			want: Snake{
				head: NewHead(NewSegment(9, 20), geometry.Up()),
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
			assert.Equal(t, tt.want, s.Move(tt.args.dimension))
		})
	}
}

func TestSnake_Grow(t *testing.T) {
	type fields struct {
		head Head
		tail Tail
	}
	type args struct {
		dimension geometry.Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Snake
	}{
		{
			name:   "just head",
			fields: fields{head: NewHead(NewSegment(10, 20), geometry.Up())},
			args:   args{geometry.NewDimension(20, 30)},
			want: Snake{
				head: NewHead(NewSegment(9, 20), geometry.Up()),
				tail: NewTail(NewSegment(10, 20)),
			},
		},
		{
			name: "one tail",
			fields: fields{
				head: NewHead(NewSegment(10, 20), geometry.Up()),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{geometry.NewDimension(20, 30)},
			want: Snake{
				head: NewHead(NewSegment(9, 20), geometry.Up()),
				tail: NewTail(
					NewSegment(10, 20),
					NewSegment(11, 20),
				),
			},
		},
		{
			name: "two tail",
			fields: fields{
				head: NewHead(NewSegment(10, 20), geometry.Up()),
				tail: NewTail(
					NewSegment(11, 20),
					NewSegment(12, 20),
				),
			},
			args: args{geometry.NewDimension(20, 30)},
			want: Snake{
				head: NewHead(NewSegment(9, 20), geometry.Up()),
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
			assert.Equal(t, tt.want, s.Grow(tt.args.dimension))
		})
	}
}

func TestSnake_Contains(t *testing.T) {
	type fields struct {
		head Head
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
				head: NewHead(NewSegment(10, 20), geometry.Up()),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{NewSegment(10, 20)},
			want: true,
		},
		{
			name: "contains in tail",
			fields: fields{
				head: NewHead(NewSegment(10, 20), geometry.Up()),
				tail: NewTail(NewSegment(11, 20)),
			},
			args: args{NewSegment(11, 20)},
			want: true,
		},
		{
			name: "not contains",
			fields: fields{
				head: NewHead(NewSegment(10, 20), geometry.Up()),
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