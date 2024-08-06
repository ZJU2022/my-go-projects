package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		Name    string
		Slice   []int
		wantVal int
		wantErr error
	}{
		{
			Name:    "测试最大值",
			Slice:   []int{1, 12, 123},
			wantVal: 123,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			res, err := Max(testcase.Slice)
			assert.Equal(t, err, testcase.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, testcase.wantVal, res)
		})
	}
}
func TestMin(t *testing.T) {
	testCases := []struct {
		Name    string
		Slice   []int
		wantVal int
		wantErr error
	}{
		{
			Name:    "测试最小值",
			Slice:   []int{1, 12, 123},
			wantVal: 1,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			res, err := Min(testcase.Slice)
			assert.Equal(t, err, testcase.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, testcase.wantVal, res)
		})
	}
}
func TestSum(t *testing.T) {
	testCases := []struct {
		Name    string
		Slice   []int
		wantVal int
		wantErr error
	}{
		{
			Name:    "测试求和",
			Slice:   []int{1, 12, 123},
			wantVal: 136,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			res, err := Sum(testcase.Slice)
			assert.Equal(t, err, testcase.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, testcase.wantVal, res)
		})
	}
}
