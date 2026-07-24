package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed rainbow_large.png
var RainbowLargePng []byte

var RainbowLarge = &Definition{
	Name: "rainbow_large",
	Middle: image.Rect(56, 56, 136, 136),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(RainbowLargePng))
	if err != nil { panic(err) }
	RainbowLarge.Image = img
}
