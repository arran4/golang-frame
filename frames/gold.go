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
	Middle: image.Rect(12, 12, 52, 52),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(GoldPng))
	if err != nil { panic(err) }
	Gold.Image = img
}
