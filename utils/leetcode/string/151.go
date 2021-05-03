package string

import (
	"bytes"
	"log"
)

//

func reverseWords(s string) string {
	// 切分
	bytesSlice := bytes.Split([]byte(s), []byte{' '})
	log.Println(bytesSlice)
	var dstBytes []byte
	for i := len(bytesSlice) - 1; i >= 0; i-- {
		// 剔除[]byte{''}的元素
		if len(bytesSlice[i]) == 0 {
			continue
		}
		if len(dstBytes) > 0 {
			dstBytes = append(dstBytes, byte(' '))
		}
		dstBytes = append(dstBytes, bytesSlice[i]...)

	}

	return string(dstBytes)
}
