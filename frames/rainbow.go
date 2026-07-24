package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed rainbow.png
var RainbowPng []byte

var Rainbow = &Definition{
	Name: "rainbow",
	Middle: image.Rect(28, 28, 68, 68),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(RainbowPng))
	if err != nil { panic(err) }
	Rainbow.Image = img
}
