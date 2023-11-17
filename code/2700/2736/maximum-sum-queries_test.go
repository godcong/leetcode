package _2736

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumSumQueries(t *testing.T) {
	type args struct {
		nums1   []int
		nums2   []int
		queries [][]int
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
			if got := maximumSumQueries(tt.args.nums1, tt.args.nums2, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumSumQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
