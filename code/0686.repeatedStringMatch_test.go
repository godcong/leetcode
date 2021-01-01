package code

import "testing"

func Test_repeatedStringMatch(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				A: "abcd",
				B: "cdabcdab",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				A: "a",
				B: "a",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				A: "aa",
				B: "a",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				A: "abcd",
				B: "abcd",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				A: "abcd",
				B: "cdabcdabcdabcdab",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
