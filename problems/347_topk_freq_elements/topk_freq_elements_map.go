// Package topk_freq_elements
// difficulty: medium
// link: https://leetcode.com/problems/top-k-frequent-elements/
package topk_freq_elements

import (
	"sort"
)

type NumCount struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	// O(n*logn) hopefully
	sort.Ints(nums)
	prevNum := nums[0]
	currentCount := 1
	var res []NumCount
	// O(n)
	for i := 1; i < len(nums); i++ {
		if nums[i] == prevNum {
			currentCount += 1
		} else {
			res = append(res, NumCount{count: currentCount, num: prevNum})
			prevNum = nums[i]
			currentCount = 1
		}
	}
	res = append(res, NumCount{count: currentCount, num: prevNum})
	// O(n*logn) hopefully
	sort.Slice(res, func(i, j int) bool {
		return res[i].count > res[j].count
	})
	maxIter := k
	if k > len(res) {
		maxIter = len(res)
	}
	// O(k)
	finalRes := make([]int, 0, maxIter)
	for i := 0; i < maxIter; i++ {
		finalRes = append(finalRes, res[i].num)
	}
	return finalRes
}
