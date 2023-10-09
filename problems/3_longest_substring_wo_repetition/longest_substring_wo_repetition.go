// Package longest_substring_wo_repetition
// difficulty: medium
// link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
package longest_substring_wo_repetition

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	l, r := 0, 1
	maxLength := 0
	charSet := make(map[uint8]struct{})
	charSet[s[l]] = struct{}{}
	for l < len(s) && r < len(s) {
		if _, exists := charSet[s[r]]; exists {
			if (r - l) > maxLength {
				maxLength = r - l
			}
			delete(charSet, s[l])
			l++
			if l >= r {
				r = l + 1
				charSet = map[uint8]struct{}{s[l]: {}}
			}
		} else {
			charSet[s[r]] = struct{}{}
			r++
		}
	}
	if (r - l) > maxLength {
		maxLength = r - l
	}
	return maxLength
}
