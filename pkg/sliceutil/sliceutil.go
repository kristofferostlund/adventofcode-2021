package sliceutil

func RemoveAt[T any](coll []T, i int) []T {
	var empty T
	out := append([]T(nil), coll...)

	copy(out[i:], out[i+1:])
	out[len(out)-1] = empty
	return out[:len(out)-1]
}
