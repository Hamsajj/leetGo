package valid_anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		name     string
		expected bool
		inputs   [2]string
	}{
		{
			name:     "are anagrams",
			expected: true,
			inputs:   [2]string{"anagram", "nagaram"},
		},
		{
			name:     "are not anangram",
			expected: false,
			inputs:   [2]string{"rat", "car"},
		},
		{
			name:     "different length",
			expected: false,
			inputs:   [2]string{"rat", "ra"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if isAnagram(tt.inputs[0], tt.inputs[1]) != tt.expected {
				t.Errorf(
					"expected result of isAnagram(%s,%s) to equal %t",
					tt.inputs[0], tt.inputs[1], tt.expected,
				)
			}
		})
	}
}
