package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "aa",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "au",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "aab",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
