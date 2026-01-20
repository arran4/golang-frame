package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_construction.png
var SignConstructionPng []byte

var SignConstruction = &Definition{
	Name: "sign_construction",
	Middle: image.Rect(16, 16, 80, 80),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignConstructionPng))
	if err != nil { panic(err) }
	SignConstruction.Image = img
}
