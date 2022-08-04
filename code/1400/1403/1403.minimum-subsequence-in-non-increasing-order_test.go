package _1403

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minSubsequence(t *testing.T) {
	type args struct {
		nums []int
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
			if got := minSubsequence(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
