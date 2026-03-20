package _2136

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_earliestFullBloom(t *testing.T) {
	type args struct {
		plantTime []int
		growTime  []int
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
			if got := earliestFullBloom(tt.args.plantTime, tt.args.growTime); got != tt.want {
				t.Errorf("earliestFullBloom() = %v, want %v", got, tt.want)
			}
		})
	}
}
