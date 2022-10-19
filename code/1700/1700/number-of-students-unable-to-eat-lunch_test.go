package _1700

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countStudents(t *testing.T) {
	type args struct {
		students   []int
		sandwiches []int
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
			if got := countStudents(tt.args.students, tt.args.sandwiches); got != tt.want {
				t.Errorf("countStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
