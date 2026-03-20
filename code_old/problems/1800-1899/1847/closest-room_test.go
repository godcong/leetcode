package _1847

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_closestRoom(t *testing.T) {
	type args struct {
		rooms   [][]int
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
			if got := closestRoom(tt.args.rooms, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}
