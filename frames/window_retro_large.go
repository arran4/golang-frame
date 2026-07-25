package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro_large.png
var WindowRetroLargePng []byte

var WindowRetroLarge = &Definition{
	Name: "window_retro_large",
	Middle: image.Rect(44, 50, 76, 180),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroLargePng))
	if err != nil { panic(err) }
	WindowRetroLarge.Image = img
}
