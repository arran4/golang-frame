package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_house_xlarge.png
var WindowHouseXlargePng []byte

var WindowHouseXlarge = &Definition{
	Name: "window_house_xlarge",
	Middle: image.Rect(48, 48, 240, 207),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowHouseXlargePng))
	if err != nil { panic(err) }
	WindowHouseXlarge.Image = img
}
