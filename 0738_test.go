package leetcode

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
	type args struct {
		N int
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
				N: 10,
			},
			want: 9,
		},
		{
			name: "",
			args: args{
				N: 1234,
			},
			want: 1234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.N); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
