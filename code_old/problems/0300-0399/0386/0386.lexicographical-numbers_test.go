package _0386

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_lexicalOrder(t *testing.T) {
	type args struct {
		n int
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
			if got := lexicalOrder(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
