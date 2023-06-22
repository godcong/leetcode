package _InterviewQuestion_16_19

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_pondSizes(t *testing.T) {
	type args struct {
		land [][]int
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
			if got := pondSizes(tt.args.land); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pondSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
