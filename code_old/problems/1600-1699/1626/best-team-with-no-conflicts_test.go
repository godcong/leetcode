package _1626

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_bestTeamScore(t *testing.T) {
	type args struct {
		scores []int
		ages   []int
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
			if got := bestTeamScore(tt.args.scores, tt.args.ages); got != tt.want {
				t.Errorf("bestTeamScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
