package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win95_window.png
var Win95WindowPng []byte

var Win95Window = &Definition{
	Name: "win95_window",
	Middle: image.Rect(22, 25, 38, 90),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win95WindowPng))
	if err != nil { panic(err) }
	Win95Window.Image = img
}
