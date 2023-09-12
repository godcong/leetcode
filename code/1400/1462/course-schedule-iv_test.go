package _1462

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkIfPrerequisite(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPrerequisite(tt.args.numCourses, tt.args.prerequisites, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkIfPrerequisite() = %v, want %v", got, tt.want)
			}
		})
	}
}
