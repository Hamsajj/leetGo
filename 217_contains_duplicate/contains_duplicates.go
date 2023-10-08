package contains_duplicates

// easy question
// https://leetcode.com/problems/contains-duplicate/

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
