package search

import (
	"testing"
)

func TestHappyPath(t *testing.T) {
	needle := 5
	haystack := []int{1, 2, 3, 4, 5}

	if result := LinearSearch(haystack, needle); result != true {
		t.Fatalf("Wrong answer")
	}
}

func TestNotFound(t *testing.T) {
	needle := 12
	haystack := []int{1, 2, 3, 4}

	if result := LinearSearch(haystack, needle); result != false {
		t.Fatalf("Wrong answer")
	}
}
