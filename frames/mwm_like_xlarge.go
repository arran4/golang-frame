package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mwm_like_xlarge.png
var MwmLikeXlargePng []byte

var MwmLikeXlarge = &Definition{
	Name: "mwm_like_xlarge",
	Middle: image.Rect(36, 48, 108, 108),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MwmLikeXlargePng))
	if err != nil { panic(err) }
	MwmLikeXlarge.Image = img
}
