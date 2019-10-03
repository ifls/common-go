package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
)

func Md5Hash(src []byte) []byte {
	data := md5.Sum(src)
	rlt := make([]byte, 0)
	for _, byte := range data {
		rlt = append(rlt, byte)
	}
	return rlt
}

func Sha1Hash(src []byte) []byte {
	data := sha1.Sum(src)
	rlt := make([]byte, 0)
	for _, byte := range data {
		rlt = append(rlt, byte)
	}
	return rlt
}

func Base64Encoding(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func Base32Encoding(data []byte) string {
	return base32.StdEncoding.EncodeToString(data)
}

func Sha256Hash(src []byte) []byte {
	data := sha256.Sum256(src)
	rlt := make([]byte, 0)
	for _, byte := range data {
		rlt = append(rlt, byte)
	}
	return rlt
}
