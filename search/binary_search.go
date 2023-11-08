package search

func BinarySearch(v int, arr []int) bool {
	len := len(arr)

	if len == 0 {
		return false
	}

	middle := len / 2
	if v == arr[middle] {
		return true
	} else if v > arr[middle] {
		return BinarySearch(v, arr[middle+1:])
	} else {
		return BinarySearch(v, arr[:middle])
	}
}
