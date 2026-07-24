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
	Middle: image.Rect(78, 87, 180, 159),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MwmLikeXlargePng))
	if err != nil { panic(err) }
	MwmLikeXlarge.Image = img
}
