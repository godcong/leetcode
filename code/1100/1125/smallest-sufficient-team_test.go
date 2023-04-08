package _1125

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_smallestSufficientTeam(t *testing.T) {
	type args struct {
		req_skills []string
		people     [][]string
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
			if got := smallestSufficientTeam(tt.args.req_skills, tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestSufficientTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}
