package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed dots_large.png
var DotsLargePng []byte

var DotsLarge = &Definition{
	Name: "dots_large",
	Middle: image.Rect(32, 32, 64, 64),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(DotsLargePng))
	if err != nil { panic(err) }
	DotsLarge.Image = img
}
