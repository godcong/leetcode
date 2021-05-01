package _0690

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getImportance(t *testing.T) {
	type args struct {
		employees []*Employee
		id        int
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
				employees: []*Employee{{1, 5, []int{2, 3}}, {2, 3, []int{}}, {3, 3, []int{}}},
				id:        1,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImportance(tt.args.employees, tt.args.id); got != tt.want {
				t.Errorf("getImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}
