package slice

import "my-go-projects/generic-utils/internal/err"

func add[T any](src []T, element T, index int) ([]T, error) {
	//在指定索引位置插入一个新元素
	length := len(src)

	if index < 0 || index > length {
		return nil, err.NewErrIndexOutOfRange(length, index)
	}

	//先将切片扩展一个元素
	var zeroValue T
	src = append(src, zeroValue)

	//index往后的所有元素全部后退一位
	//向后移动实际上是copy前一个value
	for i := len(src) - 1; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	//此时src[index] == src[index+1]，需要更新src[index]
	src[index] = element

	return src, nil

}
