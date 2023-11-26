package min_substring

import "testing"

func TestMinSubstring(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		s        string
		t        string
		expected string
	}{
		{
			name:     "testCase 1",
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
		{
			name:     "testCase 2",
			s:        "a",
			t:        "a",
			expected: "a",
		},
		{
			name:     "testCase 3",
			s:        "a",
			t:        "aa",
			expected: "",
		},
		{
			name:     "testCase 4",
			s:        "ab",
			t:        "ba",
			expected: "ab",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, got)
			}
		})
	}
}
