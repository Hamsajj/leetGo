// Package valid_anagram
// difficulty: easy
// link: https://leetcode.com/problems/valid-anagram/
package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alphabetCounts := make([]int, 24)
	for _, ch := range s {
		idx := int(ch) - int('a')
		alphabetCounts[idx] += 1
	}
	for _, ch := range t {
		idx := int(ch) - int('a')
		if alphabetCounts[idx] == 0 {
			return false
		}
		alphabetCounts[idx] -= 1
	}
	return true
}
