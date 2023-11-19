package graph

func Map[I, O any](data []I, f func(I) O) []O {
	output := make([]O, 0, len(data))

	for _, v := range data {
		output = append(output, f(v))
	}

	return output
}
