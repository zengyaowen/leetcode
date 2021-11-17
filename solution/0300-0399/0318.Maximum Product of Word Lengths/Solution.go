func maxProduct(words []string) int {
	n := len(words)
	masks := make([]int, n)
	for i, word := range words {
		for _, c := range word {
			masks[i] |= (1 << (c - 'a'))
		}
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (masks[i] & masks[j]) == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}