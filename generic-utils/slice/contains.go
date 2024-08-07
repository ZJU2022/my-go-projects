package slice

// Contains 判断 src 里面是否存在 dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, func(src T) bool {
		return src == dst
	})
}

// ContainsFunc 判断 src 里面是否存在 dst
// 优先使用Contains,这是底层实现
func ContainsFunc[T any](src []T, equal func(src T) bool) bool {
	//遍历src，拿src中的每个值和dst比较

	for _, val := range src {
		if equal(val) {
			return true
		}
	}
	return false

}

// ContainsAny 判断 src 里面是否存在 dst 中的任何一个元素

// ContainsAnyFunc 判断 src 里面是否存在 dst 中的任何一个元素
// 你应该优先使用 ContainsAny

// ContainsAll 判断 src 里面是否存在 dst 中的所有元素

// ContainsAllFunc 判断 src 里面是否存在 dst 中的所有元素
// 你应该优先使用 ContainsAll
