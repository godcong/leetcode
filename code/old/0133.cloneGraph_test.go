package old

import (
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	type args struct {
		node *Node133
	}

	node1 := &Node133{
		Val:       1,
		Neighbors: nil,
	}
	node2 := &Node133{
		Val:       2,
		Neighbors: nil,
	}
	node3 := &Node133{
		Val:       3,
		Neighbors: nil,
	}
	node4 := &Node133{
		Val:       4,
		Neighbors: nil,
	}
	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node2, node4)
	node4.Neighbors = append(node4.Neighbors, node1, node3)
	tests := []struct {
		name string
		args args
		want *Node133
	}{
		{
			name: "",
			args: args{
				node: node1,
			},
			want: node1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !deepEqualCloneGraph(t, got, tt.want, 10) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deepEqualCloneGraph(t *testing.T, got, want *Node133, count int) bool {
	if count <= 0 {
		return true
	}
	if got.Val != want.Val {
		t.Errorf("cloneGraph() = %v, want %v", got.Val, want.Val)
		return false
	}
	for i := range got.Neighbors {
		if !deepEqualCloneGraph(t, got.Neighbors[i], want.Neighbors[i], count-1) {
			return false
		}
	}
	return true
}
