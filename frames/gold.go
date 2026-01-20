package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed gold.png
var GoldPng []byte

var Gold = &Definition{
	Name: "gold",
	Middle: image.Rect(16, 16, 80, 80),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(GoldPng))
	if err != nil { panic(err) }
	Gold.Image = img
}
