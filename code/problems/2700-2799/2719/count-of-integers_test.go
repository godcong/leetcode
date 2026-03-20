package _2719

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_count(t *testing.T) {
	type args struct {
		num1    string
		num2    string
		min_sum int
		max_sum int
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
			if got := count(tt.args.num1, tt.args.num2, tt.args.min_sum, tt.args.max_sum); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
