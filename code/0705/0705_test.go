package _0705

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyHashSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Constructor()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
