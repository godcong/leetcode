package code

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
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
				nums:  strToIntArray("[8,2,4,7]"),
				limit: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
