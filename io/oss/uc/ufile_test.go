package uc

import (
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"github.com/ufilesdk-dev/ufile-gosdk/example/helper"
	"log"
	"os"
	"testing"
)

func TestUFile(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	if _, err := os.Stat(helper.FakeSmallFilePath); os.IsNotExist(err) {
		helper.GenerateFakefile(helper.FakeSmallFilePath, helper.FakeSmallFileSize)
	}
	if _, err := os.Stat(helper.FakeBigFilePath); os.IsNotExist(err) {
		helper.GenerateFakefile(helper.FakeBigFilePath, helper.FakeBigFileSize)
	}
	config, err := ufsdk.LoadConfig(helper.ConfigFile)
	if err != nil {
		panic(err.Error())
	}
	req, err := ufsdk.NewFileRequest(config, nil)
	if err != nil {
		panic(err.Error())
	}

	//可以替换为自定义源bucketName
	srcBucketName = config.BucketName

	var fileKey string
	fileKey = helper.GenerateUniqKey()
	scheduleUploadhelper(helper.FakeSmallFilePath, fileKey, putUpload, req)

	fileKey = helper.GenerateUniqKey()
	scheduleUploadhelper(helper.FakeSmallFilePath, fileKey, postUpload, req)

	fileKey = helper.GenerateUniqKey()
	scheduleUploadhelper(helper.FakeBigFilePath, fileKey, mput, req)
	fileKey = helper.GenerateUniqKey()
	scheduleUploadhelper(helper.FakeBigFilePath, fileKey, asyncmput, req)

}
