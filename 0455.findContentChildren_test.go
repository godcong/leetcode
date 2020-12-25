package leetcode

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
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
				g: strToIntArray("[1,2,3]"),
				s: strToIntArray("[1,1]"),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				g: strToIntArray("[1,2]"),
				s: strToIntArray("[1,2,3]"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
