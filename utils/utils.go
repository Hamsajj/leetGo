package utils

import (
	"cmp"
	"slices"
)

// AreSlicesEqualWithoutOrder test if two slices are equal, ignoring the order
func AreSlicesEqualWithoutOrder[T ~[]E, E cmp.Ordered](a, b T) bool {
	if len(a) != len(b) {
		return false
	}
	slices.Sort(a)
	slices.Sort(b)
	return AreSlicesEqual(a, b)
}

// AreSlicesEqual test if two slices are equal
func AreSlicesEqual[T ~[]E, E cmp.Ordered](a, b T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func AreArrayOfArraysEqual[T ~[]E, E cmp.Ordered](nums1, nums2 []T) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	if !doesAllELementsExistInArray(nums1, nums2) {
		return false
	}
	if !doesAllELementsExistInArray(nums2, nums1) {
		return false
	}
	return true
}

func doesAllELementsExistInArray[T ~[]E, E cmp.Ordered](nums1, nums2 []T) bool {
	for _, num1 := range nums1 {
		foundEqual := false
		for _, num2 := range nums2 {
			if AreSlicesEqualWithoutOrder(num1, num2) {
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
