package puzzle4

type Set[V comparable] struct {
	elements map[V]struct{}
}

func (s *Set[V]) Has(v V) bool {
	_, has := s.elements[v]
	return has
}

func (s *Set[V]) Intersection(other *Set[V]) *Set[V] {
	intersection := &Set[V]{elements: make(map[V]struct{})}
	x, y := s, other
	if len(other.elements) < len(s.elements) {
		x, y = other, s
	}

	for k := range x.elements {
		if y.Has(k) {
			intersection.elements[k] = struct{}{}
		}
	}

	return intersection
}

func (s *Set[V]) Difference(other *Set[V]) *Set[V] {
	return nil
}
