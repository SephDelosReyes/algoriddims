package main

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	a := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	result := linearSearch(a, 420)
	if !result {
		t.Errorf("Expected true, got %t", result)
	}
	result2 := linearSearch(a, 42069)
	if result2 {
		t.Errorf("Expected false, got %t", result2)
	}

}

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	result := binarySearch(a, 3)
	if !result {
		t.Errorf("Expected true, got %t", result)
	}
	fmt.Println("done")
	result2 := binarySearch(a, 42069)
	if result2 {
		t.Errorf("Expected false, got %t", result2)
	}
	fmt.Println("done")
	result3 := binarySearch(a, 69420)
	if !result3 {
		t.Errorf("Expected false, got %t", result3)
	}
	fmt.Println("done")
	result4 := binarySearch(a, 1)
	if !result4 {
		t.Errorf("Expected false, got %t", result4)
	}
	fmt.Println("done")
}
