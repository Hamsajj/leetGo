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
		if nums[i] == prevNum+1 {
			currentLCS++
		} else if nums[i] == prevNum {
			continue
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
