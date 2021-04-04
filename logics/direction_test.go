package logics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDirection(t *testing.T) {
	type args struct {
		deltaRow int
		deltaCol int
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{
			name: "new direction",
			args: args{1, 0},
			want: Direction{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewDirection(tt.args.deltaRow, tt.args.deltaCol))
		})
	}
}
