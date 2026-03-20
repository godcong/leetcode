package _2558

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_pickGifts(t *testing.T) {
	type args struct {
		gifts []int
		k     int
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
			if got := pickGifts(tt.args.gifts, tt.args.k); got != tt.want {
				t.Errorf("pickGifts() = %v, want %v", got, tt.want)
			}
		})
	}
}
