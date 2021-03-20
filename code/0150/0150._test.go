package _0150

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				tokens: StrToStringArray("[\"2\",\"1\",\"+\",\"3\",\"*\"]"),
			},
			want: 9,
		},
		{
			name: "",
			args: args{
				tokens: StrToStringArray("[\"4\",\"13\",\"5\",\"/\",\"+\"]"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
