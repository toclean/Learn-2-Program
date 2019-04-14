package main

import (
	"reflect"
	"testing"
)

func TestAddingZeros(t *testing.T) {
	s1 := "001"
	result := addZeros(s1, 2)
	if result != "00001" {
		t.Errorf("Correct number of zeros were not added, got: %v, want: %v", result, "00001")
	}
}

func TestBinaryStringAddition(t *testing.T) {
	total := addTwoBinaryStrings("01", "100001")
	if total != "100010" {
		t.Errorf("Sum was incorrect, got: %s, want: %s", total, "100010")
	}
}

func TestBubbleSort(t *testing.T) {
	original := []int{-4, 100, -33, 0}
	new := bubbleSort(original)
	if !reflect.DeepEqual(new, []int{-33, -4, 0, 100}) {
		t.Errorf("Array was not correctly sorted, got: %v, want: %v", new, []int{-33, -4, 0, 100})
	}
}

func TestInsertionSort(t *testing.T) {
	original := []int{-4, 100, -33, 0}
	new := insertionSort(original)
	if !reflect.DeepEqual(new, []int{-33, -4, 0, 100}) {
		t.Errorf("Array was not correctly sorted, got: %v, want: %v", new, []int{-33, -4, 0, 100})
	}
}

func TestQuickSort(t *testing.T) {
	original := []int{-4, 100, -33, 0}
	new := quickSort(original, 0, len(original)-1)
	if !reflect.DeepEqual(new, []int{-33, -4, 0, 100}) {
		t.Errorf("Array was not correctly sorted, got: %v, want: %v", new, []int{-33, -4, 0, 100})
	}
}

func TestLinearSearch(t *testing.T) {
	array := []int{-4, 100, -33, 0}
	result := linearSearch(array, 100)
	if result != 1 {
		t.Errorf("Item was not correctly identified in the array, got: %v, want: %v", result, 1)
	}
}
