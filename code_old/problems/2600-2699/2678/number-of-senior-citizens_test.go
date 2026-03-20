package _2678

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countSeniors(t *testing.T) {
	type args struct {
		details []string
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
			if got := countSeniors(tt.args.details); got != tt.want {
				t.Errorf("countSeniors() = %v, want %v", got, tt.want)
			}
		})
	}
}
