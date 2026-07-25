package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_house_large.png
var WindowHouseLargePng []byte

var WindowHouseLarge = &Definition{
	Name: "window_house_large",
	Middle: image.Rect(32, 32, 160, 138),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowHouseLargePng))
	if err != nil { panic(err) }
	WindowHouseLarge.Image = img
}
