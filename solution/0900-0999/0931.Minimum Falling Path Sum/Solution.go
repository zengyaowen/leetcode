func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	f := make([]int, n)
	for _, row := range matrix {
		g := make([]int, n)
		copy(g, f)
		for j, x := range row {
			if j > 0 {
				g[j] = min(g[j], f[j-1])
			}
			if j+1 < n {
				g[j] = min(g[j], f[j+1])
			}
			g[j] += x
		}
		f = g
	}
	ans := 1 << 30
	for _, x := range f {
		ans = min(ans, x)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}