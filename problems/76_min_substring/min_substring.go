package min_substring

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	tCount := stringToArr(t)
	i := 0
	j := len(t)
	sCount := stringToArr(s[i:j])
	minStr := ""
	minStrLen := len(s) + 1
	for i <= j {
		if doesSIncludeT(sCount, tCount) {
			if j-i < minStrLen {
				minStrLen = j - i
				minStr = s[i:j]
			}
			sCount[s[i]-'A']--
			i++
		} else if j <= len(s)-1 {
			j++
			sCount[s[j-1]-'A']++
		} else {
			break
		}
	}
	return minStr
}

func doesSIncludeT(sCount [58]int, tCount [58]int) bool {
	for idx, count := range tCount {
		if sCount[idx] < count {
			return false
		}
	}
	return true
}

func stringToArr(s string) [58]int {
	var counts [58]int
	for _, ch := range s {
		counts[ch-'A']++
	}
	return counts
}
