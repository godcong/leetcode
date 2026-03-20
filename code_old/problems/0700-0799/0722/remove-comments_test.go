package _0722

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_removeComments(t *testing.T) {
	type args struct {
		source []string
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
			if got := removeComments(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
