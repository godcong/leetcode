package _0937

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_reorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
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
			if got := reorderLogFiles(tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderLogFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
