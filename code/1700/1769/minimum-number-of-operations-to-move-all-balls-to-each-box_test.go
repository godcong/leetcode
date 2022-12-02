package _1769

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		boxes string
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
			if got := minOperations(tt.args.boxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
