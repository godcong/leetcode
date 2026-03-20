package _1376

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numOfMinutes(t *testing.T) {
	type args struct {
		n          int
		headID     int
		manager    []int
		informTime []int
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
			if got := numOfMinutes(tt.args.n, tt.args.headID, tt.args.manager, tt.args.informTime); got != tt.want {
				t.Errorf("numOfMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
