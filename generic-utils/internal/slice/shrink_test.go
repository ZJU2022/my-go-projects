package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestShrink(t *testing.T) {
	testCases := []struct {
		name      string
		orginCap  int
		length    int
		expectCap int
	}{
		{
			name:      "小于64",
			orginCap:  32,
			length:    4,
			expectCap: 32,
		},
		{
			name:      "小于2048, 不足1/4",
			orginCap:  1000,
			length:    200,
			expectCap: 500,
		},
		{
			name:      "小于2048, 超过1/4",
			orginCap:  1000,
			length:    260,
			expectCap: 1000,
		},
		{
			name:      "大于2048，不足一半",
			orginCap:  3000,
			length:    60,
			expectCap: 1875,
		},
		{
			name:      "大于2048，大于一半",
			orginCap:  3000,
			length:    2000,
			expectCap: 3000,
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			l := make([]int, 0, testcase.orginCap)

			for i := 0; i < testcase.length; i++ {
				l = append(l, i)
			}
			l = Shrink(l)
			assert.Equal(t, cap(l), testcase.expectCap)

		})
	}

}
