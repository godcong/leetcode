package _2569

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_handleQuery(t *testing.T) {
	type args struct {
		nums1   []int
		nums2   []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleQuery(tt.args.nums1, tt.args.nums2, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
