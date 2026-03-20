package _3607

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_processQueries(t *testing.T) {
	type args struct {
		c           int
		connections [][]int
		queries     [][]int
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
			if got := processQueries(tt.args.c, tt.args.connections, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
