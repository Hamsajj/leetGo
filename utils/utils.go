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
