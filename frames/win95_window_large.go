package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win95_window_large.png
var Win95WindowLargePng []byte

var Win95WindowLarge = &Definition{
	Name: "win95_window_large",
	Middle: image.Rect(44, 50, 76, 180),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win95WindowLargePng))
	if err != nil { panic(err) }
	Win95WindowLarge.Image = img
}
