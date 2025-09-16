package _2197

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_replaceNonCoprimes(t *testing.T) {
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
			if got := replaceNonCoprimes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceNonCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
