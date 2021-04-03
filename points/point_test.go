package points

import (
	"reflect"
	"testing"
)

func TestPointOf(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want Point
	}{
		{
			name: "PointOf",
			args: args{
				x: 2,
				y: 3,
			},
			want: Point{
				x: 2,
				y: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointOf(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Translate(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		dx int
		dy int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Point
	}{
		{
			name: "translate",
			fields: fields{
				x: 2,
				y: 3,
			},
			args: args{
				dx: 1,
				dy: 2,
			},
			want: PointOf(3, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.Translate(tt.args.dx, tt.args.dy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
