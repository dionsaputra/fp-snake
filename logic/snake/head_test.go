package snake

import (
	"github.com/dionsaputra/fp-snake/logic/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHead(t *testing.T) {
	type args struct {
		segment   Segment
		direction geometry.Direction
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
				direction: geometry.Left(),
			},
			want: Head{
				segment:   NewSegment(10, 20),
				direction: geometry.Left(),
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
		direction geometry.Direction
	}
	type args struct {
		direction geometry.Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Head
	}{
		{
			name:   "set to opposite direction",
			fields: fields{segment: NewSegment(1, 2), direction: geometry.Left()},
			args:   args{direction: geometry.Right()},
			want:   NewHead(NewSegment(1, 2), geometry.Left()),
		},
		{
			name:   "set to non-opposite direction",
			fields: fields{segment: NewSegment(1, 2), direction: geometry.Left()},
			args:   args{direction: geometry.Up()},
			want:   NewHead(NewSegment(1, 2), geometry.Up()),
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
		direction geometry.Direction
	}
	type args struct {
		dimension geometry.Dimension
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
				direction: geometry.Left(),
			},
			args: args{dimension: geometry.NewDimension(10, 10)},
			want: NewHead(NewSegment(1, 1), geometry.Left()),
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
