package util

import "testing"

func TestRandomSequence(t *testing.T) {
	RandomSequence()
	RandInt()
	RandUint8()
	RandUInt32()
	RandIntBetween(0, 100)
}
