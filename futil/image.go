package futil

import (
	"bytes"
	"github.com/ifls/gocore/util"
	"image"
	"image/jpeg"
)

func GetImageFileInfo(data []byte) (point image.Point) {
	var bs bytes.Buffer

	bs.Write(data)

	img, err := jpeg.Decode(&bs)
	if err != nil {
		util.LogErr(err)
		return
	}

	util.DevInfo("%+v\n", img.Bounds())
	return img.Bounds().Max
}
