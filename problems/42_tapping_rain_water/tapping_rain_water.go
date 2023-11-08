// Package tapping_rain_water
// difficulty: hard
// link: https://leetcode.com/problems/trapping-rain-water/
package tapping_rain_water

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	trapped := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] < maxLeft {
				trapped += maxLeft - height[left]
			} else {
				maxLeft = height[left]
			}
			left++
		} else {
			if height[right] < maxRight {
				trapped += maxRight - height[right]
			} else {
				maxRight = height[right]
			}
			right--
		}
	}
	return trapped
}
