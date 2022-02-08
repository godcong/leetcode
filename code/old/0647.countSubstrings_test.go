package old

import "testing"

func Test_countSubstrings(t *testing.T) {
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
				s: "abc",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				s: "abaa",
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				s: "abcdcbaaa",
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
