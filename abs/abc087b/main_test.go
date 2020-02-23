package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{2, 2, 2, 100},
			want: 2,
		},
		{
			args: args{5, 1, 0, 150},
			want: 0,
		},
		{
			args: args{30, 40, 50, 6000},
			want: 213,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.a, tt.args.b, tt.args.c, tt.args.x); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
