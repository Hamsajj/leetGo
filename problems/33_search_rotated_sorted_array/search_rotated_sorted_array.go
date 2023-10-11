// Package search_rotated_sorted_array
// difficulty: medium
// link: https://leetcode.com/problems/search-in-rotated-sorted-array/
package search_rotated_sorted_array

func findMinIndex(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	var middle int
	for left < right {
		middle = (left + right) / 2
		if nums[middle] < nums[right] {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

func search(nums []int, target int) int {
	minIndex := findMinIndex(nums)
	length := len(nums)

	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		ind := (minIndex + middle) % length
		if nums[ind] == target {
			return ind
		} else if nums[ind] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}
