package _2409

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countDaysTogether(t *testing.T) {
	type args struct {
		arriveAlice string
		leaveAlice  string
		arriveBob   string
		leaveBob    string
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
			if got := countDaysTogether(tt.args.arriveAlice, tt.args.leaveAlice, tt.args.arriveBob, tt.args.leaveBob); got != tt.want {
				t.Errorf("countDaysTogether() = %v, want %v", got, tt.want)
			}
		})
	}
}
