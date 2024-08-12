package slice

func Reverse[T any](src []T) []T {
	res := make([]T, 0, len(src))

	for i := len(src) - 1; i >= 0; i-- {
		res = append(res, src[i])
	}
	return res
}
