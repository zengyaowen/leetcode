func maxKelements(nums []int, k int) (ans int64) {
	h := &hp{nums}
	heap.Init(h)
	for i := 0; i < k; i++ {
		v := h.pop()
		ans += int64(v)
		h.push((v + 2) / 3)
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }