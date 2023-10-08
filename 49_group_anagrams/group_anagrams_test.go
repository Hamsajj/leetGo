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
			if !testArrayOfArraysEqual(result, tt.expected) {
				t.Errorf("expected %v, but recieved %v", tt.expected, result)
			}
		})
	}
}

func testArrayOfArraysEqual(strs1, strs2 [][]string) bool {
	if len(strs1) != len(strs2) {
		return false
	}
	for _, str1 := range strs1 {
		foundEqual := false
		for _, str2 := range strs2 {
			if utils.AreSlicesEqualWithoutOrder(str1, str2) {
				foundEqual = true
				break
			}
		}
		if !foundEqual {
			return false
		}
	}
	return true
}
