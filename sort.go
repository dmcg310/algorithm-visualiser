package main

import "math/rand"

type SortArray struct {
	Values []int
}

func NewSortArray(size, maxValue int) SortArray {
	values := make([]int, size)
	for i := range values {
		values[i] = rand.Intn(maxValue + 1)
	}

	return SortArray{
		Values: values,
	}
}
