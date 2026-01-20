package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed checkers_xlarge.png
var CheckersXlargePng []byte

var CheckersXlarge = &Definition{
	Name: "checkers_xlarge",
	Middle: image.Rect(48, 48, 96, 96),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(CheckersXlargePng))
	if err != nil { panic(err) }
	CheckersXlarge.Image = img
}
