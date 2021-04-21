package _0091

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
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
				s: "12",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "226",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := numDecodings(tt.args.s); got != tt.want {
					t.Errorf("numDecodings() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
