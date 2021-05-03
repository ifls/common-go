package file

import (
	"fmt"
	"github.com/ifls/gocore/utils"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestReadFile(t *testing.T) {
	filepath := "/Users/ifls/test.jpg"
	data, err := ReadFile(filepath)
	utils.DevInfo("%v", data)
	utils.LogErr(err)
	t.Error(WriteFile(data, "./test.jpg"))
}

func TestUpload(t *testing.T) {
	filepath := "/Users/ifls/tt.jpg"
	t.Error(uploadGcpOssName(filepath, TestBucket))
}

func TestDownload(t *testing.T) {
	//filepath := "./huoyi.jpg"
	//read(filepath, "dev_bucket-ifls", "jpg/2019/08/07/zJCBDQWEf9QIqI9FGZ87rA==.jpg")
}

func TestIsDir(t *testing.T) {
	assert.Equal(t, true, IsDir("/Users/ifls"))
	assert.Equal(t, true, IsDir("."))
	assert.Equal(t, true, IsDir(".."))
	//assert.Equal(t, false, IsDir("./file_"))
	assert.Equal(t, false, IsDir("./file_io.go"))
	//assert.Equal(t, false, IsDir("./file_io2.go"))
}

func TestWriteFile(t *testing.T) {
	t.Error(WriteFile([]byte("adfwe"), "../../test/temp_file/temp.txt"))
	allline, err := ReadlinesBuf("../../test/temp_file/temp.txt")
	if err != nil {
		t.Fatal("error occur")
	}

	utils.Judge("adfwe" == allline[0], t)
}

func TestFileIOExist(t *testing.T) {
	assert.Equal(t, true, Exist("."))
	assert.Equal(t, false, Exist("./cc"))

	assert.Equal(t, true, Exist("./file_io.go"))
	assert.Equal(t, false, Exist("./file_io2.go"))
}

func TestReadlines(t *testing.T) {
	lines, err := Readlines("/Users/ifls/.bash_profile")
	lines2, err2 := ReadlinesBuf("/Users/ifls/.bash_profile")
	fmt.Println(len(lines), "&&", len(lines2))
	utils.Judge(len(lines) == len(lines2) && err == nil && err2 == nil, t)
}

func TestIsFile(t *testing.T) {
	IsFile("")
	_ = CreateDir("")
	_ = DeleteDir("", true)
	_ = WriteUrl(nil, "")
	_, _ = ReadUrl("")
	_ = downloadGcpOss("", "", "")

	_ = uploadGcpOss("", "", "")
}
