package slice

import (
	"my-go-projects/generic-utils/internal/err"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{1, 12, 123, 1234, 12345},
			index:     0,
			wantSlice: []int{12, 123, 1234, 12345},
			wantVal:   1,
		},
		{
			name:      "index middle",
			slice:     []int{1, 12, 123, 1234, 12345},
			index:     2,
			wantSlice: []int{1, 12, 1234, 12345},
			wantVal:   123,
		},
		{
			name:      "index last",
			slice:     []int{1, 12, 123, 1234, 12345},
			index:     4,
			wantSlice: []int{1, 12, 123, 1234},
			wantVal:   12345,
		},
		{
			name:    "index out of range",
			slice:   []int{1, 12, 123, 1234, 12345},
			index:   5,
			wantErr: err.NewErrIndexOutOfRange(5, 5),
		},
		{
			name:    "index less than 0",
			slice:   []int{1, 12, 123, 1234, 12345},
			index:   -9,
			wantErr: err.NewErrIndexOutOfRange(5, -9),
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			resSlice, val, err := Delete(testcase.slice, testcase.index)
			assert.Equal(t, err, testcase.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, resSlice, testcase.wantSlice)
			assert.Equal(t, val, testcase.wantVal)
		})
	}
}
