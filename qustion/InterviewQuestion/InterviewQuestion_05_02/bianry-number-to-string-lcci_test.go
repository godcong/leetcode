package _InterviewQuestion_05_02

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_printBin(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printBin(tt.args.num); got != tt.want {
				t.Errorf("printBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
