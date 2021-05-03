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

func TestIsPalindrome1(t *testing.T) {
	for _, testdata := range testdatas {
		assert.Equal(t, testdata.result, IsPalindrome1(testdata.param))
	}
}

func BenchmarkIsPalindrome1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome1(testdatas[0].param)
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, testdata := range testdatas {
		assert.Equal(t, testdata.result, IsPalindrome2(testdata.param))
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome2(testdatas[0].param)
	}
}
