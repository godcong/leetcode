package _2410

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_matchPlayersAndTrainers(t *testing.T) {
	type args struct {
		players  []int
		trainers []int
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
			if got := matchPlayersAndTrainers(tt.args.players, tt.args.trainers); got != tt.want {
				t.Errorf("matchPlayersAndTrainers() = %v, want %v", got, tt.want)
			}
		})
	}
}
