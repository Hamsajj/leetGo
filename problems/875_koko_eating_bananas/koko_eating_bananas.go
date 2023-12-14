package koko_eating_bananas

import (
	"math"
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	left, right := 1, piles[len(piles)-1]
	for left < right {
		middle := (left + right) / 2
		if isFeasible(piles, middle, h) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left

}

func isFeasible(piles []int, k int, h int) bool {
	timeToEat := 0.0
	for _, bananas := range piles {
		timeToEat += math.Ceil(float64(bananas) / float64(k))
		if int(timeToEat) > h {
			return false
		}
	}
	return true
}
