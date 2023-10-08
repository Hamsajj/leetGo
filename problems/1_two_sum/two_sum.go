// Package two_sum
// difficulty: easy
// link: https://leetcode.com/problems/two-sum/
package two_sum

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)
	for i, num := range nums {
		if val, ok := numToIndex[target-num]; ok {
			return []int{i, val}
		}
		numToIndex[num] = i
	}
	// should not reach here because the description of the solution
	// guarantees exactly one solution exists for each test case
	return []int{}
}
