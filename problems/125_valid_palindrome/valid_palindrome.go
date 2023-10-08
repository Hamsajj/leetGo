// Package valid_palindrome
// difficulty: easy
// link: https://leetcode.com/problems/valid-palindrome/
package valid_palindrome

func isPalindrome(s string) bool {
	if (len(s)) == 1 {
		return true
	}
	var cleanedInput []rune
	for _, char := range s {
		if !isNumber(char) {
			char = lowerCaseAlphabet(char)
			if char == -1 {
				continue // skip if is not number and not alphabet
			}
		}
		cleanedInput = append(cleanedInput, char)
	}
	if (len(cleanedInput)) <= 1 {
		return true
	}
	for i, j := 0, len(cleanedInput)-1; i < j; i, j = i+1, j-1 {
		if cleanedInput[i] != cleanedInput[j] {
			return false
		}
	}
	return true
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

// returns -1 if it is not an alphabet
func lowerCaseAlphabet(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char
	}
	if char >= 'A' && char <= 'Z' {
		return char - ('A' - 'a')
	}

	return -1

}
