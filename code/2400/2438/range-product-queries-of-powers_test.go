package _2438

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_productQueries(t *testing.T) {
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
			if got := productQueries(tt.args.n, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
