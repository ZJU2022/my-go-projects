package slice

import "my-go-projects/generic-utils/internal/slice"

// 删除指定索引的val
func Delete[Src any](src []Src, index int) ([]Src, error) {
	res, _, err := slice.Delete[Src](src, index)
	return res, err
}

//原地删除切片中符合条件的元素

func RemoveIf[Src any](src []Src, m func(index int, src Src) bool) []Src {
	emptyPos := 0

	for index, _ := range src {
		//当前索引是否符合
		if m(index, src[index]) == true {
			continue
		}
		//不符合删除条件原地移动到最前面
		src[emptyPos] = src[index]
		emptyPos += 1
	}
	return src[:emptyPos]
}
