package roman_to_integer

import "testing"

func TestRomanToInteger(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "testCase 1",
			input:    "V",
			expected: 5,
		},
		{
			name:     "testCase 2",
			input:    "III",
			expected: 3,
		},
		{
			name:     "testCase 3",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "testCase 4",
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.input)
			if got != tt.expected {
				t.Errorf("expected %d, but received %d", tt.expected, got)
			}
		})
	}
}
