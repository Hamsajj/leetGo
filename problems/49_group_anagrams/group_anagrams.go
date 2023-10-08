// Package group_anagrams
// difficulty: medium
// link: https://leetcode.com/problems/group-anagrams/
package group_anagrams

func groupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[[26]int][]string)
	for _, text := range strs {
		var alphCounts [26]int
		for _, ch := range text {
			alphCounts[int(ch)-int('a')] += 1
		}
		if _, ok := anagramsMap[alphCounts]; ok {
			anagramsMap[alphCounts] = append(anagramsMap[alphCounts], text)
		} else {
			anagramsMap[alphCounts] = []string{text}
		}
	}
	result := make([][]string, 0, len(anagramsMap))
	for _, val := range anagramsMap {
		result = append(result, val)
	}
	return result
}
