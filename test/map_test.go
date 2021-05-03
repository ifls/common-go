package test

import (
	"log"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

func BenchmarkMap(b *testing.B) {
	b.StopTimer()
	cache := make(map[int]int, 200000000)
	rnd := rand.Int()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cache[rnd+i] = rnd + i
	}
	b.StopTimer()
}

func TestGOMAXPROCS(t *testing.T) {
	log.Println(runtime.GOMAXPROCS(1))
	log.Println(runtime.GOMAXPROCS(2))
}

type T struct {
	a int
}

func (tv T) Mv(a int) int          { return 0 } // value receiver
func (tp *T) Mp(f float32) float32 { return 1 } // pointer receiver

var tt T

type T2 struct {
	a int
}

func (tp *T2) Mp(f float32) float32 { return 1 } // pointer receiver

var tt2 T2

func TestPointerReceiver(t *testing.T) {
	//tt.Mv(2)
	//T.Mv(tt, 3)
	//T.Mp(tt, 3)
	//a := &tt
	//T.Mv(a, 3.0)
	//T.Mp(a, 3.0)
	//
	//T2.Mp(tt2, 3)
	//a2 := &tt2
	//T2.Mp(a2, 3.0)

}

func TestInterface(t *testing.T) {
	var i interface{}

	a := struct {
		ii int32
	}{ii: 0}

	i = a
	log.Printf("%T, %#v", a, a)
	log.Printf("%T, %#v", i, i)

	log.Println(reflect.TypeOf(i).Kind())
	log.Println(reflect.ValueOf(i).String())
}

func TestRunTime(t *testing.T) {
	var a []byte
	if a == nil {
		log.Println("nil", unsafe.Sizeof(a))
	}

	var b string
	if b == "" {
		log.Println("nil", unsafe.Sizeof(b))
	}
}
