package main

func linearSearch(A []int, target int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == target {
			return i
		}
	}

	return -1
}
