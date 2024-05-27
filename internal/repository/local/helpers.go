package local

func voteSum(votes map[uint64]int) int {
	sum := 0
	for _, v := range votes {
		sum += v
	}
	return sum
}
