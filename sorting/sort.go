package sorting

import (
	"cmp"
)

type Comparable interface {
	Less(Comparable) bool
}

func MergesortGx[T cmp.Ordered](c []T) {
	if len(c) < 2 {
		return
	}

	l := len(c)
	tmp := make([]T, l)
	// in each pass, double the size of subarray to be merged with
	// the adjacent subarray
	for width := 1; width < l; width *= 2 {
		end := 0
		for start := 0; start+width < l; start += width * 2 {
			end = min(l, start+width*2)
			mergesortGx(c, tmp, start, start+width, end)
		}
		copy(c[:end], tmp[:end])
	}
}

func mergesortGx[T cmp.Ordered](a, tmp []T, start, mid, end int) {
	i := start
	l := start
	r := mid
	for ; l < mid && r < end; i++ {
		if a[l] < a[r] {
			tmp[i] = a[l]
			l++
		} else {
			tmp[i] = a[r]
			r++
		}
	}

	for l < mid {
		tmp[i] = a[l]
		i++
		l++
	}

	for r < end {
		tmp[i] = a[r]
		i++
		r++
	}
}

// Mergesort performs an in-place sorting of the input collection in O(nlogn) time
func Mergesort(c []Comparable) {
	if len(c) < 2 {
		return
	}

	l := len(c)
	tmp := make([]Comparable, l)
	// in each pass, double the size of subarray to be merged with
	// the adjacent subarray
	for width := 1; width < l; width *= 2 {
		end := 0
		for start := 0; start+width < l; start += width * 2 {
			end = min(l, start+width*2)
			mergesort(c, tmp, start, start+width, end)
		}
		copy(c[:end], tmp[:end])
	}
}

// mergesort performs the actual merge of left subarray [start, mid)
// with the right subarray [mid, end)
func mergesort(a, tmp []Comparable, start, mid, end int) {
	i := start
	l := start
	r := mid
	for ; l < mid && r < end; i++ {
		if a[l].Less(a[r]) {
			tmp[i] = a[l]
			l++
		} else {
			tmp[i] = a[r]
			r++
		}
	}

	for l < mid {
		tmp[i] = a[l]
		i++
		l++
	}

	for r < end {
		tmp[i] = a[r]
		i++
		r++
	}
}
