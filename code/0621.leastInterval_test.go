package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
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
				tasks: strToByteArray("[\"A\",\"A\",\"A\",\"B\",\"B\",\"B\"]"),
				n:     2,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				tasks: strToByteArray("[\"A\",\"A\",\"A\",\"B\",\"B\",\"B\"]"),
				n:     0,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				tasks: strToByteArray("[\"A\",\"A\",\"A\",\"A\",\"A\",\"A\",\"B\",\"C\",\"D\",\"E\",\"F\",\"G\"]"),
				n:     2,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
