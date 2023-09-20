package sort

func BubbleSort(arr *[]int) {
	n := len(*arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			a := (*arr)[j]
			b := (*arr)[j+1]

			if a > b {
				(*arr)[j] = b
				(*arr)[j+1] = a
			}

		}
	}
}
