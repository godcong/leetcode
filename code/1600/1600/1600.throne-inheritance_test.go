package _600

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		kingName string
	}
	tests := []struct {
		name string
		args args
		want ThroneInheritance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.kingName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
