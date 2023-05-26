package printer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type StringLengthTestCase struct {
	Input    []string
	Expected int
}

func testMaxStringLength(t *testing.T, testCase *StringLengthTestCase) {
	res, err := maxStringLength(testCase.Input)

	if testCase.Expected == -1 {
		assert.NotNil(t, err)
	} else {
		assert.Equal(t, testCase.Expected, res)
	}
}

func TestMaxStringLength(t *testing.T) {
	for _, testCase := range []StringLengthTestCase{
		{
			[]string{"ab", "abc"},
			3,
		},
		{
			make([]string, 0),
			-1,
		},
	} {
		testMaxStringLength(t, &testCase)
	}
}
