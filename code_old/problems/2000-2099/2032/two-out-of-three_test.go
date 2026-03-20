package _2032

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_twoOutOfThree(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
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
			if got := twoOutOfThree(tt.args.nums1, tt.args.nums2, tt.args.nums3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoOutOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
