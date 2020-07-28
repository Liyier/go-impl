package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var cases = [][]int{
		{1, 2, 3, 4, 5},
		{2, 1, 4, 6, 5},
		{5, 4, 3, 2, 1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			bubbleSort(c)
			t.Log(c)
			if !isSorted(c, false) {
				t.Error("sort fail")
			}
		})
	}

}

func isSorted(arr []int, desc bool) bool {
	length := len(arr)
	if desc {
		for i := 0; i < length-1; i++ {
			if arr[i] < arr[i+1] {
				return false
			}
		}
	} else {
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] {
				return false
			}
		}
	}
	return true
}
