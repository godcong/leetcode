package main

import (
	"reflect"
	"testing"
)

func Test_pointArray(t *testing.T) {
	tests := []struct {
		name string
		want []*Val
	}{
		// TODO: Add test cases.
		{
			name: "",
			want: []*Val{
				{
					Val: "1",
				},
				{
					Val: "2",
				},
				{
					Val: "3",
				},
				{
					Val: "4",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := pointArray(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("pointArray() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
