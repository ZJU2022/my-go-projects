package slice

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		Name    string
		Slice   []Number
		wantVal Number
		found   bool
		macth   matchFunc[Number]
	}{
		{
			Name:    "找到了",
			Slice:   []Number{{val: 1}, {val: 12}, {val: 123}},
			wantVal: Number{val: 123},
			found:   true,
			macth:   func(src Number) bool { return src.val == 123 },
		},
		{
			Name:  "没找到",
			Slice: []Number{{val: 1}, {val: 12}, {val: 123}},
			found: false,
			macth: func(src Number) bool { return src.val == 1234 },
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			val, found := Find(testcase.Slice, testcase.macth)
			assert.Equal(t, testcase.wantVal, val)
			assert.Equal(t, testcase.found, found)
		})
	}
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		Name      string
		Slice     []Number
		wantSlice []Number
		match     matchFunc[Number]
	}{
		{
			Name:      "没有符合条件的",
			Slice:     []Number{{val: 2}, {val: 4}},
			wantSlice: []Number{},
			match:     func(src Number) bool { return src.val%4 == 1 },
		},
		{
			Name:      "有符合条件的",
			Slice:     []Number{{val: 2}, {val: 4}},
			wantSlice: []Number{{val: 2}, {val: 4}},
			match:     func(src Number) bool { return src.val == 0 },
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			res := FindAll(testcase.Slice, testcase.match)
			assert.Equal(t, testcase.wantSlice, res)
		})
	}

}

type Number struct {
	val int
}
