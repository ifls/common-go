package file

import (
	"github.com/ifls/gocore/utils/log"
	"image/jpeg"
	"os"
	"testing"
)

func TestImage(t *testing.T) {
	filepath := "./tt.jpg"
	fd, err := os.Open(filepath)
	if err != nil {
		log.LogErr(err)
		return
	}
	defer func() {
		_ = fd.Close()
	}()

	img, err := jpeg.Decode(fd)
	if err != nil {
		log.LogErr(err)
		return
	}

	log.DevInfo("%+v\n", img.Bounds())
}

func TestGetImageFileInfo(t *testing.T) {
	GetImageFileInfo(nil)
}
