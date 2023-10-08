// Package topk_freq_elements
// difficulty: medium
// link: https://leetcode.com/problems/top-k-frequent-elements/
package topk_freq_elements

import (
	"container/heap"
	"sort"
)

type MaxHeap []NumCount

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(NumCount))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequentHeap(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	// O(n*logn) hopefully
	sort.Ints(nums)
	prevNum := nums[0]
	currentCount := 1
	maxHeap := &MaxHeap{}
	// O(n*logn)
	for i := 1; i < len(nums); i++ {
		if nums[i] == prevNum {
			currentCount += 1
		} else {
			heap.Push(maxHeap, NumCount{count: currentCount, num: prevNum})
			prevNum = nums[i]
			currentCount = 1
		}
	}
	heap.Push(maxHeap, NumCount{count: currentCount, num: prevNum})
	// O(n*logn) hopefully

	maxIter := k
	if k > maxHeap.Len() {
		maxIter = maxHeap.Len()
	}
	// O(k)
	finalRes := make([]int, maxIter)
	for i := 0; i < maxIter; i++ {
		finalRes[i] = heap.Pop(maxHeap).(NumCount).num
	}
	return finalRes
}
