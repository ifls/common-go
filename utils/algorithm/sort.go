package algorithm

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

type Sort struct {
	data     []int
	Length   int
	LessTime int64
	swapTime int64
	spend    int64
	sorted   bool
}

func (s *Sort) Len() int {
	return s.Length
}

func (s *Sort) Swap(i, j int) {
	s.swapTime++
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s *Sort) Less(i, j int) bool {
	s.LessTime++
	return s.data[i] < s.data[j]
}

func (s *Sort) randomArray() {
	for i := 0; i < s.Length; i++ {
		s.data = append(s.data, rand.Int())
	}
}

func (s *Sort) print() {
	s.isSorted()
	log.Printf("isSorted = %v Compare = %v, Swap = %v, Spend = %#v\n", s.sorted, s.LessTime, s.swapTime, time.Duration(s.spend).String())
}

func (s *Sort) isSorted() {
	for i := 0; i < len(s.data)-1; i++ {
		if !s.Less(i, i+1) {
			s.sorted = false
			break
		}
	}
	s.sorted = true
}

func (s *Sort) once() {
	start := time.Now().UnixNano()
	//s.BubbleSort()
	//s.InsertSort()
	//s.QuickSort()
	sort.Sort(s)
	s.spend = time.Now().UnixNano() - start
	s.print()
}

func (s *Sort) BubbleSort() {
	for i := 0; i < s.Length; i++ {
		for j := 0; j < s.Length-i-1; j++ {
			if !s.Less(j, j+1) {
				s.Swap(j, j+1)
			}
		}
	}
}

func (s *Sort) InsertSort() {
	for i := 0; i < s.Length; i++ {
		for j := i + 1; j < s.Length-1; j++ {
			if !s.Less(j, j+1) {
				s.Swap(j, j+1)
			}
		}
	}
}

func (s *Sort) QuickSort() {
	s._quick_sort1(0, len(s.data)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
func (s *Sort) _quick_sort1(begin, end int) {
	// terminal
	if begin >= end {
		return
	}
	// current logic
	//mid := (begin + end) >> 1
	mid := s._partition(begin, end)
	// drill down
	s._quick_sort1(begin, mid-1)
	s._quick_sort1(mid+1, end)
}

// 分区、返回基准索引p、使得arr[l:p - 1] < arr[p] && arr[p + 1: r] > arr[p]
func (s *Sort) _partition(l int, r int) int {
	p := l     // 取第一个元素为基点
	j := p + 1 // j 表示大于基点和小于基点的分界下标
	for i := l + 1; i <= r; i++ {
		if s.Less(i, p) {
			s.Swap(i, j)
			j++
		}
	}
	// 整个数组遍历完成之后、j所指向的元素就是第一个大于arr[p]的元素、所以将p与j-1进行交换
	// 最后返回j - 1
	s.Swap(p, j-1)
	return j - 1
}
