func pileBox(box [][]int) (ans int) {
	sort.Slice(box, func(i, j int) bool {
		a, b := box[i], box[j]
		return a[0] < b[0] || (a[0] == b[0] && b[1] < a[1])
	})
	n := len(box)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if box[j][1] < box[i][1] && box[j][2] < box[i][2] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += box[i][2]
		ans = max(ans, f[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}