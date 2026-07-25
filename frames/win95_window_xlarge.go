package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win95_window_xlarge.png
var Win95WindowXlargePng []byte

var Win95WindowXlarge = &Definition{
	Name: "win95_window_xlarge",
	Middle: image.Rect(66, 75, 114, 270),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win95WindowXlargePng))
	if err != nil { panic(err) }
	Win95WindowXlarge.Image = img
}
