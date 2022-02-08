package old

import "testing"

func Test_judgePoint24(t *testing.T) {
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
				nums: []int{
					4, 1, 8, 7,
				},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{
					1, 2, 1, 2,
				},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				nums: []int{
					1, 3, 4, 6,
				},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{
					3, 3, 8, 8,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgePoint24(tt.args.nums); got != tt.want {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}
