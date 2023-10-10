// Package longest_repeating_character_replacement
// difficulty: medium
// link: https://leetcode.com/problems/longest-repeating-character-replacement/
package longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	if k == len(s) {
		return k
	}
	left := 0
	maxLength := 1
	charCount := map[byte]int{
		s[0]: 1,
	}
	maxCharCount := 1
	for right := 1; right < len(s); right++ {
		charCount[s[right]]++
		maxCharCount = max(maxCharCount, charCount[s[right]])
		replacementNeeded := right - left + 1 - maxCharCount
		if replacementNeeded > k {
			charCount[s[left]]--
			left++
		}
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
