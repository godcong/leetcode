package code

import (
	"reflect"
	"testing"
)

func Test_findCriticalAndPseudoCriticalEdges(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				n:     5,
				edges: strToIntArrArray("[[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]"),
			},
			want: strToIntArrArray("[[0,1],[2,3,4,5]]"),
		},
		{
			name: "",
			args: args{
				n:     4,
				edges: strToIntArrArray("[[0,1,1],[1,2,1],[2,3,1],[0,3,1]]"),
			},
			want: strToIntArrArray("[[],[0,1,2,3]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCriticalAndPseudoCriticalEdges(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCriticalAndPseudoCriticalEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
