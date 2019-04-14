package main

func partition(A []int, lo int, hi int) int {
	pivot := A[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if A[j] < pivot {
			A[i], A[j] = A[j], A[i]
			i = i + 1
		}
	}
	A[i], A[hi] = A[hi], A[i]

	return i
}

func quickSort(A []int, lo int, hi int) []int {
	if lo < hi {
		p := partition(A, lo, hi)
		quickSort(A, lo, p-1)
		quickSort(A, p+1, hi)
	}

	return A
}
