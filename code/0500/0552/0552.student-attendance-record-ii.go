package _0552

const mod int = 1e9 + 7

type matrix [6][6]int

var _matrix = matrix{
	{1, 1, 0, 1, 0, 0},
	{1, 0, 1, 1, 0, 0},
	{1, 0, 0, 1, 0, 0},
	{0, 0, 0, 1, 1, 0},
	{0, 0, 0, 1, 0, 1},
	{0, 0, 0, 1, 0, 0},
}

func checkRecord(n int) (ans int) {
	res := _matrix.pow(n)
	for _, v := range res[0] {
		ans = (ans + v) % mod
	}
	return ans
}

func (a matrix) mul(b matrix) matrix {
	c := matrix{}
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int) matrix {
	res := matrix{}
	for i := range res {
		res[i][i] = 1
	}
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}
