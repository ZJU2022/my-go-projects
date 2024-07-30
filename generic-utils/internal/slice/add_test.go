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
			slice:     []int{1, 12},
			addVal:    123,
			index:     0,
			wantSlice: []int{123, 1, 12},
		},
		{
			name:      "index middle",
			slice:     []int{1, 12, 123},
			addVal:    1234,
			index:     1,
			wantSlice: []int{1, 1234, 12, 123},
		},
		{
			name:    "index out of range",
			slice:   []int{1, 12, 123},
			addVal:  1234,
			index:   7,
			wantErr: err.NewErrIndexOutOfRange(3, 7),
		},
		{
			name:    "index less than 0",
			slice:   []int{1, 12, 123},
			addVal:  1234,
			index:   -1,
			wantErr: err.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:      "index last",
			slice:     []int{123, 100, 101, 102, 102, 102},
			addVal:    233,
			index:     5,
			wantSlice: []int{123, 100, 101, 102, 102, 233, 102},
		},
		{
			name:      "append on last",
			slice:     []int{123, 100, 101, 102, 102, 102},
			addVal:    1234,
			index:     6,
			wantSlice: []int{123, 100, 101, 102, 102, 102, 1234},
		},
		{
			name:    "index out of range",
			slice:   []int{123, 100, 101, 102, 102, 102},
			addVal:  1234,
			index:   7,
			wantErr: err.NewErrIndexOutOfRange(6, 7),
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := Add(testcase.slice, testcase.addVal, testcase.index)
			assert.Equal(t, testcase.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, testcase.wantSlice, res)
		})
	}

}
