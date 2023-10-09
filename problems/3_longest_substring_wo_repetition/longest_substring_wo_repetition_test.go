package longest_substring_wo_repetition

import (
	"testing"
)

func TestLengthOfLongest(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "testCase 1",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "testCase 2",
			input:    "bbbbbbb",
			expected: 1,
		},
		{
			name:     "testCase 3",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "testCase 4",
			input:    "a",
			expected: 1,
		},
		{
			name:     "testCase 5",
			input:    "",
			expected: 0,
		},
		{
			name:     "testCase 6",
			input:    "au",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.input)
			if got != tt.expected {
				t.Errorf("Expected %d, received %d", tt.expected, got)
			}
		})
	}
}
