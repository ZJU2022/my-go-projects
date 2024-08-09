package slice

// 将切片转化为mao，切片值为key
func toMap[T comparable](src []T) map[T]struct{} {
	res := make(map[T]struct{}, len(src))

	for _, val := range src {
		//空结构体不占内存
		res[val] = struct{}{}
	}
	return res
}
