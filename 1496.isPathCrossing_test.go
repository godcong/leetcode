package leetcode

import "testing"

func Test_isPathCrossing(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				path: "NNSWWEWSSESSWENNW",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPathCrossing(tt.args.path); got != tt.want {
				t.Errorf("isPathCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
