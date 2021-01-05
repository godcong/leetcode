package code

import "testing"

func Test_threeConsecutiveOdds(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				arr: []int{
					2, 6, 4, 1,
				},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				arr: []int{
					1, 2, 34, 3, 4, 5, 7, 23, 12,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeConsecutiveOdds(tt.args.arr); got != tt.want {
				t.Errorf("threeConsecutiveOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
