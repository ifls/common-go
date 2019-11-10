package net

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	slice := make([]int, 3)
	a := new(int)
	fmt.Printf("%T\n%T\n", slice, a)
	var s []int
	s1 := make([]float64, 0)
	fmt.Printf("%T\n%T\n", s, s1)
	fmt.Printf("%v\n%v\n", s, s1)

	var colors map[string]string

	// 将Red的代码加入到映射

	fmt.Printf("[%v-\n", colors["Red"])
}

func TestEncode(t *testing.T) {
	//dataText := &pb.User{
	//	Uid:  100001,
	//	Id:   "432352235325353",
	//	Name: "heyifeng",
	//}
	//
	//bytes, err := Encode(dataText)
	//if err != nil {
	//	fmt.Printf("%v \n ", err)
	//}
	//
	//fmt.Printf("%v %v\n", string(bytes), bytes)
	//
	//user, err := Decode(bytes, "User")
	//if err != nil {
	//	fmt.Printf("error %v\n", err)
	//	return
	//}
	//fmt.Printf("1 %v\n", dataText)
	//fmt.Printf("2, %v\n", user.(*pb.User))
}

func TestToBytes(t *testing.T) {
	ToBytes(nil)
	UInt32ToBytes(2)
	BytesToUInt32([]byte("2"))
	BytesToUInt64([]byte("2"))
	BytesToInt64([]byte("2"))
	BytesToFloat32([]byte("2"))
	BytesToFloat64([]byte("2"))
	BytesToInt32([]byte("2"))
}
