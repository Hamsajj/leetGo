package valid_sudoku

import "testing"

func TestIsValidSudoku(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		board    [][]byte
		expected bool
	}{
		{
			name:     "testCase 1",
			expected: true,
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := isValidSudoku(tt.board)
			if got != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, got)
			}
		})
	}
}
