package group_anagrams

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "testCase 1",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:     "testCase 1",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "testCase 1",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "testCase 1",
			strs:     []string{"ab", "ba"},
			expected: [][]string{{"ab", "ba"}},
		},
		{
			name:     "testCase 1",
			strs:     []string{"a", ""},
			expected: [][]string{{"a"}, {""}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.strs)
			if !utils.AreArrayOfArraysEqual(result, tt.expected) {
				t.Errorf("expected %v, but recieved %v", tt.expected, result)
			}
		})
	}
}
