package logics

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewHead(t *testing.T) {
	type args struct {
		segment   Segment
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want Head
	}{
		{
			name: "new head",
			args: args{
				segment:   NewSegment(10, 20),
				direction: Left(),
			},
			want: Head{
				segment:   NewSegment(10, 20),
				direction: Left(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewHead(tt.args.segment, tt.args.direction))
		})
	}
}

func TestHead_SetDirection(t *testing.T) {
	type fields struct {
		segment   Segment
		direction Direction
	}
	type args struct {
		direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Head
	}{
		{
			name:   "set to opposite direction",
			fields: fields{segment: NewSegment(1, 2), direction: Left()},
			args:   args{direction: Right()},
			want:   NewHead(NewSegment(1, 2), Left()),
		},
		{
			name:   "set to non-opposite direction",
			fields: fields{segment: NewSegment(1, 2), direction: Left()},
			args:   args{direction: Up()},
			want:   NewHead(NewSegment(1, 2), Up()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Head{
				segment:   tt.fields.segment,
				direction: tt.fields.direction,
			}
			assert.Equal(t, tt.want, h.SetDirection(tt.args.direction))
		})
	}
}

func TestHead_Move(t *testing.T) {
	type fields struct {
		segment   Segment
		direction Direction
	}
	type args struct {
		dimension Dimension
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Head
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Head{
				segment:   tt.fields.segment,
				direction: tt.fields.direction,
			}
			if got := h.Move(tt.args.dimension); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHead_Move1(t *testing.T) {
	type fields struct {
		segment   Segment
		direction Direction
	}
	type args struct {
		dimension Dimension
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
				segment:   NewSegment(1, 2),
				direction: Left(),
			},
			args: args{dimension: NewDimension(10, 10)},
			want: NewHead(NewSegment(1, 1), Left()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Head{
				segment:   tt.fields.segment,
				direction: tt.fields.direction,
			}
			assert.Equal(t, tt.want, h.Move(tt.args.dimension))
		})
	}
}
