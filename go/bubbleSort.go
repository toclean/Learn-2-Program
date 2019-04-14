package main

func bubbleSort(A []int) []int {
	n := len(A)
	for {
		swapped := false
		for i := 0; i < n-1; i++ {
			if A[i] > A[i+1] {
				A[i], A[i+1] = A[i+1], A[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return A
}
