package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	param  string
	result bool
}

var testdatas []TestData

func init() {
	testdatas = []TestData{
		{
			param:  "A man, a plan, a canal: Panama",
			result: true,
		},
		{
			param:  "race a car",
			result: false,
		},
	}
}

func Test125(t *testing.T) {
	for _, testdata := range testdatas {
		assert.Equal(t, testdata.result, isPalindrome(testdata.param))
	}
}
