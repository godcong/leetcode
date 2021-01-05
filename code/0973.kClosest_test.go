package code

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		K      int
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
				points: strToIntArrArray("[[1,3],[-2,2]]"),
				K:      1,
			},
			want: strToIntArrArray("[[-2,2]]"),
		},
		{
			name: "",
			args: args{
				points: strToIntArrArray("[[3,3],[5,-1],[-2,4]]"),
				K:      2,
			},
			want: strToIntArrArray("[[3,3],[-2,4]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
