package net

import (
	"bytes"
	"encoding/binary"
)

//struct => []byte
//func Encode(srt interface{}) ([]byte, error) {
//	switch tp := srt.(type) {
//	case *pb.User:
//		if srt, ok := srt.(*pb.User); !ok {
//			return nil, errors.New("type convert err")
//		} else {
//			return proto.Marshal(srt)
//		}
//	default:
//		_ = tp
//	}
//	return nil, errors.New("not find type")
//}
//
////bytes => struct
//func Decode(data []byte, typeName string) (interface{}, error) {
//	switch typeName {
//	case "User":
//		template := &pb.User{}
//		return template, proto.Unmarshal(data, template)
//	}
//
//	return nil, errors.New("no match")
//}

func ToBytes(x interface{}) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func UInt32ToBytes(x uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func BytesToUInt32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func BytesToUInt64(b []byte) uint64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint64
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func BytesToFloat32(b []byte) float32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x float32
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func BytesToFloat64(b []byte) float64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x float64
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
