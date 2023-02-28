package _2363

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_mergeSimilarItems(t *testing.T) {
	type args struct {
		items1 [][]int
		items2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSimilarItems(tt.args.items1, tt.args.items2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSimilarItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
