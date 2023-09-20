package search

import (
	"math"
)

func BinarySearch(arr []int, needle int) bool {
	lo := 0
	hi := len(arr)

	for {
		m := int(math.Round(float64(lo + (hi-lo)/2)))
		v := arr[m]

		if v == needle {
			return true
		}

		if v > needle {
			hi = m
		}

		if v < needle {
			lo = m + 1
		}

		if lo >= hi {
			break
		}
	}

	return false
}
