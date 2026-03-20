package _2099

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSubsequence(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubsequence(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
