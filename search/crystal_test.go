package search

import (
	"testing"
)

func TestLinearCrystalGood(t *testing.T) {
	haystack := []bool{false, false, false, true, true, true, true, true}

	if res := LinearCrystalBall(haystack); res != 3 {
		t.Fatalf("wrong")
	}
}

func TestLinearCrystalWrong(t *testing.T) {
	haystack := []bool{false, false, false, false, false}

	if res := LinearCrystalBall(haystack); res != -1 {
		t.Fatalf("wrong")
	}
}

func TestBinaryCrystalGood(t *testing.T) {
	haystack := []bool{false, false, false, true, true, true, true, true}

	if res := BinaryCrystalBall(haystack); res != 3 {
		t.Fatalf("wrong")
	}
}

func TestBinaryCrystalWrong(t *testing.T) {
	haystack := []bool{false, false, false, false, false}

	if res := BinaryCrystalBall(haystack); res != -1 {
		t.Fatalf("wrong")
	}
}
