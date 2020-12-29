package leetcode

import "testing"

func Test_minPatches(t *testing.T) {
	type args struct {
		nums []int
		n    int
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
				nums: nil,
				n:    0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPatches(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("minPatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
