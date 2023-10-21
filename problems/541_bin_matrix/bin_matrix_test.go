package bin_matrix

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		mat      [][]int
		expected [][]int
	}{
		{
			name:     "testCase 1",
			mat:      [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expected: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			name:     "testCase 2",
			mat:      [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			expected: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := updateMatrix(tt.mat)
			if !utils.AreArrayOfArraysEqual(got, tt.expected) {
				t.Errorf("expected %v, but recieved %v", tt.expected, got)
			}
		})
	}
}
