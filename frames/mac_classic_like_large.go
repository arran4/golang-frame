package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_classic_like_large.png
var MacClassicLikeLargePng []byte

var MacClassicLikeLarge = &Definition{
	Name: "mac_classic_like_large",
	Middle: image.Rect(8, 8, 56, 56),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacClassicLikeLargePng))
	if err != nil { panic(err) }
	MacClassicLikeLarge.Image = img
}
