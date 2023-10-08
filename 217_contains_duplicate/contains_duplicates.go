// Package contains_duplicates
// difficulty: easy
// link: https://leetcode.com/problems/contains-duplicate/
package contains_duplicates

func containsDuplicate(nums []int) bool {
	numExists := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := numExists[num]; ok {
			return true
		} else {
			numExists[num] = struct{}{}
		}
	}
	return false
}
