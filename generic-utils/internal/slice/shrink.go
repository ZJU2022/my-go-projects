package slice

//go只有扩容机制，没有缩容机制，自定义编写缩容函数

func capaCity(capacity, length int) (int, bool) {

	if capacity <= 64 {
		return capacity, false
	}
	if capacity > 2048 && (capacity/length >= 2) {
		factor := 0.625
		return int(float32(capacity) * float32(factor)), true
	}
	if capacity <= 2048 && (capacity/length >= 4) {
		return capacity / 2, true
	}
	return capacity, false
}

func Shrink[T any](src []T) []T {
	capacity, length := cap(src), len(src)
	newcapcity, changed := capaCity(capacity, length)

	if !changed {
		return src
	}
	s := make([]T, 0, newcapcity)
	s = append(s, src...)
	return s
}
