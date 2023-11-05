package sort

func merge(left, right []int) []int {
	lenLeft := len(left)
	lenRight := len(right)

	aux := make([]int, lenLeft+lenRight)

	i, j := 0, 0
	idx := 0
	for i < lenLeft && j < lenRight {
		if left[i] > right[j] {
			aux[idx] = right[j]
			j += 1
		} else {
			aux[idx] = left[i]
			i += 1
		}
		idx += 1
	}

	for k := i; k < lenLeft; k++ {
		aux[idx] = left[k]
		idx += 1
	}

	for k := j; k < lenRight; k++ {
		aux[idx] = right[k]
		idx += 1
	}

	return aux
}

func mergesort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}

	middle := len / 2
	left := arr[:middle]
	right := arr[middle:]

	l := mergesort(left)
	r := mergesort(right)

	return merge(l, r)
}

func MergeSort(arr []int) []int {
	return mergesort(arr)
}
