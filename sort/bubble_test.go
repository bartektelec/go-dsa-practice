package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{100, 5, 20, 58, 69, 420}

	BubbleSort(&arr)

	if !reflect.DeepEqual(arr, []int{5, 20, 58, 69, 100, 420}) {
		t.Fatal()
	}
}
