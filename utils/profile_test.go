package utils

import "testing"

func TestBeginTrace_ns(t *testing.T) {
	BegintraceNs()
	EndtraceNs()
	BeginTrace()
	EndTrace()
}
