package sorting

import (
	"cmp"
)

const degradeThreshold int = 8

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

func ParallelMergesort[T cmp.Ordered](input []T) {
	if len(input) < 2 {
		return
	}

	// temporary array for saving intermediate merged results
	tmp := make([]T, len(input))
	parallelMergesort(input, tmp, 0, len(input))
}

func parallelMergesort[T cmp.Ordered](input, tmp []T, start, end int) {
	// degrade to insertion sort if the number of elements to sort is low
	if end-start <= degradeThreshold {
		insertionSort(input, start, end)
		return
	}

	splitPoint := start + (end-start)/2

	ch := make(chan struct{}, 2)
	var sortInHalf = func(s, e int) {
		parallelMergesort(input, tmp, s, e)
		// signal the completion of sorting current half
		ch <- struct{}{}
	}

	// sort the left half [start, splitPoint)
	// and the right half [splitPoint, end) in parallel
	go sortInHalf(start, splitPoint)
	go sortInHalf(splitPoint, end)

	// wait until the sorting of both halves is completed
	<-ch
	<-ch

	parallelMerge(input, tmp, start, splitPoint, end)
}

func parallelMerge[T cmp.Ordered](input, tmp []T, start, mid, end int) {
	parallelMergeSubarrays(input, tmp, start, mid, mid, end, start)
	copy(input[start:end], tmp[start:end])
}

// parallelMergeSubarrays merges the subarray input[lStart, lEnd) with subarray input[rStart, rEnd)
// and write the merged result into tmp[tStart, tStart + lEnd - lStart + rEnd - rStart)
func parallelMergeSubarrays[T cmp.Ordered](input, tmp []T, lStart, lEnd, rStart, rEnd, tStart int) {
	// if both subarrays are empty
	if lStart >= lEnd && rStart >= rEnd {
		return
	}

	// make sure [lStart, lEnd) always represents the larger range
	if rEnd-rStart > lEnd-lStart {
		lStart, lEnd, rStart, rEnd = rStart, rEnd, lStart, lEnd
	}

	lPivotPoint := lStart + (lEnd-lStart)/2
	v := input[lPivotPoint]
	rPivotPoint := findPivotPoint(input, rStart, rEnd, v)
	tPivot := tStart + (lPivotPoint - lStart) + (rPivotPoint - rStart)
	tmp[tPivot] = v

	ch := make(chan struct{}, 2)
	mergeInHalf := func(start1, end1, start2, end2, tStart int) {
		parallelMergeSubarrays(input, tmp, start1, end1, start2, end2, tStart)
		ch <- struct{}{}
	}

	// merge the two halves in parallel
	go mergeInHalf(lStart, lPivotPoint, rStart, rPivotPoint, tStart)
	go mergeInHalf(lPivotPoint+1, lEnd, rPivotPoint, rEnd, tPivot+1)

	// wait until the merge of both halves is completed
	<-ch
	<-ch
}

// findPivotPoint locates the leftmost index i in t[start:end)
// such that t[i] >= target using binary search
func findPivotPoint[T cmp.Ordered](t []T, start, end int, target T) int {
	// loop invariant:
	// t[start, l) < target
	// t[r, end) >= target
	// [l, r) unchecked
	l, r := start, end
	for l < r {
		mid := l + (r-l)/2
		if t[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func insertionSort[T cmp.Ordered](t []T, start, end int) {
	for i := start + 1; i < end; i++ {
		// the element to be swapped with earlier sorted elements
		v := t[i]
		j := i - 1
		for ; j >= start && t[j] > v; j-- {
			t[j+1] = t[j]
		}
		t[j+1] = v
	}
}
