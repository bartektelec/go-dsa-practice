package search

func LinearSearch(haystack []int, needle int) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}
