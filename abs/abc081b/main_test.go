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
		{
			args: args{
				numbers: []int{8, 12, 40},
			},
			want: 2,
		},
		{
			args: args{
				numbers: []int{5, 6, 8, 10},
			},
			want: 0,
		},
		{
			args: args{
				numbers: []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.numbers); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
