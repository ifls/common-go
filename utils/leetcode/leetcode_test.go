package leetcode

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"testing"
)

func TestA(t *testing.T) {
	assert.Equal(t, 55, climbStairs3(9))
}

func climbStairs3(n int) int {
	sqrt5 := math.Sqrt(5.0)
	fibn := math.Pow((1.0+sqrt5)/2.0, float64(n+1)) - math.Pow((1.0-sqrt5)/2.0, float64(n+1))
	fibn = math.Ceil(fibn)
	log.Println(sqrt5, fibn)
	return (int)(fibn / sqrt5)
}
