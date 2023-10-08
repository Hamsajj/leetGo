// Package longest_consecutive_sequence
// difficulty: medium
// link: https://leetcode.com/problems/longest-consecutive-sequence/
package longest_consecutive_sequence

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	sort.Ints(nums)
	currentLCS := 1
	maxLCS := currentLCS
	prevNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prevNum {
			continue
		}
		if nums[i] == prevNum+1 {
			currentLCS++
		} else {
			if currentLCS > maxLCS {
				maxLCS = currentLCS
			}
			currentLCS = 1
		}
		prevNum = nums[i]
	}
	if currentLCS > maxLCS {
		maxLCS = currentLCS
	}
	return maxLCS
}

func longestConsecutiveSet(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	hashSet := make(map[int]struct{})
	for _, num := range nums {
		hashSet[num] = struct{}{}
	}

	currentLCS := 1
	maxLCS := 1
	for num, _ := range hashSet {
		if _, ok := hashSet[num-1]; ok {
			continue
		}
		nextExists := true
		for i := 1; nextExists; i++ {
			if _, ok := hashSet[num+i]; ok {
				currentLCS += 1
			} else {
				nextExists = false
				if currentLCS > maxLCS {
					maxLCS = currentLCS
				}
				currentLCS = 1
			}
		}
	}
	if currentLCS > maxLCS {
		maxLCS = currentLCS
	}
	return maxLCS
}
