package _2037

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minMovesToSeat(t *testing.T) {
	type args struct {
		seats    []int
		students []int
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
			if got := minMovesToSeat(tt.args.seats, tt.args.students); got != tt.want {
				t.Errorf("minMovesToSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
