package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed floral.png
var FloralPng []byte

var Floral = &Definition{
	Name: "floral",
	Middle: image.Rect(32, 32, 64, 64),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(FloralPng))
	if err != nil { panic(err) }
	Floral.Image = img
}
