package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed macosx_like_large.png
var MacosxLikeLargePng []byte

var MacosxLikeLarge = &Definition{
	Name: "macosx_like_large",
	Middle: image.Rect(108, 48, 120, 120),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacosxLikeLargePng))
	if err != nil { panic(err) }
	MacosxLikeLarge.Image = img
}
