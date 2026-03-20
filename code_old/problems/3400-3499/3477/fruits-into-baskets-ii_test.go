package _3477

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numOfUnplacedFruits(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
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
			if got := numOfUnplacedFruits(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("numOfUnplacedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
