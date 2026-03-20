package _2612

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minReverseOperations(t *testing.T) {
	type args struct {
		n      int
		p      int
		banned []int
		k      int
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
			if got := minReverseOperations(tt.args.n, tt.args.p, tt.args.banned, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minReverseOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
