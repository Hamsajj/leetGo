// Package product_array_except_self
// difficulty: medium
// link:https://leetcode.com/problems/product-of-array-except-self/
package product_array_except_self

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	multiplier := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = multiplier
		multiplier *= nums[i]
	}
	multiplier = 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= multiplier
		multiplier *= nums[i]
	}
	return answer
}
