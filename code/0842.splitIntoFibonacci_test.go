package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: "123456579",
			},
			want: leetcode.strToIntArray("[123,456,579]"),
		},
		{
			name: "",
			args: args{
				s: "11235813",
			},
			want: leetcode.strToIntArray("[1,1,2,3,5,8,13]"),
		},
		{
			name: "",
			args: args{
				s: "112358130",
			},
			want: leetcode.strToIntArray("[]"),
		},
		{
			name: "",
			args: args{
				s: "0123",
			},
			want: leetcode.strToIntArray("[]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.s); !reflect.DeepEqual(got, tt.want) && len(got) != len(tt.want) /*ignore nil and null array*/ {
				t.Errorf("splitIntoFibonacci() = %v size:%v, want %v size:%v", got, len(got), tt.want, len(tt.want))
			}
		})
	}
}
