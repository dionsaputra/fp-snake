package fp

import (
	"reflect"
	"strconv"
	"testing"
)

func TestRangeOf(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want Range
	}{
		{
			name: "when empty args got fp 0 to 0",
			args: args{args: []int{}},
			want: Range{0, 0},
		},
		{
			name: "when args size is 1 got fp 0 to args[0]",
			args: args{args: []int{5}},
			want: Range{0, 5},
		},
		{
			name: "when args size is 2 got fp args[0] to args[1]",
			args: args{args: []int{2, 5}},
			want: Range{2, 5},
		},
		{
			name: "when args size is more than 2 got fp args[0] to args[1]",
			args: args{args: []int{2, 5, 7}},
			want: Range{2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeOf(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_ForEach(t *testing.T) {
	type fields struct {
		start int
		end   int
	}
	type args struct {
		action func(index int)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "action print",
			fields: fields{1, 5},
			args: args{action: func(index int) {
				print(index)
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Range{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			r.Forward(tt.args.action)
		})
	}
}

func TestRange_ReduceToString(t *testing.T) {
	type fields struct {
		start int
		end   int
	}
	type args struct {
		action func(index int) string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "reduce fp 1 to 5 got 1234",
			fields: fields{1, 5},
			args: args{
				action: func(index int) string {
					return strconv.Itoa(index)
				},
			},
			want: "1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Range{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := r.reduceToString(tt.args.action); got != tt.want {
				t.Errorf("reduceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_JoinToString(t *testing.T) {
	type fields struct {
		start int
		end   int
	}
	type args struct {
		separator string
		transform func(index int) string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "join fp 1 to 5 separator comma got 1,2,3,4",
			fields: fields{1, 5},
			args: args{
				separator: ",",
				transform: func(index int) string {
					return strconv.Itoa(index)
				},
			},
			want: "1,2,3,4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Range{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := r.JoinToString(tt.args.separator, tt.args.transform); got != tt.want {
				t.Errorf("JoinToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_Iterate(t *testing.T) {
	type fields struct {
		start int
		end   int
	}
	type args struct {
		stop func(index int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "stop at middle",
			fields: fields{
				start: 0,
				end:   10,
			},
			args: args{
				stop: func(index int) bool {
					return index == 5
				},
			},
			want: 5,
		},

		{
			name: "stop at end",
			fields: fields{
				start: 0,
				end:   10,
			},
			args: args{
				stop: func(index int) bool {
					return index == 12
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Range{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := r.Iterate(tt.args.stop); got != tt.want {
				t.Errorf("Iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_Contains(t *testing.T) {
	type fields struct {
		start int
		end   int
	}
	type args struct {
		predicate func(index int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "contains",
			fields: fields{
				start: 0,
				end:   10,
			},
			args: args{
				predicate: func(index int) bool {
					return index == 5
				},
			},
			want: true,
		},
		{
			name: "not contains",
			fields: fields{
				start: 0,
				end:   10,
			},
			args: args{
				predicate: func(index int) bool {
					return index == 20
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Range{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := r.Contains(tt.args.predicate); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
