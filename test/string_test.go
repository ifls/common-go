package test

import (
	"fmt"
	"log"
	"os"
	"sort"
	"testing"
)

type I interface {
	V()
}

type I2 interface {
	P()
}

type S struct {
}

func (s S) V()  {

}

func (s *S) P()  {

}

func ee(i I)  {

}

func ee2(i I2)  {

}

func TestPv(t *testing.T) {
	s := S{}
	s.P()
	s.V()

	ee(s)
	//ee2(s)
	s2 := &S{}
	s2.V()
	s2.P()

	ee(s2)
	ee2(s2)
}
func TestStringSort(t *testing.T) {
	s := "wrffxxxs"

	ints := make([]int, 0)
	for i, r := range s {
		log.Printf("%T %#v", i, i)
		log.Printf("%T %#v", r, r)
		ints = append(ints, int(r))
	}
	sort.Ints(ints)
	var a  map[string]bool = nil
	print("sss")
	log.Println(a["s"])
	delete(a, "ss")
	log.Println("Hello", 23)
	fmt.Print("ss", 3, 4, 5, "s")
	fmt.Fprint(os.Stdout, "22")
	fmt.Printf("%q, %#q", 3, "3")
	fmt.Printf("%x\n", int(^uint(0) >> 1))
	fmt.Printf("%x\n", ^uint64(0))
}

func TestStringByte(t *testing.T)  {
	s := "ssss"
	log.Printf("%T %#v", s[2], s[2])
}
