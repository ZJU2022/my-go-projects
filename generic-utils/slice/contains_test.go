package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		Name    string
		Slice   []int
		dst     int
		wantVal bool
	}{
		{
			Name:    "dst exist",
			Slice:   []int{1, 12, 123},
			dst:     1,
			wantVal: true,
		},
		{
			Name:    "dst not exist",
			Slice:   []int{1, 12, 123},
			dst:     11,
			wantVal: false,
		},
		{
			Name:    "src length is 0",
			Slice:   []int{},
			dst:     1,
			wantVal: false,
		},
		{
			Name:    "src is nil",
			dst:     1,
			wantVal: false,
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			res := Contains(testcase.Slice, testcase.dst)
			assert.Equal(t, testcase.wantVal, res)
		})
	}

}
