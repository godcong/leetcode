package _0911

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		persons []int
		times   []int
	}
	tests := []struct {
		name string
		args args
		want TopVotedCandidate
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.persons, tt.args.times); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
