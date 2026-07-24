package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_system_3_like_xlarge.png
var MacSystem3LikeXlargePng []byte

var MacSystem3LikeXlarge = &Definition{
	Name: "mac_system_3_like_xlarge",
	Middle: image.Rect(6, 66, 144, 144),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacSystem3LikeXlargePng))
	if err != nil { panic(err) }
	MacSystem3LikeXlarge.Image = img
}
