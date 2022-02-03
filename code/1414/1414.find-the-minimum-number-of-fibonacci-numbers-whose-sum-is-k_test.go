package _1414

import (
	"testing"
)

func Test_findMinFibonacciNumbers(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinFibonacciNumbers(tt.args.k); got != tt.want {
				t.Errorf("findMinFibonacciNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
