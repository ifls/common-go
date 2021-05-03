package net

import (
	"github.com/ifls/gocore/utils"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"strings"
)

func CopyFile(w io.Reader, url string) {
	filname := strings.ReplaceAll(url, "/", "_")
	path := "/Users/ifls/Downloads/logs/imgs/" + filname
	f, err := os.Create(path)
	if err != nil {
		utils.LogErr(err, zap.String("reason", "file create error"))
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	if _, err = io.Copy(f, w); err != nil {
		utils.LogErr(err, zap.String("reason", "data copy error"))
	}
	if err := f.Close(); err != nil {
		utils.LogErr(err, zap.String("reason", "file close error"))
	}
}
