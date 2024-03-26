package _2642

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want Graph
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {
	type args struct {
		edge []int
	}
	tests := []struct {
		name string
		this *Graph
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddEdge(tt.args.edge)
		})
	}
}

func TestGraph_ShortestPath(t *testing.T) {
	type args struct {
		node1 int
		node2 int
	}
	tests := []struct {
		name string
		this *Graph
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.ShortestPath(tt.args.node1, tt.args.node2); got != tt.want {
				t.Errorf("Graph.ShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
