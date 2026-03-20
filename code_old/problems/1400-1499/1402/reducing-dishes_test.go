package _1402

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSatisfaction(t *testing.T) {
	type args struct {
		satisfaction []int
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
			if got := maxSatisfaction(tt.args.satisfaction); got != tt.want {
				t.Errorf("maxSatisfaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
