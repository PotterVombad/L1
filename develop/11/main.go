package main

import "fmt"

type Set[T comparable] struct {
	el []T
}

func MakeSet[T comparable](s []T) Set[T] {
	dict := make(map[T]bool)
	for _, v := range s {
		_, ok := dict[v]
		if !ok {
			dict[v] = true
		}
	}
	keys := make([]T, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	return Set[T] {
		el: keys,
	}
}

func (s Set[T]) Intersection(sec Set[T]) []T {
	dict := make(map[T]bool)
	res := make([]T, 0, len(s.el))
	for _, v := range s.el {
		dict[v] = true
	}
	for _, w := range sec.el {
		_, ok := dict[w]
		if ok {
			res = append(res, w)
		}
	}
	return res
}

func main() {
	first := []int{1, 2, 3, 4}
	second := []int{2, 1, 5, 0}

	fmt.Println(MakeSet(first).Intersection(MakeSet(second)))
}
