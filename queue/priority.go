package queue

import (
	"container/heap"

	"github.com/c2nc/snippets/cmp"
)

type priorityQueueImpl []cmp.Lesser

func (pq priorityQueueImpl) Len() int {
	return len(pq)
}

func (pq priorityQueueImpl) Less(i, j int) bool {
	return pq[i].Less(pq[j])
}

func (pq priorityQueueImpl) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueueImpl) Push(x interface{}) {
	item := x.(cmp.Lesser)
	*pq = append(*pq, item)
}

func (pq *priorityQueueImpl) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type PriorityQueue struct {
	priorityQueueImpl
}

func NewPriorityQueue() *PriorityQueue {
	var pq PriorityQueue
	heap.Init(&pq.priorityQueueImpl)
	return &pq
}

func (pq *PriorityQueue) Push(item cmp.Lesser) {
	heap.Push(&pq.priorityQueueImpl, item)
}

func (pq *PriorityQueue) Pop() cmp.Lesser {
	return heap.Pop(&pq.priorityQueueImpl).(cmp.Lesser)
}

func (pq *PriorityQueue) Front() cmp.Lesser {
	return pq.priorityQueueImpl[0]
}

func (pq *PriorityQueue) Length() int {
	return pq.priorityQueueImpl.Len()
}
