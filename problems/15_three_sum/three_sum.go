// Package three_sum
// difficulty: medium
// link: https://leetcode.com/problems/3sum/
package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	answersSet := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		for l, r := i+1, len(nums)-1; l < r; {
			if nums[l]+nums[r] == target {
				answersSet = append(answersSet, []int{nums[i], nums[l], nums[r]})
				for j := l + 1; j < len(nums); j++ {
					if nums[j] == nums[l] {
						l++
					} else {
						l++
						break
					}
				}
				for j := r - 1; j >= 0; j-- {
					if nums[j] == nums[r] {
						r--
					} else {
						r--
						break
					}
				}
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				l++
			}
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				i++
			}
		}
	}
	return answersSet
}
