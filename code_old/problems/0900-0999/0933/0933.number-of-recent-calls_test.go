package _0933

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want RecentCounter
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

func TestRecentCounter_Ping(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		this *RecentCounter
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Ping(tt.args.t); got != tt.want {
				t.Errorf("RecentCounter.Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}
