package _3433

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countMentions(t *testing.T) {
	type args struct {
		numberOfUsers int
		events        [][]string
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
			if got := countMentions(tt.args.numberOfUsers, tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countMentions() = %v, want %v", got, tt.want)
			}
		})
	}
}
