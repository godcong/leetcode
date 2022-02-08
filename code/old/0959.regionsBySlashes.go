package old

/*
959. 由斜杠划分区域
在由 1 x 1 方格组成的 N x N 网格 grid 中，每个 1 x 1 方块由 /、\ 或空格构成。这些字符会将方块划分为一些共边的区域。

（请注意，反斜杠字符是转义的，因此 \ 用 "\\" 表示。）。

返回区域的数目。



示例 1：

输入：
[
  " /",
  "/ "
]
输出：2
解释：2x2 网格如下：

示例 2：

输入：
[
  " /",
  "  "
]
输出：1
解释：2x2 网格如下：

示例 3：

输入：
[
  "\\/",
  "/\\"
]
输出：4
解释：（回想一下，因为 \ 字符是转义的，所以 "\\/" 表示 \/，而 "/\\" 表示 /\。）
2x2 网格如下：

示例 4：

输入：
[
  "/\\",
  "\\/"
]
输出：5
解释：（回想一下，因为 \ 字符是转义的，所以 "/\\" 表示 /\，而 "\\/" 表示 \/。）
2x2 网格如下：

示例 5：

输入：
[
  "//",
  "/ "
]
输出：3
解释：2x2 网格如下：



提示：

1 <= grid.length == grid[0].length <= 30
grid[i][j] 是 '/'、'\'、或 ' '。
*/
func regionsBySlashes(grid []string) int {
	n := len(grid)
	var parent = make([]int, 4*n*n)

	for i := range parent {
		parent[i] = i
	}
	var regionsBySlashesFind func(int) int
	regionsBySlashesFind = func(x int) int {
		if parent[x] != x {
			parent[x] = regionsBySlashesFind(parent[x])
		}
		return parent[x]
	}
	var regionsBySlashesUnion func(int, int)
	regionsBySlashesUnion = func(x, y int) {
		parent[regionsBySlashesFind(x)] = regionsBySlashesFind(y)
	}
	for i, s := range grid {
		for j, c := range s {
			num := i*n + j
			if c != '\\' {
				regionsBySlashesUnion(4*num+0, 4*num+1)
				regionsBySlashesUnion(4*num+2, 4*num+3)
			}
			if c != '/' {
				regionsBySlashesUnion(4*num+0, 4*num+3)
				regionsBySlashesUnion(4*num+2, 4*num+1)
			}
			if i > 0 {
				regionsBySlashesUnion(4*num, 4*num-4*n+2)
			}
			if i < n-1 {
				regionsBySlashesUnion(4*num+2, 4*num+4*n+0)
			}
			if j > 0 {
				regionsBySlashesUnion(4*num+1, 4*num-4+3)
			}
			if j < n-1 {
				regionsBySlashesUnion(4*num+3, 4*num+4+1)
			}
		}
	}
	ans := 0
	for i := range parent {
		if regionsBySlashesFind(i) == i {
			ans++
		}
	}
	return ans
}
