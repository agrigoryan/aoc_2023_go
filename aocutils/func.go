package aocutils

func Filter[T any](arr []T, filter func(T) bool) []T {
	filtered := []T{}
	for _, item := range arr {
		if filter(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func Map[T any, R any](arr []T, transform func(T) R) []R {
	transformed := []R{}
	for _, item := range arr {
		transformed = append(transformed, transform(item))
	}
	return transformed
}

func Mapper[T any, R any](transform func(T) R) func([]T) []R {
	return func(arr []T) []R {
		transformed := []R{}
		for _, item := range arr {
			transformed = append(transformed, transform(item))
		}
		return transformed
	}
}

func Filterer[T any](filter func(T) bool) func([]T) []T {
	return func(arr []T) []T {
		filtered := []T{}
		for _, item := range arr {
			if filter(item) {
				filtered = append(filtered, item)
			}
		}
		return filtered
	}
}
