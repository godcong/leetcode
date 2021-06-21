package _0401

import (
	"reflect"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	type args struct {
		turnedOn int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatch(tt.args.turnedOn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBinaryWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
