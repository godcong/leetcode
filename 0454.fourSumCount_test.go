package leetcode

import "testing"

func Test_fourSumCount(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
		D []int
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
				A: strToIntArray("[ 1, 2]"),
				B: strToIntArray("[-2,-1]"),
				C: strToIntArray("[-1, 2]"),
				D: strToIntArray("[ 0, 2]"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.A, tt.args.B, tt.args.C, tt.args.D); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
