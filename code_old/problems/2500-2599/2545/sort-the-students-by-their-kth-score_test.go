package _2545

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sortTheStudents(t *testing.T) {
	type args struct {
		score [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortTheStudents(tt.args.score, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTheStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
