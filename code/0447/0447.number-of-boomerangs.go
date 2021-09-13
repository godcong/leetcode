package _0447

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for _, p := range points {
		d := make(map[int]int, len(points))
		for _, q := range points {
			d[(p[0]-q[0])*(p[0]-q[0])+(p[1]-q[1])*(p[1]-q[1])]++
		}
		for _, c := range d {
			ans += c * (c - 1)
		}
	}
	return ans
}
