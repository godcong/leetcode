package old

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
		{
			name: "",
			args: args{
				nums:  strToIntArray("[10,1,2,4,7,2]"),
				limit: 5,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums:  strToIntArray("[4,2,2,2,4,4,2,2]"),
				limit: 0,
			},
			want: 3,
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
