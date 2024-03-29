package old

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	type args struct {
		A [][]int
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
				A: strToIntArrArray("[[1,1,0],[1,0,1],[0,0,0]]"),
			},
			want: strToIntArrArray("[[1,0,0],[0,1,0],[1,1,1]]"),
		},
		{
			name: "",
			args: args{
				A: strToIntArrArray("[[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]"),
			},
			want: strToIntArrArray("[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
