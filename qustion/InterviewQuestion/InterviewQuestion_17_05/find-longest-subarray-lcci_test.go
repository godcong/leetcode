package _InterviewQuestion_17_05

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findLongestSubarray(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestSubarray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLongestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
