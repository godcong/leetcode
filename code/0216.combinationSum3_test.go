package code

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				k: 3,
				n: 7,
			},
			want: [][]int{
				{1, 2, 4},
			},
		},
		{
			name: "",
			args: args{
				k: 3,
				n: 9,
			},
			want: [][]int{
				{1, 2, 6}, {1, 3, 5}, {2, 3, 4},
			},
		},
		{
			name: "",
			args: args{
				k: 4,
				n: 54,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
