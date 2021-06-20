package _1600

import (
	"runtime/debug"
)

type ThroneInheritance struct {
	king  string
	edges map[string][]string
	dead  map[string]bool
}

func init() { debug.SetGCPercent(-1) }

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{kingName, map[string][]string{}, map[string]bool{}}
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	t.edges[parentName] = append(t.edges[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	var preorder func(string)
	var ans []string
	preorder = func(name string) {
		if !t.dead[name] {
			ans = append(ans, name)
		}
		for _, childName := range t.edges[name] {
			preorder(childName)
		}
	}
	preorder(t.king)
	return ans
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
