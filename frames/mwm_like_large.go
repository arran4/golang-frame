package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mwm_like_large.png
var MwmLikeLargePng []byte

var MwmLikeLarge = &Definition{
	Name: "mwm_like_large",
	Middle: image.Rect(24, 30, 72, 72),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MwmLikeLargePng))
	if err != nil { panic(err) }
	MwmLikeLarge.Image = img
}
