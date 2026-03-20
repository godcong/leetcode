package _2169

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countOperations(t *testing.T) {
	type args struct {
		num1 int
		num2 int
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
			if got := countOperations(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("countOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
