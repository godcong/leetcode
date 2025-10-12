package _3539

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_magicalSum(t *testing.T) {
	type args struct {
		m    int
		k    int
		nums []int
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
			if got := magicalSum(tt.args.m, tt.args.k, tt.args.nums); got != tt.want {
				t.Errorf("magicalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
