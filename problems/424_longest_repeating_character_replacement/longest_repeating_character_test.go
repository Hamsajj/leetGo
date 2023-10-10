package longest_repeating_character_replacement

import "testing"

func TestCharacterReplacement(t *testing.T) {

	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "testCase 1",
			s:        "ABAB",
			k:        2,
			expected: 4,
		},
		{
			name:     "testCase 2",
			s:        "AABABBA",
			k:        1,
			expected: 4,
		},
		{
			name:     "edgeCase 1",
			s:        "AABABBA",
			k:        0,
			expected: 2,
		},
		{
			name:     "edgeCase 2",
			s:        "ZX",
			k:        2,
			expected: 2,
		},
		{
			name:     "edgeCase 3",
			s:        "ABAA",
			k:        0,
			expected: 2,
		},
		{
			name:     "edgeCase 4",
			s:        "ABCCCCC",
			k:        2,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if tt.expected != got {
				t.Errorf("Expected %d, but recieved %d", tt.expected, got)
			}
		})
	}
}
