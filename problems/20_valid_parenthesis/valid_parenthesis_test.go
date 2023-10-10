package valid_parenthesis

import "testing"

func TestValidParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "testCase 1",
			input:    "()",
			expected: true,
		},
		{
			name:     "testCase 2",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "testCase 3",
			input:    "(]",
			expected: false,
		},
		{
			name:     "edgeCase 1",
			input:    "]",
			expected: false,
		},
		{
			name:     "edgeCase 2",
			input:    "(",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			if got != tt.expected {
				t.Errorf("Expected %t, but recieved %t", tt.expected, got)
			}
		})
	}
}
