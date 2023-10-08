package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "testCase 1",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "testCase 2",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "testCase 3 - empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "testCase 4 - one char",
			input:    " a",
			expected: true,
		},
		{
			name:     "testCase 5 - one space",
			input:    " ",
			expected: true,
		},
		{
			name:     "testCase 6",
			input:    "0P",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if got != tt.expected {
				t.Errorf("expected %t, but got %t", tt.expected, got)
			}
		})
	}
}
