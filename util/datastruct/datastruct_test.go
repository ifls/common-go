package datastruct

import (
	"container/list"
	"log"
	"testing"
)

func TestA(t *testing.T) {
	l := list.New()
	l.Init()
	l.PushBack(1)
	l.PushFront("s")
	log.Printf("%#v \n", l)

	l2 := new(list.List)
	l2.PushFront(1)
	log.Printf("%#v \n", l)
}
