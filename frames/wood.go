package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed wood.png
var WoodPng []byte

var Wood = &Definition{
	Name: "wood",
	Middle: image.Rect(16, 16, 80, 80),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WoodPng))
	if err != nil { panic(err) }
	Wood.Image = img
}
