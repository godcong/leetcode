package code

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
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
				s: "Hello World",
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				s: "Hello World ",
			},
			want: 5,
		},
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
			want: 0,
		},
		{
			name: "",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "a ",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: " a",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: " a ",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "b   a    ",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
