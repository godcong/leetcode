package _3531

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countCoveredBuildings(t *testing.T) {
	type args struct {
		n         int
		buildings [][]int
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
			if got := countCoveredBuildings(tt.args.n, tt.args.buildings); got != tt.want {
				t.Errorf("countCoveredBuildings() = %v, want %v", got, tt.want)
			}
		})
	}
}
