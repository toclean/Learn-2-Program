package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addZeros(originalString string, zerosToAdd int) string {
	var sb strings.Builder
	for index := 0; index < zerosToAdd; index++ {
		sb.WriteString("0")
	}
	return sb.String() + originalString
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Print("2 arguments were not provided")
	}

	s1 := args[0]
	s2 := args[1]

	zerosToAdd := 0
	// Add extra zeros
	if len(s1) > len(s2) {
		zerosToAdd = len(s1) - len(s2)
		s2 = addZeros(s2, zerosToAdd)
	} else if len(s2) > len(s1) {
		zerosToAdd = len(s2) - len(s1)
		s1 = addZeros(s1, zerosToAdd)
	}

	fmt.Printf("s1: %s, s2: %s\n", s1, s2)

	result := ""

	var carry int64
	for index := len(s1) - 1; index >= 0; index-- {
		digit1, err1 := strconv.ParseInt(string(s1[index]), 10, 64)
		digit2, err2 := strconv.ParseInt(string(s2[index]), 10, 64)
		if err1 != nil || err2 != nil {
			return
		}

		sum := digit1 + digit2 + carry
		if sum == 3 {
			sum = 1
			carry = 1
		} else if sum == 2 {
			sum = 0
			carry = 1
		} else {
			carry = 0
		}

		result = strconv.FormatInt(sum, 2) + result
		if index == 0 && carry > 0 {
			result = strconv.FormatInt(carry, 2) + result
		}
	}

	fmt.Printf("Result: %s\n", result)
}
