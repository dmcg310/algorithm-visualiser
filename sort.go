package main

import "math/rand"

type SortArray struct {
	Values []int
}

type Algorithm interface {
	Step(array *SortArray)
	IsFinished() bool
	Reset(array *SortArray)
}

type Bubble struct {
	outerIndex int
	innerIndex int
	swapped    bool
	finished   bool
}

type SortingAlgorithm struct {
	Name      string
	Array     SortArray
	Algorithm Algorithm
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

func NewSortingAlgorithm(name string) SortingAlgorithm {
	array := NewSortArray(50, 20)

	var alg Algorithm
	switch name {
	case "Bubble":
		alg = &Bubble{}
	default:
		alg = &Bubble{}
	}

	alg.Reset(&array)

	return SortingAlgorithm{
		Name:      name,
		Array:     array,
		Algorithm: alg,
	}
}

func (sa *SortingAlgorithm) Step() {
	sa.Algorithm.Step(&sa.Array)
}

func (sa *SortingAlgorithm) IsFinished() bool {
	return sa.Algorithm.IsFinished()
}

func (sa *SortingAlgorithm) Reset() {
	sa.Array = NewSortArray(50, 20)
	sa.Algorithm.Reset(&sa.Array)
}

func (b *Bubble) Step(array *SortArray) {
	if b.outerIndex >= len(array.Values)-1 {
		b.finished = true
		return
	}

	if b.innerIndex >= len(array.Values)-1-b.outerIndex {
		if !b.swapped {
			b.finished = true
		}

		b.outerIndex++
		b.innerIndex = 0
		b.swapped = false

		return
	}

	if array.Values[b.innerIndex] > array.Values[b.innerIndex+1] {
		array.Values[b.innerIndex], array.Values[b.innerIndex+1] =
			array.Values[b.innerIndex+1], array.Values[b.innerIndex]
		b.swapped = true
	}

	b.innerIndex++
}

func (b *Bubble) IsFinished() bool {
	return b.finished
}

func (b *Bubble) Reset(array *SortArray) {
	b.outerIndex = 0
	b.innerIndex = 0
	b.swapped = false
	b.finished = false
}
