package test

import (
	"log"
	"sort"
	"testing"
)

func TestStringSort(t *testing.T) {
	s := "wrffxxxs"

	ints := make([]int, 0)
	for _, r := range s {
		ints = append(ints, int(r))
	}
	sort.Ints(ints)
	log.Println()

}
