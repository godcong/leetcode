package _2956

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findIntersectionValues(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
			if got := findIntersectionValues(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIntersectionValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
