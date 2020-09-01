package leetcode

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "1 ",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: ".2e81",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "+100",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "5e2",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "-123",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "3.1416",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "-1E-16",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "0123",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "12e",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "1a3.14",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "1.2.3",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "+-5",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "12e+5.4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
