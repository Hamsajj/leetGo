// Package valid_parenthesis
// difficulty: easy
// link: https://leetcode.com/problems/valid-parentheses/

package valid_parenthesis

func isValid(input string) bool {
	charPair := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}
	var s stack
	for _, char := range input {
		if opening, exists := charPair[char]; exists {
			var last int32
			if s.Len() == 0 {
				return false
			}
			s, last = s.Pop()
			if last != opening {
				return false
			}
		} else {
			s = s.Push(char)
		}
	}
	if s.Len() > 0 {
		return false
	}
	return true
}

type stack []int32

func (s stack) Push(v int32) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int32) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Len() int {
	return len(s)
}
