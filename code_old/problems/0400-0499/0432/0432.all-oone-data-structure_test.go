package _0432

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want AllOne
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllOne_Inc(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		this *AllOne
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Inc(tt.args.key)
		})
	}
}

func TestAllOne_Dec(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		this *AllOne
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Dec(tt.args.key)
		})
	}
}

func TestAllOne_GetMaxKey(t *testing.T) {
	tests := []struct {
		name string
		this *AllOne
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.GetMaxKey(); got != tt.want {
				t.Errorf("AllOne.GetMaxKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllOne_GetMinKey(t *testing.T) {
	tests := []struct {
		name string
		this *AllOne
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.GetMinKey(); got != tt.want {
				t.Errorf("AllOne.GetMinKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
