func maximumGood(statements [][]int) int {
	n := len(statements)
	check := func(k int) int {
		cnt := 0
		for i, s := range statements {
			if ((k >> i) & 1) == 1 {
				for j := 0; j < n; j++ {
					if s[j] < 2 && ((k>>j)&1) != s[j] {
						return 0
					}
				}
				cnt++
			}
		}
		return cnt
	}
	ans := 0
	for k := 0; k < (1 << n); k++ {
		ans = max(ans, check(k))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}