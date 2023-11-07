// Package accounts_merge
// difficulty: medium
// link: https://leetcode.com/problems/accounts-merge/
package accounts_merge

import (
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	emailMap := make(map[string][][]string)
	for _, acc := range accounts {
		emails := acc[1:]
		name := acc[0]
		sort.Strings(emails)
		emails = removeDuplicates(emails)
		if emailLists, ok := emailMap[name]; ok {
			for i, emailList := range emailLists {
				overlaps, combined := combineIfOverLap(emails, emailList)
				if overlaps {
					emailMap[name][i] = combined
				} else {
					emailMap[name] = append(emailMap[name], emails)
				}
			}
		} else {
			emailMap[name] = [][]string{emails}
		}
	}
	var res [][]string
	for key, val := range emailMap {
		for _, emails := range val {
			res = append(res, append([]string{key}, emails...))
		}
	}
	return res
}

func removeDuplicates(arr []string) []string {
	out := []string{arr[0]}
	prev := arr[0]
	for _, val := range arr[1:] {
		if val == prev {
			continue
		}
		out = append(out, val)
		prev = val
	}
	return out
}

func combineIfOverLap(arr1 []string, arr2 []string) (bool, []string) {
	out := make([]string, 0, len(arr1))
	i, j := 0, 0
	overlap := false
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			out = append(out, arr1[i])
			i++
			j++
			overlap = true
		} else if arr1[i] < arr2[j] {
			out = append(out, arr1[i])
			i++
		} else {
			out = append(out, arr2[j])
			j++
		}
	}
	if !overlap {
		return overlap, []string{}
	}
	if j < len(arr2) {
		out = append(out, arr2[j:]...)
	}
	if i < len(arr1) {
		out = append(out, arr1[i:]...)
	}
	return overlap, out
}
