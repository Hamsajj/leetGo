// Package kth_largest
// difficulty: easy
// link: https://leetcode.com/problems/kth-largest-element-in-a-stream/
package kth_largest

import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	if len(nums) <= k {
		h := (*IntHeap)(&nums)
		heap.Init(h)
		return KthLargest{h: h, k: k}
	}
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums[:k] {
		heap.Push(h, num)
	}
	ins := KthLargest{h: h, k: k}
	for _, num := range nums[k:] {
		ins.Add(num)
	}
	return ins
}

func (this *KthLargest) Add(val int) int {
	if this.h.Len() < this.k {
		heap.Push(this.h, val)
	} else if val > (*this.h)[0] {
		(*this.h)[0] = val
		heap.Fix(this.h, 0)
	}
	return (*this.h)[0]
}
