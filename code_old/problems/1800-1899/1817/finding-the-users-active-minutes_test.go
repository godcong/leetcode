package _1817

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findingUsersActiveMinutes(t *testing.T) {
	type args struct {
		logs [][]int
		k    int
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
			if got := findingUsersActiveMinutes(tt.args.logs, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findingUsersActiveMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
