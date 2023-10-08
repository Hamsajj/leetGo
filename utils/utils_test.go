package utils

import "testing"

func TestAreSlicesEqual(t *testing.T) {
	tests := []struct {
		name     string
		inputs   [2][]int
		expected bool
	}{
		{
			name:     "equal arrays",
			inputs:   [2][]int{{1, 2, 3}, {1, 2, 3}},
			expected: true,
		},
		{
			name:     "unequal arrays because order",
			inputs:   [2][]int{{1, 3, 2}, {1, 2, 3}},
			expected: false,
		},
		{
			name:     "unequal arrays because length",
			inputs:   [2][]int{{1, 2, 3}, {1, 2}},
			expected: false,
		},
		{
			name:     "unequal arrays",
			inputs:   [2][]int{{1, 2, 3}, {1, 2, 4}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if AreSlicesEqual(tt.inputs[0], tt.inputs[1]) != tt.expected {
				t.Errorf("expected the result to be %t", tt.expected)
			}
		})
	}
}

func TestAreSlicesEqualWithoutOrder(t *testing.T) {
	tests := []struct {
		name     string
		inputs   [2][]int
		expected bool
	}{
		{
			name:     "equal arrays",
			inputs:   [2][]int{{1, 2, 3}, {1, 2, 3}},
			expected: true,
		},
		{
			name:     "equal arrays despite order",
			inputs:   [2][]int{{1, 3, 2}, {1, 2, 3}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if AreSlicesEqualWithoutOrder(tt.inputs[0], tt.inputs[1]) != tt.expected {
				t.Errorf("expected the result to be %t", tt.expected)
			}
		})
	}
}
