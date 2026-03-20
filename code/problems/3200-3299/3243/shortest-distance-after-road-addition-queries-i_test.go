package _3243

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_shortestDistanceAfterQueries(t *testing.T) {
	type args struct {
		n       int
		queries [][]int
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
			if got := shortestDistanceAfterQueries(tt.args.n, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestDistanceAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
