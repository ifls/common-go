package algorithm

import (
	"math/rand"
	"testing"
	"time"
)

var datas []Sort

func init() {
	datas = make([]Sort, 0)
	rand.Seed(time.Now().UnixNano() % 1000)
	for i := 0; i < 100; i++ {
		s := Sort{Length: 10000}
		s.randomArray()
		datas = append(datas, s)
	}

}

func TestBubbleSort(t *testing.T) {
	for _, s := range datas {
		s.once()
	}
}

func TestStdSort(t *testing.T) {
	for _, s := range datas {
		s.once()
	}
}
