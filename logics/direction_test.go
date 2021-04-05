package logics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirection_IsOpposite(t *testing.T) {
	type fields struct {
		vertical   int
		horizontal int
	}
	type args struct {
		direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "opposite",
			fields: fields{
				vertical:   1,
				horizontal: 0,
			},
			args: args{direction: Up()},
			want: true,
		},
		{
			name: "not opposite",
			fields: fields{
				vertical:   1,
				horizontal: 0,
			},
			args: args{direction: Right()},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Direction{
				vertical:   tt.fields.vertical,
				horizontal: tt.fields.horizontal,
			}
			assert.Equal(t, tt.want, d.IsOpposite(tt.args.direction))
		})
	}
}