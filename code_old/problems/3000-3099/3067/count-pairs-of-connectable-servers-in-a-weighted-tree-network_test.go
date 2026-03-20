package _3067

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countPairsOfConnectableServers(t *testing.T) {
	type args struct {
		edges       [][]int
		signalSpeed int
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
			if got := countPairsOfConnectableServers(tt.args.edges, tt.args.signalSpeed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPairsOfConnectableServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
