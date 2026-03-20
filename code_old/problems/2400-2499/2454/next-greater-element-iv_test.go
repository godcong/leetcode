package _2454

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_secondGreaterElement(t *testing.T) {
	type args struct {
		nums []int
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
			if got := secondGreaterElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("secondGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
