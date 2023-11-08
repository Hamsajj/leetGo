package min_cost_climbing

func minCostClimbingStairs(cost []int) int {
	first := cost[0]
	second := cost[1]
	if len(cost) <= 2 {
		return min(first, second)
	}

	for i := 2; i < len(cost); i++ {
		curr := cost[i] + min(second, first)
		first = second
		second = curr
	}
	return min(first, second)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
