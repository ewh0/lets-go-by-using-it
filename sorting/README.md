## Go Sort

This module contains a simple implementation of merge sort. The implementation is according to the pseudo-code depicted on the Wikipedia https://en.wikipedia.org/wiki/Merge_sort#Bottom-up_implementation. 

```c
// array A[] has the items to sort; array B[] is a work array
void BottomUpMergeSort(A[], B[], n)
{
    // Each 1-element run in A is already "sorted".
    // Make successively longer sorted runs of length 2, 4, 8, 16... until the whole array is sorted.
    for (width = 1; width < n; width = 2 * width)
    {
        // Array A is full of runs of length width.
        for (i = 0; i < n; i = i + 2 * width)
        {
            // Merge two runs: A[i:i+width-1] and A[i+width:i+2*width-1] to B[]
            // or copy A[i:n-1] to B[] ( if (i+width >= n) )
            BottomUpMerge(A, i, min(i+width, n), min(i+2*width, n), B);
        }
        // Now work array B is full of runs of length 2*width.
        // Copy array B to array A for the next iteration.
        // A more efficient implementation would swap the roles of A and B.
        CopyArray(B, A, n);
        // Now array A is full of runs of length 2*width.
    }
}

//  Left run is A[iLeft :iRight-1].
// Right run is A[iRight:iEnd-1  ].
void BottomUpMerge(A[], iLeft, iRight, iEnd, B[])
{
    i = iLeft, j = iRight;
    // While there are elements in the left or right runs...
    for (k = iLeft; k < iEnd; k++) {
        // If left run head exists and is <= existing right run head.
        if (i < iRight && (j >= iEnd || A[i] <= A[j])) {
            B[k] = A[i];
            i = i + 1;
        } else {
            B[k] = A[j];
            j = j + 1;    
        }
    } 
}

void CopyArray(B[], A[], n)
{
    for (i = 0; i < n; i++)
        A[i] = B[i];
}
```









## APIs

TODO: summarize exported APIs









## Notes



### 1. Why not using sort.Interface as the type for the input collection?

Without use reflection, I don't find out a way to easily make a copy of the underlying collection value as pointed by the [sort.Interface](https://devdocs.io/go/sort/index#Interface). This might be resolved by use type assertion and type switch, but I don't want to continue in that direction. Instead, I create a custom interface to set the expectation for the element passed into the sort API

```go
type Comparable interface {
	Less(Comparable) bool
}
```

and the merge sort expects the elements to be passed in as a slice of Comparable `[]Comparable`

```go
func Mergesort(c []Comparable)
```



To make our custom type comparable

```go
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
```



Testing it

```go
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
```





There is just a little bit more work to make the builtin type to become comparable

```go
type myInt int

func (m myInt) Less(c Comparable) bool {
	n := c.(myInt)
	return m < n
}
```







### 2. Utilize the cmp.Ordered type constraint

[cmp.Ordered](https://devdocs.io/go/cmp/index#Ordered) is a type constraint that is recently introduced in Go 1.21 and it is basically a union of all builtin types which supports  < <= >= > operators. With it, we can easily implement a generic version of merge sort 

```go
func MergesortGx[T cmp.Ordered](c []T) 
```



However, the convenience it offers is limited. Unless later Go decides to support operator overloading, right now we can only pass the builtin types 
