package maputil

import "fmt"

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

// func (p Point) Equal(other Point) bool {
// 	return p.X == other.X && p.Y == other.Y
// }

func init() {
	lookup := map[Point]int{
		{0, 0}: 0,
		{1, 1}: 1,
	}

	fmt.Println(Values(lookup))
	fmt.Println(Keys(lookup))
}

func Values[K comparable, V any](m map[K]V) []V {
	out := make([]V, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}

func Keys[K comparable, V any](m map[K]V) []K {
	out := make([]K, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
