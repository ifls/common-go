package futil

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/ifls/gocore/util"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
)

const (
	TEST_BUCKET = "dev_bucket-ifls"
	GCP_OSS_URL = "https://storage.cloud.google.com/"
)

var client *storage.Client

func init() {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}
	var err error
	ctx := context.Background()
	client, err = storage.NewClient(ctx)
	if err != nil {
		util.LogErr(err)
		return
	}
}

func createBucket() {

}

func listBucket() {

}

func getBucketInfo() {

}

type OnUploadSucc func(string)

func WriteGcpOss(data []byte, bucket string, object string, cb OnUploadSucc) error {
	buf := bytes.Buffer{}
	buf.Write(data)
	return WriteGcpOssFromReader(&buf, bucket, object, cb)
}

func WriteGcpOssFromReader(reader io.Reader, bucket string, object string, cb OnUploadSucc) error {
	util.LogInfo("write to gcp oss", zap.String("bucket", bucket), zap.String("objectName", object))
	ctx := context.Background()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err := io.Copy(wc, reader); err != nil {
		util.LogErr(err, zap.String("reason", "copy to gcp oss error"))
		return err
	}
	if err := wc.Close(); err != nil {
		util.LogErr(err, zap.String("reason", "gcp oss close error"))
		return err
	}

	if cb != nil {
		util.LogInfo("write to gcp oss cb", zap.String("bucket", bucket), zap.String("objectName", object))
		cb(GCP_OSS_URL + bucket + "/" + object)
	}

	return nil
}

func ReadGcpOss(bucket string, object string) ([]byte, error) {
	ctx := context.Background()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		util.LogErr(err, zap.String("reason", "read from gcp oss error"))
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		util.LogErr(err, zap.String("reason", "ioutil.ReadAll() from gcp oss error"))
		return nil, err
	}
	return data, nil
}

func listObject() {

}

func getObjectInfo() {

}

func deleteObject() {

}

func renameObject() {

}

func moveObject() {

}

func lockObject() {

}

func public() {

}

func GetDir(fileType string, filename string) string {
	//now := time.Now()
	var buf bytes.Buffer
	//yearStr := now.Year()
	//monthStr := int(now.Month())
	//dayStr := now.Day()

	//fmt.Fprintf(&buf, "%s/%d/%02d/%02d/", fileType, yearStr, monthStr, dayStr)
	fmt.Fprintf(&buf, "%s/%s/%0s/", filename[0:1], filename[1:3], filename[3:6])
	return buf.String()
}
