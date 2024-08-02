package slice

import (
	"errors"

	"github.com/ecodeclub/ekit"
)

var ErrEmptySlice = errors.New("slice is empty")

// 返回切片最大值
// 在使用 float32 或者 float64 的时候要小心精度问题
func Max[T ekit.RealNumber](ts []T) (T, error) {

	if len(ts) == 0 {
		var zero T
		return zero, ErrEmptySlice
	}

	res := ts[0]

	for i := 1; i < len(ts); i++ {
		if ts[i] > res {
			res = ts[i]
		}
	}
	return res, nil
}

// 返回切片最小值，该方法会假设你至少会传入一个值
// 在使用 float32 或者 float64 的时候要小心精度问题
func Min[T ekit.RealNumber](ts []T) (T, error) {

	if len(ts) == 0 {
		var zero T
		return zero, ErrEmptySlice
	}

	res := ts[0]

	for i := 1; i < len(ts); i++ {
		if ts[i] < res {
			res = ts[i]
		}
	}
	return res, nil
}

// 切片内数据求和
// 在使用 float32 或者 float64 的时候要小心精度问题
func Sum[T ekit.Number](ts []T) (T, error) {

	if len(ts) == 0 {
		var zero T
		return zero, ErrEmptySlice
	}

	var res T

	for _, n := range ts {
		res += n
	}
	return res, nil
}
