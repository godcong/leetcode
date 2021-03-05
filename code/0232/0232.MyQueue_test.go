package _0232

import (
	"reflect"
	"strconv"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums [][]int
		step []string
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
				nums: strToIntArrArray("[[], [1], [2], [], [], []]"),
				step: strToStringArray("[\"MyQueue\", \"push\", \"push\", \"peek\", \"pop\", \"empty\"]"),
			},
			want: []string{
				"null", "null", "null", "1", "1", "false",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var my MyQueue
			for i, s := range tt.args.step {
				switch s {
				case "MyQueue":
					my = Constructor()
				case "push":
					my.Push(tt.args.nums[i][0])
				case "peek":
					p := my.Peek()
					if strconv.Itoa(p) != tt.want[i] {
						t.Errorf("Peek() = %v, want %v", p, tt.want[i])
					}
				case "pop":
					p := my.Pop()
					if strconv.Itoa(p) != tt.want[i] {
						t.Errorf("Pop() = %v, want %v", p, tt.want[i])
					}
				case "empty":
					e := my.Empty()
					if strconv.FormatBool(e) != tt.want[i] {
						t.Errorf("Empty() = %v, want %v", e, tt.want[i])
					}
				}
			}

			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
