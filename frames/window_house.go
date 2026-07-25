package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_house.png
var WindowHousePng []byte

var WindowHouse = &Definition{
	Name: "window_house",
	Middle: image.Rect(16, 16, 80, 69),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowHousePng))
	if err != nil { panic(err) }
	WindowHouse.Image = img
}
