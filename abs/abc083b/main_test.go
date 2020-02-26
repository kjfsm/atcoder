package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		n int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{20, 2, 5},
			want: 84,
		},
		{
			args: args{10, 1, 2},
			want: 13,
		},
		{
			args: args{100, 4, 16},
			want: 4554,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
