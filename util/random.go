package util

import (
	"math"
	"math/rand"
	"time"
)

func Random() {

}

const n int = 150000

func RandomSequence() []int {

	var s []int
	for i := 0; i < n; i++ {
		s = append(s, rand.Int())
	}

	return s
}

func RandUInt32() uint32 {
	return rand.Uint32()
}

func RandUint8() uint8 {
	return uint8(rand.Uint32() % 256)
}

func RandInt() int {
	return rand.Int()
}

func RandIntBetween(min, max int) int {
	min = int(math.Abs(float64(min)))
	max = int(math.Abs(float64(max)))
	min = int(math.Min(float64(min), float64(max)))
	max = int(math.Max(float64(min), float64(max)))

	return int(rand.Int31n(int32(max-min))) + min
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
