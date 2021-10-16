package main

import "testing"

func Test_addIntegers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero", args{[]int{0}}, 0},
		{"two positive integers", args{[]int{1, 2}}, 3},
		{"three negative integers", args{[]int{-10, -20, -30}}, -60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addIntegers(tt.args.nums...); got != tt.want {
				t.Errorf("addIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
