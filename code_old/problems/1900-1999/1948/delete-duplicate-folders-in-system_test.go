package _1948

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_deleteDuplicateFolder(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicateFolder(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicateFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
