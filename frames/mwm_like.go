package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mwm_like.png
var MwmLikePng []byte

var MwmLike = &Definition{
	Name: "mwm_like",
	Middle: image.Rect(12, 18, 36, 36),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MwmLikePng))
	if err != nil { panic(err) }
	MwmLike.Image = img
}
