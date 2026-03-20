package _2391

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_garbageCollection(t *testing.T) {
	type args struct {
		garbage []string
		travel  []int
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
			if got := garbageCollection(tt.args.garbage, tt.args.travel); got != tt.want {
				t.Errorf("garbageCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
