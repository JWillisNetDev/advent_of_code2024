package set

type unit struct{}

type Set[T comparable] struct {
	inner map[T]unit
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{}
}

// Create a new Set from a slice.
func NewSetFromSlice[T comparable](elements []T) *Set[T] {
	inner := sliceToMap(elements)
	return &Set[T]{inner}
}

// Returns whether or not the element was successfully added to the set.
// A return value of false indicates that the element already exists in the set.
func (s *Set[T]) Add(element T) bool {
	if _, ok := s.inner[element]; ok {
		return false
	}

	s.inner[element] = unit{}
	return true
}

// Returns whether or not the element was successfully removed from the set.
// A return value of false indicates the element did not exist in the set to begin with.
func (s *Set[T]) Remove(element T) bool {
	if _, ok := s.inner[element]; ok {
		delete(s.inner, element)
		return true
	}

	return false
}

func (s *Set[T]) Contains(element T) bool {
	_, ok := s.inner[element]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.inner)
}

// Returns a new set containing the unique elements from both the source set and the other set.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	ss := &Set[T]{}
	for e := range s.inner {
		ss.inner[e] = unit{}
	}
	for e := range other.inner {
		ss.inner[e] = unit{}
	}
	return ss
}

// Returns a new set containing only the elements which both the source set and other set share.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	ss := &Set[T]{}
	for e := range s.inner {
		if _, ok := other.inner[e]; ok {
			ss.inner[e] = unit{}
		}
	}
	return ss
}

// Returns a new set containing elements in the source set, but not in the supplied other set.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	ss := &Set[T]{}
	for e := range s.inner {
		if _, ok := other.inner[e]; ok {
			continue
		}
		ss.inner[e] = unit{}
	}
	return ss
}

// Returns a new set containing elements exlcusive to both the source and supplied other set, but not elements which are shared between both sets.
func (s *Set[T]) ExclusiveDifference(other *Set[T]) *Set[T] {
	ss := &Set[T]{}
	for e := range s.inner {
		if _, ok := other.inner[e]; ok {
			continue
		}
		ss.inner[e] = unit{}
	}
	for e := range other.inner {
		if _, ok := s.inner[e]; ok {
			continue
		}
		ss.inner[e] = unit{}
	}
	return ss
}

// Returns a boolean value indicating whether or not the source set is a subset of the given other set.
// A subset contains all the elements of the superset and the superset may contain additional elements.
func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	for e := range other.inner {
		if _, ok := s.inner[e]; !ok {
			return false
		}
	}
	return true
}

// Returns a boolean value indicating wheter or not the source set is a superset of the given other set.
// A subset contains all the elements of the superset and the superset may contain additional elements.
func (s *Set[T]) IsSupersetOf(other *Set[T]) bool {

}

func sliceToMap[T comparable](s []T) map[T]unit {
	m := map[T]unit{}
	for _, v := range s {
		m[v] = unit{}
	}
	return m
}
