package sort

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx += 1
			arr[idx], arr[i] = arr[i], arr[idx]
		}
	}

	idx += 1
	arr[hi], arr[idx] = arr[idx], arr[hi]

	return idx
}

func quicksort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(arr, lo, hi)
	quicksort(arr, lo, pivot-1)
	quicksort(arr, pivot+1, hi)
}

func QuickSort(arr []int) []int {
	quicksort(arr, 0, len(arr)-1)
	return arr
}
