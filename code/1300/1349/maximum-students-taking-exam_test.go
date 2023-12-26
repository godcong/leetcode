package _1349

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxStudents(t *testing.T) {
	type args struct {
		seats [][]byte
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
			if got := maxStudents(tt.args.seats); got != tt.want {
				t.Errorf("maxStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
