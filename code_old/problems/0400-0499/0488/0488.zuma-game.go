package _0488

var min int
var cache = make(map[string]bool, 0)

func dfs(board string, m map[byte]int, count int) {
	if cache[board] {
		return
	}
	cache[board] = true
	if count >= min {
		return
	}
	for {
		end := true
		for i, num := 1, 1; i < len(board); i++ {
			if board[i] == board[i-1] {
				num++
			} else {
				if num >= 3 {
					board = board[0:i-num] + board[i:]
					end = false
					break
				}
				num = 1
			}
			if i == len(board)-1 && num >= 3 {
				board = board[0:i-num+1] + board[i+1:]
			}
		}
		if end {
			break
		}
	}

	if len(board) < 3 {
		if len(board) == 1 && m[board[0]] > 1 {
			board = ""
			count = count + 2
		}
		if len(board) == 2 {
			if board[0] == board[1] && m[board[0]] > 0 {
				board = ""
				count = count + 1
			} else if m[board[0]] > 1 && m[board[1]] > 1 {
				board = ""
				count = count + 4
			}
		}
		if len(board) == 0 {
			if min > count {
				min = count
			}
			return
		}
	}
	for i, num := 1, 1; i < len(board); i++ {
		if board[i] == board[i-1] {
			num++
		} else {
			if num == 2 && m[board[i-1]] > 0 {
				m[board[i-1]]--
				dfs(board[0:i-num]+board[i:], m, count+1)
				m[board[i-1]]++
			} else if m[board[i-1]] > 1 {
				m[board[i-1]] = m[board[i-1]] - 2
				dfs(board[0:i-num]+board[i:], m, count+2)
				m[board[i-1]] = m[board[i-1]] + 2
			}
			num = 1
		}
	}
	for i := 1; i < len(board); i++ {
		if board[i] == board[i-1] {
			for k, v := range m {
				if k != board[i] && v > 0 {
					m[k]--
					dfs(board[0:i]+string(k)+board[i:], m, count+1)
					m[k]++
				}
			}
		}
	}
}

func findMinStep(board string, hand string) int {
	min = 10
	m := map[byte]int{}
	for i, _ := range hand {
		m[hand[i]]++
	}
	dfs(board, m, 0)
	if min == 10 {
		return -1
	}
	return min
}
