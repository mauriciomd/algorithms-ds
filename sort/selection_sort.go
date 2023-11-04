package sort

func SelectionSort(arr []int) []int {
	len := len(arr)

	for i := 0; i < len; i++ {
		min := i

		for j := i + 1; j < len; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
