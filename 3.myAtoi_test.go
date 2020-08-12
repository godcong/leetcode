package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				str: "42",
			},
			want: 42,
		},
		{
			name: "",
			args: args{
				str: "   -42",
			},
			want: -42,
		},
		{
			name: "",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "",
			args: args{
				str: "+-123",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				str: "+0 456",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				str: "9223372036854775808",
			},
			want: 2147483647,
		},
		{
			name: "",
			args: args{
				str: "  0000000000012345678",
			},
			want: 12345678,
		},
		{
			name: "",
			args: args{
				str: "-   234",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				str: "0-1",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				str: "      -11919730356x",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
