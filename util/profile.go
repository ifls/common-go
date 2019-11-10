package util

import (
	"log"
	"time"
)

var gap float64

func BeginTrace() {
	t := time.Now()
	s := float64(t.Unix())
	n := float64(t.Nanosecond()) / 1000000000.0

	log.Printf("begin is %v s %v ns \n", s, n)

	gap = s + n

}

func EndTrace() {
	t := time.Now()
	s := float64(t.Unix())
	n := float64(t.Nanosecond()) / 1000000000.0

	log.Printf("end is %v s %v ns \n", s, n)

	gap = s + n - gap

	log.Printf("spendTime is %v s \n", gap)
}

var gapS int64
var gapNs int64

func BegintraceNs() {
	t := time.Now()
	gapS = t.Unix()
	gapNs = int64(t.Nanosecond())

	//log.Printf("begin is %v s %v ns \n", gap_s, gap_ns)
}

func EndtraceNs() int64 {
	t := time.Now()
	s := t.Unix()
	n := int64(t.Nanosecond())

	log.Printf("end is %v s %v ns \n", s, n)

	st := (s-gapS)*1000000000 + n - gapNs

	log.Printf("spendTime is %v ns \n", st)
	return st
}
