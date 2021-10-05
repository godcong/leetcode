package _0284

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		iter *Iterator
	}
	tests := []struct {
		name string
		args args
		want *PeekingIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.iter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
