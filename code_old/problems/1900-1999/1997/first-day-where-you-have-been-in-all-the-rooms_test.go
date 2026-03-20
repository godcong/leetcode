package _1997

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_firstDayBeenInAllRooms(t *testing.T) {
	type args struct {
		nextVisit []int
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
			if got := firstDayBeenInAllRooms(tt.args.nextVisit); got != tt.want {
				t.Errorf("firstDayBeenInAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
