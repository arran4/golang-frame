package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed metal.png
var MetalPng []byte

var Metal = &Definition{
	Name: "metal",
	Middle: image.Rect(8, 8, 40, 40),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MetalPng))
	if err != nil { panic(err) }
	Metal.Image = img
}
