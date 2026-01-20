package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed macosx_like_xlarge.png
var MacosxLikeXlargePng []byte

var MacosxLikeXlarge = &Definition{
	Name: "macosx_like_xlarge",
	Middle: image.Rect(162, 72, 180, 180),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacosxLikeXlargePng))
	if err != nil { panic(err) }
	MacosxLikeXlarge.Image = img
}
