package leetcode

import "testing"

func Test_largestTimeFromDigits(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				A: []int{
					1, 2, 3, 4,
				},
			},
			want: "23:41",
		},
		{
			name: "",
			args: args{
				A: []int{
					4, 2, 3, 1,
				},
			},
			want: "23:41",
		},
		{
			name: "",
			args: args{
				A: []int{
					0, 0, 0, 0,
				},
			},
			want: "00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTimeFromDigits(tt.args.A); got != tt.want {
				t.Errorf("largestTimeFromDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
