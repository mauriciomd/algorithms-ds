package sort

import (
	"strconv"
	"strings"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{2, 1, 5, 8, 9, 4, 3, 7, 0, 6}
	output := BubbleSort(input)

	got := strings.Join(Map(output, strconv.Itoa), ",")
	want := "0,1,2,3,4,5,6,7,8,9"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
