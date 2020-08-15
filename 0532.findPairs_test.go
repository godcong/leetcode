package leetcode

import "testing"

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{
					3, 1, 4, 1, 5,
				},
				k: 2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{
					1, 1, 1, 2, 1,
				},
				k: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{
					1, 2, 3, 4, 5,
				},
				k: 1,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{
					1, 3, 1, 5, 4,
				},
				k: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
