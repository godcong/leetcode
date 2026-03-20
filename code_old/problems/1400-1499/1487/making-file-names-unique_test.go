package _1487

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getFolderNames(t *testing.T) {
	type args struct {
		names []string
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
			if got := getFolderNames(tt.args.names); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFolderNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
