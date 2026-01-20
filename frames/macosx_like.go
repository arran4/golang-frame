package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed macosx_like.png
var MacosxLikePng []byte

var MacosxLike = &Definition{
	Name: "macosx_like",
	Middle: image.Rect(8, 16, 40, 40),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacosxLikePng))
	if err != nil { panic(err) }
	MacosxLike.Image = img
}
