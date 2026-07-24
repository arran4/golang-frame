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
	Middle: image.Rect(33, 87, 111, 111),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MwmLikeXlargePng))
	if err != nil { panic(err) }
	MwmLikeXlarge.Image = img
}
