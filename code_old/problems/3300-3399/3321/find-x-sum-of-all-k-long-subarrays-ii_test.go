package _3321

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findXSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
		x    int
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
			if got := findXSum(tt.args.nums, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findXSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
