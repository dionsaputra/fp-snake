package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDimension(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	tests := []struct {
		name string
		args args
		want Dimension
	}{
		{
			name: "new dimension",
			args: args{20, 30},
			want: Dimension{20, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewDimension(tt.args.height, tt.args.width))
		})
	}
}
