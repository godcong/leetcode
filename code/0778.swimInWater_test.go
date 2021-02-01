package code

import "testing"

func Test_swimInWater(t *testing.T) {
	type args struct {
		grid [][]int
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
				grid: strToIntArrArray("[[0,2],[1,3]]"),
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				grid: strToIntArrArray("[[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]"),
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swimInWater(tt.args.grid); got != tt.want {
				t.Errorf("swimInWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
