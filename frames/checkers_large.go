package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed checkers_large.png
var CheckersLargePng []byte

var CheckersLarge = &Definition{
	Name: "checkers_large",
	Middle: image.Rect(32, 32, 64, 64),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(CheckersLargePng))
	if err != nil { panic(err) }
	CheckersLarge.Image = img
}
