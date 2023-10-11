package _2512

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_topStudents(t *testing.T) {
	type args struct {
		positive_feedback []string
		negative_feedback []string
		report            []string
		student_id        []int
		k                 int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topStudents(tt.args.positive_feedback, tt.args.negative_feedback, tt.args.report, tt.args.student_id, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
