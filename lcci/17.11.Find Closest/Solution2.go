func findClosest(words []string, word1 string, word2 string) int {
	d := map[string][]int{}
	for i, w := range words {
		d[w] = append(d[w], i)
	}
	idx1, idx2 := d[word1], d[word2]
	i, j, m, n := 0, 0, len(idx1), len(idx2)
	ans := 100000
	for i < m && j < n {
		t := abs(idx1[i] - idx2[j])
		if t < ans {
			ans = t
		}
		if idx1[i] < idx2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}