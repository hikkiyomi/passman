package util

func Filter[T any](arr []T, f func(x T) bool) []T {
	result := make([]T, 0, len(arr))

	for _, x := range arr {
		if f(x) {
			result = append(result, x)
		}
	}

	return result
}
