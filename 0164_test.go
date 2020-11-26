package leetcode

import "testing"

func Test_maximumGap(t *testing.T) {
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
			if got := maximumGap(tt.args.nums); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
