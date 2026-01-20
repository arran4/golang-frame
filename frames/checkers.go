package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed checkers.png
var CheckersPng []byte

var Checkers = &Definition{
	Name: "checkers",
	Middle: image.Rect(16, 16, 32, 32),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(CheckersPng))
	if err != nil { panic(err) }
	Checkers.Image = img
}
