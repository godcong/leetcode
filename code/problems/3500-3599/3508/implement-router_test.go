package _3508

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		memoryLimit int
	}
	tests := []struct {
		name string
		args args
		want Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.memoryLimit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_AddPacket(t *testing.T) {
	type args struct {
		source      int
		destination int
		timestamp   int
	}
	tests := []struct {
		name string
		this *Router
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.AddPacket(tt.args.source, tt.args.destination, tt.args.timestamp); got != tt.want {
				t.Errorf("Router.AddPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_ForwardPacket(t *testing.T) {
	tests := []struct {
		name string
		this *Router
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.ForwardPacket(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router.ForwardPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_GetCount(t *testing.T) {
	type args struct {
		destination int
		startTime   int
		endTime     int
	}
	tests := []struct {
		name string
		this *Router
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.GetCount(tt.args.destination, tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("Router.GetCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
