package old

import "testing"

func Test_makeGood(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			name: "",
			args: args{
				s: "abBAcC",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "s",
			},
			want: "s",
		},
		{
			name: "",
			args: args{
				s: "Pp",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.args.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
