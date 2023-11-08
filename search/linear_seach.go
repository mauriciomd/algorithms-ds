package search

func LinearSearch(v int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if v == arr[i] {
			return true
		}
	}

	return false
}
