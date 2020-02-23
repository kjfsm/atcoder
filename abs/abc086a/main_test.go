package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				a: 3,
				b: 4,
			},
			want: "Even",
		},
		{
			args: args{
				a: 1,
				b: 21,
			},
			want: "Odd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
