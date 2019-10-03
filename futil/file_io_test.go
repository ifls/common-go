package futil

import (
	"fmt"
	"gocore/util"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestReadFile(t *testing.T) {
	filepath := "/Users/ifls/test.jpg"
	data, err := ReadFile(filepath)
	util.DevInfo("%v", data)
	util.LogErr(err)
	WriteFile(data, "./test.jpg")
}

func TestUpload(t *testing.T) {
	filepath := "/Users/ifls/tt.jpg"
	uploadGcpOssName(filepath, TEST_BUCKET)
}

func TestDownload(t *testing.T) {
	//filepath := "./huoyi.jpg"
	//read(filepath, "dev_bucket-ifls", "jpg/2019/08/07/zJCBDQWEf9QIqI9FGZ87rA==.jpg")
}

func TestIsDir(t *testing.T) {
	assert.Equal(t, IsDir("/Users/ifls"), false)
}

func TestWriteFile(t *testing.T) {
	WriteFile([]byte("adfwe"), "../../test/temp_file/temp.txt")
	allline, err := Readlines_buf("../../test/temp_file/temp.txt")
	if err != nil {
		t.Fatal("error occur")
	}

	util.Judge("adfwe" == allline[0], t)
}

func TestFileIOExist(t *testing.T) {
	assert.Equal(t, Exist("./file_io.go"), true)
}

func TestReadlines(t *testing.T) {
	lines, err := Readlines("/Users/ifls/.bash_profile")
	lines2, err2 := Readlines_buf("/Users/ifls/.bash_profile")
	fmt.Println(len(lines), "&&", len(lines2))
	util.Judge(len(lines) == len(lines2) && err == nil && err2 == nil, t)
}
