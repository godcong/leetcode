package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				strs: strToStringArray("[\"eat\", \"tea\", \"tan\", \"ate\", \"nat\", \"bat\"]"),
			},
			want: strToStringArrayArray("[\n  [\"ate\",\"eat\",\"tea\"],\n  [\"nat\",\"tan\"],\n  [\"bat\"]\n]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			for i := range tt.want {
				sort.Strings(tt.want[i])
				sort.Strings(got[i])
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
