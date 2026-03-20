package _2200

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findKDistantIndices(t *testing.T) {
	type args struct {
		nums []int
		key  int
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
			if got := findKDistantIndices(tt.args.nums, tt.args.key, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKDistantIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
