package logics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSegment(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want Segment
	}{
		{
			name: "new segment",
			args: args{10, 20},
			want: Segment{10, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewSegment(tt.args.row, tt.args.col))
		})
	}
}

func TestSegment_Move(t *testing.T) {
	type fields struct {
		row int
		col int
	}
	type args struct {
		direction Direction
		dimension Dimension
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
			args:   args{Left(), NewDimension(20, 30)},
			want:   NewSegment(10, 11),
		},
		{
			name:   "out of height dimension",
			fields: fields{19, 12},
			args:   args{Down(), NewDimension(20, 30)},
			want:   NewSegment(0, 12),
		},
		{
			name:   "out of width dimension",
			fields: fields{19, 29},
			args:   args{Right(), NewDimension(20, 30)},
			want:   NewSegment(19, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Segment{
				row: tt.fields.row,
				col: tt.fields.col,
			}
			assert.Equal(t, tt.want, s.Move(tt.args.direction, tt.args.dimension))
		})
	}
}
