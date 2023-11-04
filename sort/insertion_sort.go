package sort

func InsertionSort(arr []int) []int {
	len := len(arr)

	for i := 1; i < len; i++ {
		j := i

		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j -= 1
		}
	}

	return arr
}
