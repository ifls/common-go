package futil

import (
	"github.com/ifls/gocore/util"
	"image/jpeg"
	"os"
	"testing"
)

func TestImage(t *testing.T) {
	filepath := "./tt.jpg"
	fd, err := os.Open(filepath)
	if err != nil {
		util.LogErr(err)
		return
	}
	defer fd.Close()

	img, err := jpeg.Decode(fd)
	if err != nil {
		util.LogErr(err)
		return
	}

	util.DevInfo("%+v\n", img.Bounds())
}
