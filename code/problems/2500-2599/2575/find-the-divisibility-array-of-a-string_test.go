package _2575

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_divisibilityArray(t *testing.T) {
	type args struct {
		word string
		m    int
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
			if got := divisibilityArray(tt.args.word, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divisibilityArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
