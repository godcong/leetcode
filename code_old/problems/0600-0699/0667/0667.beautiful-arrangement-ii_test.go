package _0667

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_constructArray(t *testing.T) {
	type args struct {
		n int
		k int
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
			if got := constructArray(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
