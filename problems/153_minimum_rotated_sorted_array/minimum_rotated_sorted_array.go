// Package minimum_rotated_sorted_array
// difficulty: medium
// link: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
package minimum_rotated_sorted_array

func findMin(nums []int) int {
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
	return nums[left]
}
