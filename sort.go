package main

import "math/rand"

type SortArray struct {
	Values []int
}

type Algorithm interface {
	Step(array *SortArray)
	IsFinished() bool
	Reset(array *SortArray)
	GetCurrentIndices() (int, int)
}

type Bubble struct {
	outerIndex int
	innerIndex int
	swapped    bool
	finished   bool
}

type Selection struct {
	currentIndex int
	minIndex     int
	finished     bool
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
	case "Selection":
		alg = &Selection{}
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

// BUBBLE

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

func (b *Bubble) GetCurrentIndices() (int, int) {
	return b.innerIndex, b.innerIndex + 1
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

// SELECTION

func (s *Selection) Step(array *SortArray) {
	if s.currentIndex >= len(array.Values)-1 {
		s.finished = true
		return
	}

	// Find the minimum element in the unsorted portion
	if s.minIndex == s.currentIndex {
		s.minIndex = s.currentIndex
		for i := s.currentIndex + 1; i < len(array.Values); i++ {
			if array.Values[i] < array.Values[s.minIndex] {
				s.minIndex = i
			}
		}
	}

	// Swap the found minimum element with the first element of the unsorted portion
	if s.minIndex != s.currentIndex {
		array.Values[s.currentIndex], array.Values[s.minIndex] =
			array.Values[s.minIndex], array.Values[s.currentIndex]
	}

	// Move to the next element
	s.currentIndex++
	s.minIndex = s.currentIndex
}

func (s *Selection) GetCurrentIndices() (int, int) {
	return s.currentIndex, s.minIndex
}

func (s *Selection) IsFinished() bool {
	return s.finished
}

func (s *Selection) Reset(array *SortArray) {
	s.currentIndex = 0
	s.minIndex = 0
	s.finished = false
}

