package _2446

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_haveConflict(t *testing.T) {
	type args struct {
		event1 []string
		event2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := haveConflict(tt.args.event1, tt.args.event2); got != tt.want {
				t.Errorf("haveConflict() = %v, want %v", got, tt.want)
			}
		})
	}
}
