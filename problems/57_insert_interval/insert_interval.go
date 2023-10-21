// Package insert_interval
// difficulty: medium
// link: https://leetcode.com/problems/i(nsert-interval/
package insert_interval

func insert(intervals [][]int, ni []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{ni}
	}
	var finalRes [][]int
	i := 0
	for ; i < len(intervals); i++ {
		if intervals[i][1] < ni[0] {
			finalRes = append(finalRes, intervals[i])
		} else {
			break
		}
	}
	mergedInterval := []int{-1, -1}
	if i == len(intervals) {
		mergedInterval[0] = ni[0]
	} else {
		mergedInterval[0] = minInt(intervals[i][0], ni[0])
	}
	for ; i < len(intervals); i++ {
		if intervals[i][0] > ni[1] {
			break
		}
	}
	if i == 0 {
		mergedInterval[1] = ni[1]
	} else {
		mergedInterval[1] = maxInt(intervals[i-1][1], ni[1])
	}
	finalRes = append(finalRes, mergedInterval)
	for ; i < len(intervals); i++ {
		finalRes = append(finalRes, intervals[i])
	}
	return finalRes
}

func maxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func minInt(i1, i2 int) int {
	if i1 > i2 {
		return i2
	}
	return i1
}
