package code

import "testing"

func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
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
				points: strToIntArrArray("[[3,12],[-2,5],[-4,1]]"),
			},
			want: 18,
		},
		{
			name: "",
			args: args{
				points: strToIntArrArray("[[0,0],[1,1],[1,0],[-1,1]]"),
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				points: strToIntArrArray("[[-1000000,-1000000],[1000000,1000000]]"),
			},
			want: 4000000,
		},
		{
			name: "",
			args: args{
				points: strToIntArrArray("[[0,0]]"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
