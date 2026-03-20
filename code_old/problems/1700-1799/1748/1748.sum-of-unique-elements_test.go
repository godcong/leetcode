package _1748

import (
	"testing"
)

func Test_sumOfUnique(t *testing.T) {
	type args struct {
		nums []int
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
			if got := sumOfUnique(tt.args.nums); got != tt.want {
				t.Errorf("sumOfUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
