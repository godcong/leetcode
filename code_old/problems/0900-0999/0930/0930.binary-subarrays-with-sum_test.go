package _0930

import (
	"testing"
)

func Test_numSubarraysWithSum(t *testing.T) {
	type args struct {
		nums []int
		goal int
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
			if got := numSubarraysWithSum(tt.args.nums, tt.args.goal); got != tt.want {
				t.Errorf("numSubarraysWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
