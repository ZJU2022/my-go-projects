package slice

//自己传入匹配函数，提高代码通用性，不同的业务需要匹配切片逻辑不同

type matchFunc[T any] func(T) bool

// 匹配切片内一个值
func Find[T any](src []T, match matchFunc[T]) (T, bool) {

	for _, val := range src {
		if match(val) {
			return val, true
		}
	}
	var t T
	return t, false
}

// 匹配切片内所有符合条件的值并返回
func FindAll[T any](src []T, match matchFunc[T]) []T {
	res := make([]T, len(src), len(src)*2)

	for _, val := range src {
		if match(val) {
			res = append(res, val)
		}
	}
	return res
}
