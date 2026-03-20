package _0212

type Trie struct {
	isEnd bool
	cnt   int
	ch    [26]*Trie
	word  string
}

func (t *Trie) insert(s string) {
	node := t
	for _, c := range s {
		if node.ch[c-'a'] == nil {
			node.ch[c-'a'] = new(Trie)
		}
		node = node.ch[c-'a']
		node.cnt++
	}
	node.isEnd = true
	node.word = s
}
func (t *Trie) delete(s string) {
	node := t
	for _, c := range s {
		node = node.ch[c-'a']
		node.cnt--
	}
	node.isEnd = false
}

func (t *Trie) search(s string) *Trie {
	node := t
	for _, c := range s {
		if node.ch[c-'a'] == nil {
			return nil
		}
		node = node.ch[c-'a']
	}
	return node
}

func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	trie := new(Trie)
	for _, w := range words {
		trie.insert(w)
	}
	var ans []string
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(int, int, *Trie)
	dfs = func(x, y int, t *Trie) {
		if t == nil || t.cnt == 0 {
			return
		}
		if t.isEnd {
			ans = append(ans, t.word)
			trie.delete(t.word)
		}
		oldC := board[x][y]
		board[x][y] = 0
		for _, d := range dirs {
			a, b := x+d[0], y+d[1]
			if a >= 0 && a < m && b >= 0 && b < n && board[a][b] != 0 {
				dfs(a, b, t.ch[board[a][b]-'a'])
			}
		}
		board[x][y] = oldC
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, trie.ch[board[i][j]-'a'])
		}
	}
	return ans
}
