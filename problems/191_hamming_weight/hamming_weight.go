package hamming_weight

func hammingWeight(num uint32) int {
	weight := 0
	var candid uint32 = 1
	for i := 0; i < 32; i++ {
		if candid == candid&num {
			weight += 1
		}
		candid = candid << 1
	}

	return weight
}
