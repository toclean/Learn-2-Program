package main

import "fmt"

func main() {
	fmt.Println(addTwoBinaryStrings("0001", "0010"))
	fmt.Println(bubbleSort([]int{-1, 15, 0, 3}))
	fmt.Println(insertionSort([]int{-1, 100, 15, 3}))
	array := []int{-1, 100, 15, 3}
	fmt.Println(quickSort(array, 0, len(array)-1))
}
