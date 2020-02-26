package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		n int
		y int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{9, 45000}, true},
		{"", args{20, 196000}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Calc(tt.args.n, tt.args.y)
			if tt.want {
				if 10000*got+5000*got1+1000*got2 != tt.args.y {
					t.Errorf("Calc() got = (%v, %v, %v ) want %v", got, got1, got2, tt.args.y)
				}
			} else {
				if got != -1 || got1 != -1 || got2 != -1 {
					t.Errorf("Calc() got = (%v, %v, %v ) want %v", got, got1, got2, -1)
				}
			}
		})
	}
}
