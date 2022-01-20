package set

import (
	"fmt"
	"strings"
)

var (
	keyExist = struct{}{}
)

const maxInt = int(^uint(0) >> 1)

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable](ts ...T) *Set[T] {
	s := &Set[T]{make(map[T]struct{}, len(ts))}
	s.Add(ts...)
	return s
}

func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		s.m[item] = keyExist
	}
}

func (s *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s.m, item)
	}
}

func (s *Set[T]) Pop() T {
	var nonExistence = new(T)
	for item := range s.m {
		delete(s.m, item)
		return item
	}
	return *nonExistence
}

func (s *Set[T]) Pop2() (T, bool) {
	var nonExistence = new(T)
	for item := range s.m {
		delete(s.m, item)
		return item, true
	}
	return *nonExistence, false
}

func (s *Set[T]) Has(items ...T) bool {
	has := false
	for _, item := range items {
		// Note: trick: has will override the initial value of has(false)
		if _, has = s.m[item]; !has {
			break
		}
	}
	return has
}

func (s *Set[T]) HasAny(items ...T) bool {
	has := false
	for _, item := range items {
		if _, has = s.m[item]; has {
			break
		}
	}
	return has
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set[T]) IsEqual(t *Set[T]) bool {
	if s.Size() != t.Size() {
		return false
	}

	eq := true

	s.Each(func(item T) bool {
		_, eq = t.m[item]
		return eq
	})

	return eq
}

func (s *Set[T]) IsSubset(t *Set[T]) bool {
	if s.Size() < t.Size() {
		return false
	}

	subset := true

	t.Each(func(item T) bool {
		_, subset = s.m[item]
		return subset
	})

	return subset
}

func (s *Set[T]) IsSuperset(t *Set[T]) bool {
	return t.IsSubset(s)
}

func (s *Set[T]) Each(f func(item T) bool) {
	for item := range s.m {
		if !f(item) {
			break
		}
	}
}

func (s *Set[T]) Copy() *Set[T] {
	newSet := &Set[T]{make(map[T]struct{}, s.Size())}
	s.Each(func(item T) bool {
		newSet.Add(item)
		return true
	})
	return newSet
}

func (s *Set[T]) String() string {
	v := make([]string, 0, s.Size())
	for item := range s.m {
		v = append(v, fmt.Sprintf("%v", item))
	}
	return fmt.Sprintf("[%s]", strings.Join(v, ", "))
}

func (s *Set[T]) List() []T {
	res := make([]T, 0, s.Size())
	s.Each(func(item T) bool {
		res = append(res, item)
		return true
	})
	return res
}

func (s *Set[T]) Merge(t *Set[T]) {
	for item := range t.m {
		s.m[item] = keyExist
	}
}

func (s *Set[T]) Separate(t *Set[T]) {
	for item := range t.m {
		delete(s.m, item)
	}
}

func Union[T comparable](sets ...*Set[T]) *Set[T] {
	maxPos := -1
	maxSize := 0
	for i, set := range sets {
		if l := set.Size(); l > maxSize {
			maxSize = l
			maxPos = i
		}
	}
	if maxSize == 0 {
		return &Set[T]{}
	}

	u := sets[maxPos].Copy()
	for i, set := range sets {
		if i == maxPos {
			continue
		}
		for item := range set.m {
			u.m[item] = keyExist
		}
	}
	return u
}

func Difference[T comparable](s1 *Set[T], sets ...*Set[T]) *Set[T] {
	s := s1.Copy()
	for _, set := range sets {
		s.Separate(set)
	}
	return s
}

func Intersection[T comparable](sets ...*Set[T]) *Set[T] {
	minPos := -1
	minSize := maxInt
	for i, set := range sets {
		if l := set.Size(); l < minSize {
			minSize = l
			minPos = i
		}
	}
	if minSize == maxInt || minSize == 0 {
		return &Set[T]{}
	}

	t := sets[minPos].Copy()
	for i, set := range sets {
		if i == minPos {
			continue
		}
		for item := range t.m {
			if _, has := set.m[item]; !has {
				delete(t.m, item)
			}
		}
	}
	return t
}

func SymmetricDifference[T comparable](s *Set[T], t *Set[T]) *Set[T] {
	u := Difference(s, t)
	v := Difference(t, s)
	return Union(u, v)
}
