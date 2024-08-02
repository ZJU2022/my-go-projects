package slice

import (
	"my-go-projects/generic-utils/internal/err"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{1, 12, 123},
			addVal:    1234,
			index:     0,
			wantSlice: []int{1234, 1, 12, 123},
		},
		{
			name:    "index -1",
			slice:   []int{1, 12, 123},
			addVal:  0,
			index:   -1,
			wantErr: err.NewErrIndexOutOfRange(3, -1),
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := Add(testcase.slice, testcase.index, testcase.addVal)
			assert.Equal(t, testcase.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, res, testcase.wantSlice)
		})
	}

}
