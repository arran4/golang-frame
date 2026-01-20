package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed rainbow_xlarge.png
var RainbowXlargePng []byte

var RainbowXlarge = &Definition{
	Name: "rainbow_xlarge",
	Middle: image.Rect(24, 48, 168, 168),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(RainbowXlargePng))
	if err != nil { panic(err) }
	RainbowXlarge.Image = img
}
