package leetcode

import "testing"

func Test_minCount(t *testing.T) {
	type args struct {
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				coins: []int{4, 2, 1},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				coins: []int{2, 3, 10},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCount(tt.args.coins); got != tt.want {
				t.Errorf("minCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
