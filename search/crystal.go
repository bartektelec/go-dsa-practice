package search

import "math"

func LinearCrystalBall(arr []bool) int {
	for i := 0; i < len(arr); i++ {

		v := arr[i]

		if v {
			return i
		}

	}

	return -1
}

func BinaryCrystalBall(arr []bool) int {
	lo := 0
	hi := len(arr)

	found := -1

	for {

		f_mid := float64(lo + (hi-lo)/2)
		mid := int(math.Round(f_mid))
		v := arr[mid]

		if v {
			found = mid
			hi = mid
		} else {
			lo = mid + 1
		}

		if lo >= hi {
			break
		}
	}

	return found
}
