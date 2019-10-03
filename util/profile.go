package util

import (
	"fmt"
	"time"
)

var gap float64

func BeginTrace() {
	t := time.Now()
	var s float64 = float64(t.Unix())
	var n float64 = float64(float64(t.Nanosecond()) / 1000000000.0)

	fmt.Printf("begin is %v s %v ns \n", s, n)

	gap = s + n

}

func EndTrace() {
	t := time.Now()
	var s float64 = float64(t.Unix())
	var n float64 = float64(float64(t.Nanosecond()) / 1000000000.0)

	fmt.Printf("end is %v s %v ns \n", s, n)

	gap = s + n - gap

	fmt.Printf("spendTime is %v s \n", gap)
}

var gap_s int64
var gap_ns int64

func BeginTrace_ns() {
	t := time.Now()
	gap_s = t.Unix()
	gap_ns = int64(t.Nanosecond())

	//fmt.Printf("begin is %v s %v ns \n", gap_s, gap_ns)
}

func EndTrace_ns() int64 {
	t := time.Now()
	var s int64 = t.Unix()
	var n int64 = int64(t.Nanosecond())

	fmt.Printf("end is %v s %v ns \n", s, n)

	var st int64 = (s-gap_s)*1000000000 + n - gap_ns

	fmt.Printf("spendTime is %v ns \n", st)
	return st
}
