package _InterviewQuestion_17_19

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_missingTwo(t *testing.T) {
	type args struct {
		nums []int
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
			if got := missingTwo(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
