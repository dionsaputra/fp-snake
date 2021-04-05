package snake

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			args: args{[]Segment{NewSegment(10, 20), NewSegment(20, 30)}},
			want: Tail{[]Segment{NewSegment(10, 20), NewSegment(20, 30)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewTail(tt.args.segments...))
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
			name:   "empty tail",
			fields: fields{},
			args:   args{NewSegment(10, 20)},
			want:   NewTail(NewSegment(10, 20)),
		},
		{
			name:   "non-empty tail",
			fields: fields{[]Segment{NewSegment(10, 20)}},
			args:   args{NewSegment(20, 30)},
			want:   NewTail(NewSegment(20, 30), NewSegment(10, 20)),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tail{
				segments: tt.fields.segments,
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
			name:   "empty tail",
			fields: fields{},
			want:   NewTail(),
		},
		{
			name:   "not-empty tail",
			fields: fields{[]Segment{NewSegment(10, 20), NewSegment(20, 30)}},
			want:   NewTail(NewSegment(10, 20)),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tail{
				segments: tt.fields.segments,
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
			name:   "empty tail",
			fields: fields{},
			want:   true,
		},
		{
			name:   "not-empty tail",
			fields: fields{[]Segment{NewSegment(10, 20)}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tail{
				segments: tt.fields.segments,
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
			fields: fields{[]Segment{NewSegment(10, 20), NewSegment(11, 20)}},
			args:   args{NewSegment(11, 20)},
			want:   true,
		},
		{
			name:   "not contains",
			fields: fields{[]Segment{NewSegment(10, 20), NewSegment(11, 20)}},
			args:   args{NewSegment(12, 20)},
			want:   false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tail{
				segments: tt.fields.segments,
			}
			assert.Equal(t1, tt.want, t.Contains(tt.args.segment))
		})
	}
}
