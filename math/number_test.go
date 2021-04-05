package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMod(t *testing.T) {
	type args struct {
		num int
		div int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive number",
			args: args{
				num: 10,
				div: 3,
			},
			want: 1,
		},
		{
			name: "negative number",
			args: args{
				num: -10,
				div: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Mod(tt.args.num, tt.args.div))
		})
	}
}
