package snake

import (
	"github.com/dionsaputra/fp-snake/math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnake_Move(t *testing.T) {
	type fields struct {
		head Head
		tail Tail
	}
	type args struct {
		dimension math.Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Snake
	}{
		{
			name:   "just Head",
			fields: fields{head: Head{Segment{10, 20}, math.Up()}},
			args:   args{math.Dimension{Height: 20, Width: 30}},
			want:   Snake{Head: Head{Segment{9, 20}, math.Up()}},
		},
		{
			name: "one Tail",
			fields: fields{
				head: Head{Segment{10, 20}, math.Up()},
				tail: Tail{[]Segment{{11, 20}}},
			},
			args: args{math.Dimension{Height: 20, Width: 30}},
			want: Snake{
				Head: Head{Segment{9, 20}, math.Up()},
				Tail: Tail{[]Segment{{10, 20}}},
			},
		},
		{
			name: "two Tail",
			fields: fields{
				head: Head{Segment{10, 20}, math.Up()},
				tail: Tail{[]Segment{{11, 20}, {12, 20}}},
			},
			args: args{math.Dimension{Height: 20, Width: 30}},
			want: Snake{
				Head: Head{Segment{9, 20}, math.Up()},
				Tail: Tail{[]Segment{{10, 20}, {11, 20}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Snake{
				Head: tt.fields.head,
				Tail: tt.fields.tail,
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
		dimension math.Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Snake
	}{
		{
			name:   "just Head",
			fields: fields{head: Head{Segment{10, 20}, math.Up()}},
			args:   args{math.Dimension{Height: 20, Width: 30}},
			want: Snake{
				Head: Head{Segment{9, 20}, math.Up()},
				Tail: Tail{[]Segment{{10, 20}}},
			},
		},
		{
			name: "one Tail",
			fields: fields{
				head: Head{Segment{10, 20}, math.Up()},
				tail: Tail{[]Segment{{11, 20}}},
			},
			args: args{math.Dimension{Height: 20, Width: 30}},
			want: Snake{
				Head: Head{Segment{9, 20}, math.Up()},
				Tail: Tail{[]Segment{{10, 20}, {11, 20}}},
			},
		},
		{
			name: "two Tail",
			fields: fields{
				head: Head{Segment{10, 20}, math.Up()},
				tail: Tail{[]Segment{{11, 20}, {12, 20}}},
			},
			args: args{math.Dimension{Height: 20, Width: 30}},
			want: Snake{
				Head: Head{Segment{9, 20}, math.Up()},
				Tail: Tail{[]Segment{{10, 20}, {11, 20}, {12, 20}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Snake{
				Head: tt.fields.head,
				Tail: tt.fields.tail,
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
			name: "equals with Head",
			fields: fields{
				head: Head{Segment{10, 20}, math.Up()},
				tail: Tail{[]Segment{{11, 20}}},
			},
			args: args{Segment{10, 20}},
			want: true,
		},
		{
			name: "contains in Tail",
			fields: fields{
				head: Head{Segment{10, 20}, math.Up()},
				tail: Tail{[]Segment{{11, 20}}},
			},
			args: args{Segment{11, 20}},
			want: true,
		},
		{
			name: "not contains",
			fields: fields{
				head: Head{Segment{10, 20}, math.Up()},
				tail: Tail{[]Segment{{11, 20}}},
			},
			args: args{Segment{12, 20}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Snake{
				Head: tt.fields.head,
				Tail: tt.fields.tail,
			}
			assert.Equal(t, tt.want, s.Contains(tt.args.segment))
		})
	}
}

func TestSegment_Move(t *testing.T) {
	type fields struct {
		row int
		col int
	}
	type args struct {
		direction math.Direction
		dimension math.Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Segment
	}{
		{
			name:   "in dimension range",
			fields: fields{10, 12},
			args:   args{math.Left(), math.Dimension{Height: 20, Width: 30}},
			want:   Segment{10, 11},
		},
		{
			name:   "out of height dimension",
			fields: fields{19, 12},
			args:   args{math.Down(), math.Dimension{Height: 20, Width: 30}},
			want:   Segment{0, 12},
		},
		{
			name:   "out of width dimension",
			fields: fields{19, 29},
			args:   args{math.Right(), math.Dimension{Height: 20, Width: 30}},
			want:   Segment{19, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Segment{
				Row: tt.fields.row,
				Col: tt.fields.col,
			}
			assert.Equal(t, tt.want, s.Move(tt.args.direction, tt.args.dimension))
		})
	}
}

func TestNewHead(t *testing.T) {
	type args struct {
		segment   Segment
		direction math.Direction
	}
	tests := []struct {
		name string
		args args
		want Head
	}{
		{
			name: "new Head",
			args: args{
				segment:   Segment{10, 20},
				direction: math.Left(),
			},
			want: Head{
				Segment:   Segment{10, 20},
				Direction: math.Left(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Head{tt.args.segment, tt.args.direction})
		})
	}
}

func TestHead_SetDirection(t *testing.T) {
	type fields struct {
		segment   Segment
		direction math.Direction
	}
	type args struct {
		direction math.Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Head
	}{
		{
			name:   "set to opposite Direction",
			fields: fields{segment: Segment{1, 2}, direction: math.Left()},
			args:   args{direction: math.Right()},
			want:   Head{Segment{1, 2}, math.Left()},
		},
		{
			name:   "set to non-opposite Direction",
			fields: fields{segment: Segment{1, 2}, direction: math.Left()},
			args:   args{direction: math.Up()},
			want:   Head{Segment{1, 2}, math.Up()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Head{
				Segment:   tt.fields.segment,
				Direction: tt.fields.direction,
			}
			assert.Equal(t, tt.want, h.SetDirection(tt.args.direction))
		})
	}
}

func TestHead_Move(t *testing.T) {
	type fields struct {
		segment   Segment
		direction math.Direction
	}
	type args struct {
		dimension math.Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Head
	}{
		{
			name: "move",
			fields: fields{
				segment:   Segment{1, 2},
				direction: math.Left(),
			},
			args: args{dimension: math.Dimension{Height: 10, Width: 10}},
			want: Head{Segment{1, 1}, math.Left()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Head{
				Segment:   tt.fields.segment,
				Direction: tt.fields.direction,
			}
			assert.Equal(t, tt.want, h.Move(tt.args.dimension))
		})
	}
}

func TestNewTail(t *testing.T) {
	type args struct {
		segments []Segment
	}
	tests := []struct {
		name string
		args args
		want Tail
	}{
		{
			name: "empty arguments",
			args: args{},
			want: Tail{},
		},
		{
			name: "non-empty arguments",
			args: args{[]Segment{{10, 20}, {20, 30}}},
			want: Tail{[]Segment{{10, 20}, {20, 30}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Tail{tt.args.segments})
		})
	}
}

func TestTail_AddFirst(t1 *testing.T) {
	type fields struct {
		segments []Segment
	}
	type args struct {
		segment Segment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Tail
	}{
		{
			name:   "empty Tail",
			fields: fields{},
			args:   args{Segment{10, 20}},
			want:   Tail{[]Segment{{10, 20}}},
		},
		{
			name:   "non-empty Tail",
			fields: fields{[]Segment{{10, 20}}},
			args:   args{Segment{20, 30}},
			want:   Tail{[]Segment{{20, 30}, {10, 20}}},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tail{
				Segments: tt.fields.segments,
			}
			assert.Equal(t1, tt.want, t.AddFirst(tt.args.segment))
		})
	}
}

func TestTail_DropLast(t1 *testing.T) {
	type fields struct {
		segments []Segment
	}
	tests := []struct {
		name   string
		fields fields
		want   Tail
	}{
		{
			name:   "empty Tail",
			fields: fields{},
			want:   Tail{},
		},
		{
			name:   "not-empty Tail",
			fields: fields{[]Segment{{10, 20}, {20, 30}}},
			want:   Tail{[]Segment{{10, 20}}},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tail{
				Segments: tt.fields.segments,
			}
			assert.Equal(t1, tt.want, t.DropLast())
		})
	}
}

func TestTail_IsEmpty(t1 *testing.T) {
	type fields struct {
		segments []Segment
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "empty Tail",
			fields: fields{},
			want:   true,
		},
		{
			name:   "not-empty Tail",
			fields: fields{[]Segment{{10, 20}}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tail{
				Segments: tt.fields.segments,
			}
			assert.Equal(t1, tt.want, t.IsEmpty())
		})
	}
}

func TestTail_Contains(t1 *testing.T) {
	type fields struct {
		segments []Segment
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
			name:   "contains",
			fields: fields{[]Segment{{10, 20}, {11, 20}}},
			args:   args{Segment{11, 20}},
			want:   true,
		},
		{
			name:   "not contains",
			fields: fields{[]Segment{{10, 20}, {11, 20}}},
			args:   args{Segment{12, 20}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tail{
				Segments: tt.fields.segments,
			}
			assert.Equal(t1, tt.want, t.Contains(tt.args.segment))
		})
	}
}
