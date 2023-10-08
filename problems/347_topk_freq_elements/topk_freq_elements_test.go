package topk_freq_elements

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

var tests = []struct {
	name     string
	nums     []int
	k        int
	expected []int
}{
	{
		name:     "testCase 1",
		nums:     []int{4, 2, 2, 3, 1, 1, 1},
		k:        2,
		expected: []int{1, 2},
	},
	{
		name:     "testCase 2",
		nums:     []int{1},
		k:        1,
		expected: []int{1},
	},
	{
		name:     "testCase 3",
		nums:     []int{-1, -1},
		k:        1,
		expected: []int{-1},
	},
	{
		name:     "testCase 4",
		nums:     []int{4, 1, -1, 2, -1, 2, 3},
		k:        2,
		expected: []int{-1, 2},
	},
}

func TestTopKFrequent(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if !utils.AreSlicesEqualWithoutOrder(got, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, got)
			}
		})
	}
}

func TestTopKFrequentHeap(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequentHeap(tt.nums, tt.k)
			if !utils.AreSlicesEqualWithoutOrder(got, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, got)
			}
		})
	}
}
