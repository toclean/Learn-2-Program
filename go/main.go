package main

import "fmt"

func main() {
	fmt.Println(addTwoBinaryStrings("100", "11101"))
	fmt.Println(bubbleSort([]int{0, -100, 22, -14}))
	fmt.Println(insertionSort([]int{-13, -22, 100, -3, 23712}))
	array := []int{-122, -31231, 0, 2, 3, 7}
	fmt.Println(quickSort(array, 0, len(array)-1))
	fmt.Println(linearSearch([]int{13, 44, -77, 8, 13, 200, 11}, 200))
}
