package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_classic_like_xlarge.png
var MacClassicLikeXlargePng []byte

var MacClassicLikeXlarge = &Definition{
	Name: "mac_classic_like_xlarge",
	Middle: image.Rect(60, 66, 132, 132),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacClassicLikeXlargePng))
	if err != nil { panic(err) }
	MacClassicLikeXlarge.Image = img
}
