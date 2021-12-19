package _0997

type degrees struct {
	in, out int
}

func findJudge(n int, trust [][]int) int {
	d := make([]degrees, n)
	var degreesMax int
	for _, t := range trust {
		d[t[1]-1].in++
		d[t[0]-1].out++
		degreesMax = updateDegreesMax(d, t[1]-1, degreesMax)
	}

	if d[degreesMax].in == n-1 && d[degreesMax].out == 0 {
		return degreesMax
	}
	return -1
}

func updateDegreesMax(d []degrees, l, r int) int {
	if d[l].in > d[r].in {
		return l
	}
	return r
}
