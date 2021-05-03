package file

import (
	"bytes"
	"github.com/ifls/gocore/utils"
	"image"
	"image/jpeg"
)

func GetImageFileInfo(data []byte) (point image.Point) {
	var bs bytes.Buffer

	bs.Write(data)

	img, err := jpeg.Decode(&bs)
	if err != nil {
		utils.LogErr(err)
		return
	}

	utils.DevInfo("%+v\n", img.Bounds())
	return img.Bounds().Max
}
