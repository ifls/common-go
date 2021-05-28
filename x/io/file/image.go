package file

import (
	"bytes"
	"github.com/ifls/gocore/utils/log"
	"image"
	"image/jpeg"
)

func GetImageFileInfo(data []byte) (point image.Point) {
	var bs bytes.Buffer

	bs.Write(data)

	img, err := jpeg.Decode(&bs)
	if err != nil {
		log.LogErr(err)
		return
	}

	log.DevInfo("%+v\n", img.Bounds())
	return img.Bounds().Max
}
