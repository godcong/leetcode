package _1606

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_busiestServers(t *testing.T) {
	type args struct {
		k       int
		arrival []int
		load    []int
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
			if got := busiestServers(tt.args.k, tt.args.arrival, tt.args.load); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("busiestServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
