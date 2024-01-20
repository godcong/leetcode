package _2788

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_splitWordsBySeparator(t *testing.T) {
	type args struct {
		words     []string
		separator byte
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
			if got := splitWordsBySeparator(tt.args.words, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitWordsBySeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
