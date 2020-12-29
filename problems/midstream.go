package problems

import "container/heap"

type MedianFinder struct {
	maxHeap PriorityQueue
	minHeap PriorityQueue
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{}
	mf.Init()
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	item := Item{
		value:    num,
		priority: num,
	}

	heap.Push(&this.maxHeap, &item)
	//this.maxHeap.Push(&item)
	midItem := heap.Pop(&this.maxHeap).(*Item)
	//midItem := this.maxHeap.Pop().(*Item)
	midItem.index = 0
	midItem.priority = -midItem.priority
	heap.Push(&this.minHeap, &midItem)
	//this.minHeap.Push(midItem)

	if len(this.minHeap)-len(this.maxHeap) == 2 {
		midItem = heap.Pop(&this.minHeap).(*Item)
		//midItem = this.minHeap.Pop().(*Item)
		midItem.index = 0
		midItem.priority = -midItem.priority
		heap.Push(&this.maxHeap, midItem)
		//this.maxHeap.Push(midItem)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.minHeap) == 0 {
		return 0
	}

	if len(this.maxHeap) == len(this.minHeap) {
		a := this.maxHeap[0].value
		b := this.minHeap[0].value
		return (float64(a) + float64(b)) / 2
	}

	return float64(this.minHeap[0].value)
}

func (this *MedianFinder) Init() {
	heap.Init(&this.maxHeap)
	heap.Init(&this.minHeap)
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
