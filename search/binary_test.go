package search

import (
	"testing"
)

func TestHappy(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	needle := 7

	if result := BinarySearch(haystack, needle); result != true {
		t.Fatalf("Wrong result")
	}
}

func TestNotFoudn(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	needle := 10

	if result := BinarySearch(haystack, needle); result != false {
		t.Fatalf("Wrong result")
	}
}
