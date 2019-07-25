package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodePriorityQueue []*ListNode

func (pq ListNodePriorityQueue) Len() int {
	return len(pq)
}

func (pq ListNodePriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq ListNodePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *ListNodePriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *ListNodePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *ListNodePriorityQueue) Remove(i int) {
	n := pq.Len() - 1
	pq.Swap(i, n)
	pq.Pop()
}
