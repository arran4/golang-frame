package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed dots.png
var DotsPng []byte

var Dots = &Definition{
	Name: "dots",
	Middle: image.Rect(16, 16, 32, 32),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(DotsPng))
	if err != nil { panic(err) }
	Dots.Image = img
}
