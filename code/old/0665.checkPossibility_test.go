package old

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				nums: []int{4, 2, 3},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{4, 2, 1},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				nums: []int{3, 4, 2, 3},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				nums: []int{3, 2, 2, 3, 5},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 3, 2},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{
					5, 7, 1, 8,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
