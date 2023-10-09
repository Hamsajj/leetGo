// Package container_with_most_water
// difficulty: medium
// link: https://leetcode.com/problems/container-with-most-water/
package container_with_most_water

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		var currentArea int
		if height[left] < height[right] {
			currentArea = height[left] * (right - left)
			left++
		} else {
			currentArea = height[right] * (right - left)
			right--
		}
		if maxArea < currentArea {
			maxArea = currentArea
		}
	}
	return maxArea
}
