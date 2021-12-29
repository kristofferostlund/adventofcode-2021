package intutil

func SumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}
