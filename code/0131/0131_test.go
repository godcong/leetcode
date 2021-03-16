package _0131

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
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
				s: "aab",
			},
			want: StrToStringArrayArray("[[\"a\",\"a\",\"b\"],[\"aa\",\"b\"]]"),
		},
		{
			name: "",
			args: args{
				s: "a",
			},
			want: StrToStringArrayArray("[[\"a\"]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
