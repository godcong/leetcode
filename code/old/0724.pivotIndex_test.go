package old

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
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
				nums: strToIntArray("[1, 7, 3, 6, 5, 6]"),
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[1, 2, 3]"),
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[2, 1, -1]"),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: strToIntArray("[0, 0, 0, 0, 1]"),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
