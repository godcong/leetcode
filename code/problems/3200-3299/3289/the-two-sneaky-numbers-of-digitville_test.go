package _3289

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getSneakyNumbers(t *testing.T) {
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
			if got := getSneakyNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSneakyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
