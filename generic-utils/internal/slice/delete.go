package slice

import "my-go-projects/generic-utils/internal/err"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)

	if index < 0 || index >= length {
		var zero T
		return nil, zero, err.NewErrIndexOutOfRange(length, index)
	}

	//返回删除的值
	res := src[index]

	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	//去掉最后一个重复的元素，左闭右开
	src = src[:length-1]

	return src, res, nil

}
