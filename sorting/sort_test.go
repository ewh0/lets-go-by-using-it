package sorting

import (
	"math/rand"
	"slices"
	"testing"
)

type myInt int

func (m myInt) Less(c Comparable) bool {
	n := c.(myInt)
	return m < n
}

type Character struct {
	first, last string
}

func (p Character) Less(c Comparable) bool {
	n := c.(Character)
	if p.first != n.first {
		return p.first < n.first
	}

	return p.last < n.last
}

func TestMergesort(t *testing.T) {
	var input1 = []Comparable{myInt(5), myInt(4), myInt(1), myInt(2), myInt(3)}
	Mergesort(input1)

	var sortFunc = func(a, b Comparable) int {
		i := a.(myInt)
		j := b.(myInt)
		return int(i) - int(j)
	}

	if !slices.IsSortedFunc(input1, sortFunc) {
		t.Errorf("failed to sort input slice [5, 4, 1, 2, 3]")
	}

	var input2 = []Comparable{myInt(1), myInt(2), myInt(3), myInt(4), myInt(5)}
	Mergesort(input2)
	if !slices.IsSortedFunc(input2, sortFunc) {
		t.Errorf("failed to sort input slice [1, 2, 3, 4, 5]")
	}

	var input3 = []Comparable{myInt(5), myInt(4), myInt(3), myInt(2), myInt(1)}
	Mergesort(input3)
	if !slices.IsSortedFunc(input3, sortFunc) {
		t.Errorf("failed to sort input slice [5, 4, 3, 2, 1]")
	}

	var input4 = []Comparable{Character{first: "Tom", last: "The Cat"}, Character{first: "Jerry", last: "The Mouse"}, Character{first: "Spike", last: "The Dog"}}
	Mergesort(input4)
	if !slices.IsSortedFunc(input4, func(a, b Comparable) int {
		p := a.(Character)
		q := b.(Character)
		if p.Less(q) {
			return -1
		} else if q.Less(p) {
			return 1
		} else {
			return 0
		}
	}) {
		t.Errorf("failed to sort customized types")
	}
}

func TestMergesortGx(t *testing.T) {
	var input = make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		input = append(input, rand.Int())
	}
	MergesortGx(input)

	if !slices.IsSorted(input) {
		t.Errorf("failed to sort the randomly generated integer slice")
	}

	var input2 = make([]float64, 0, 10000)
	for i := 0; i < 10000; i++ {
		input2 = append(input2, rand.Float64())
	}
	MergesortGx(input2)
	if !slices.IsSorted(input2) {
		t.Errorf("failed to sort the randomly generated floating number slice")
	}
}

func TestParallelMergesort(t *testing.T) {
	var input = []int{5, 4, 1, 2, 4, 6, 3, 1, 8, 9}
	ParallelMergesort(input)

	if !slices.IsSorted(input) {
		t.Errorf("failed to sort the randomly generated integer slice")
	}

	var input2 = make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		input2 = append(input2, rand.Int())
	}
	ParallelMergesort(input2)
	if !slices.IsSorted(input2) {
		t.Errorf("failed to sort the randomly generated integer slice")
	}
}
