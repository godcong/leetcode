package _2483

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_bestClosingTime(t *testing.T) {
	type args struct {
		customers string
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
			if got := bestClosingTime(tt.args.customers); got != tt.want {
				t.Errorf("bestClosingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
