package utils

import (
	"encoding/base64"
	"github.com/ifls/gocore/utils/log"
	"testing"
)

func TestMd5Hash(t *testing.T) {
	rlt := Md5Hash([]byte("fwefwfwfwfwefwefwfwfwefefgebggiowoefjwoefhweofowfoewnflwiengfowgjwpoeffwfwfwfewef"))
	//rlt2 := string(rlt)
	log.DevInfo("rlt = %v", rlt)
	log.DevInfo("length = %d", len(rlt))
}

func TestShaHash(t *testing.T) {
	rlt := Sha1Hash([]byte("fwefwfwfwfwefwefwfwfwefefgebggiowoefjwoefhweofowfoewnflwiengfowgjwpoeffwfwfwfewef"))
	//rlt2 := string(rlt)
	log.DevInfo("rlt = %v", rlt)
	log.DevInfo("length = %d", len(rlt))
}

func TestBase64Encoding(t *testing.T) {
	str := Base64Encoding([]byte(";.3%^&*(((*^%$##"))
	log.DevInfo("%v %s\n", str, str)
	str2, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		log.LogErr(err)
	}
	log.DevInfo("%v %s\n", str2, str2)
}
