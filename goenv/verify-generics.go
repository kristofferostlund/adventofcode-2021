package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	sum := SumIntsOrFloats(ints)
	fmt.Println(sum)
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
