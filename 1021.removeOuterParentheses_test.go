package leetcode

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				S: "(()())(())",
			},
			want: "()()()",
		},
		{
			name: "",
			args: args{
				S: "(()())(())(()(()))",
			},
			want: "()()()()(())",
		},
		{
			name: "",
			args: args{
				S: "()()",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.S); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
