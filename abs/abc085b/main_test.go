package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{10, 8, 8, 6}}, 3},
		{"", args{[]int{15, 15, 15}}, 1},
		{"", args{[]int{50, 30, 50, 100, 50, 80, 30}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.numbers); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
