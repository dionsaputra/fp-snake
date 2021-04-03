package snakes

import (
	"github.com/dionsaputra/fp-snake/deques"
	"github.com/dionsaputra/fp-snake/points"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewSnake(t *testing.T) {
	type args struct {
		head points.Point
	}
	tests := []struct {
		name string
		args args
		want *Snake
	}{
		{
			name: "SnakeOf",
			args: args{head: points.PointOf(1, 2)},
			want: &Snake{
				head: points.PointOf(1, 2),
				tail: deques.DequeOf(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeOf(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SnakeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnake_Move(t *testing.T) {
	type fields struct {
		head points.Point
		tail *deques.Deque
	}
	type args struct {
		dx int
		dy int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Snake
	}{
		{
			name: "just head",
			fields: fields{
				head: points.PointOf(2, 3),
				tail: deques.DequeOf(),
			},
			args: args{
				dx: 1,
				dy: 0,
			},
			want: SnakeOf(points.PointOf(3, 3)),
		},
		{
			name: "with tail",
			fields: fields{
				head: points.PointOf(2, 3),
				tail: deques.DequeOf(points.PointOf(2, 4), points.PointOf(2, 5)),
			},
			args: args{
				dx: 1,
				dy: 0,
			},
			want: &Snake{
				head: points.PointOf(3, 3),
				tail: deques.DequeOf(points.PointOf(2, 3), points.PointOf(2, 4)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Snake{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			got := s.Move(tt.args.dx, tt.args.dy)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSnake_Contains(t *testing.T) {
	type fields struct {
		head points.Point
		tail *deques.Deque
	}
	type args struct {
		point points.Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "contains in head",
			fields: fields{
				head: points.PointOf(3, 4),
				tail: deques.DequeOf(points.PointOf(3, 5), points.PointOf(3, 6)),
			},
			args: args{point: points.PointOf(3, 4)},
			want: true,
		},
		{
			name: "contains in tail",
			fields: fields{
				head: points.PointOf(3, 4),
				tail: deques.DequeOf(points.PointOf(3, 5), points.PointOf(3, 6)),
			},
			args: args{point: points.PointOf(3, 6)},
			want: true,
		},
		{
			name: "not contains",
			fields: fields{
				head: points.PointOf(3, 4),
				tail: deques.DequeOf(points.PointOf(3, 5), points.PointOf(3, 6)),
			},
			args: args{point: points.PointOf(3, 7)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Snake{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := s.Contains(tt.args.point); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
