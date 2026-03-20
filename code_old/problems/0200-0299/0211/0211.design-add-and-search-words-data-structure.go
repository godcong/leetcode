package _0211

import (
	"runtime/debug"
)

type WordDictionary struct {
	items map[int][]string
}

func Constructor() WordDictionary {
	return WordDictionary{make(map[int][]string)}
}

func (this *WordDictionary) AddWord(word string) {
	l := len(word)
	if list, ok := this.items[l]; ok {
		this.items[l] = append(list, word)
	} else {
		this.items[l] = []string{word}
	}
}

func (this *WordDictionary) Search(word string) bool {
	l := len(word)
	if list, ok := this.items[l]; ok {
		for _, w := range list {
			i := 0
			for i < l {
				if w[i] != word[i] && word[i] != '.' {
					break
				}
				i++
			}
			if i == l {
				return true
			}
		}
	}
	return false
}
func init() { debug.SetGCPercent(-1) }

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
