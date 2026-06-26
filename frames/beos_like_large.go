package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed beos_like_large.png
var BeosLikeLargePng []byte

var BeosLikeLarge = &Definition{
	Name: "beos_like_large",
	Middle: image.Rect(8, 8, 56, 56),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(BeosLikeLargePng))
	if err != nil { panic(err) }
	BeosLikeLarge.Image = img
}
