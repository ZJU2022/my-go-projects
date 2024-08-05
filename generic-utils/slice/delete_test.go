package slice

import (
	"my-go-projects/generic-utils/internal/err"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		Name      string
		Slice     []int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			Name:      "index 0",
			Slice:     []int{1, 12, 123},
			index:     0,
			wantSlice: []int{12, 123},
		},
		{
			Name:    "index -1",
			Slice:   []int{1, 12, 123},
			index:   -1,
			wantErr: err.NewErrIndexOutOfRange(3, -1),
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			res, err := Delete(testcase.Slice, testcase.index)
			assert.Equal(t, testcase.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, res, testcase.wantSlice)
		})
	}
}
