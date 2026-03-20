package _3272

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countGoodIntegers(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodIntegers(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("countGoodIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
