package leetcode

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				A: []string{
					"bella", "label", "roller",
				},
			},
			want: []string{
				"e", "l", "l",
			},
		},
		{
			name: "",
			args: args{
				A: []string{
					"cool", "lock", "cook",
				},
			},
			want: []string{
				"c", "o",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
