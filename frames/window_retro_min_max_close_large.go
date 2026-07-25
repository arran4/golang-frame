package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro_min_max_close_large.png
var WindowRetroMinMaxCloseLargePng []byte

var WindowRetroMinMaxCloseLarge = &Definition{
	Name: "window_retro_min_max_close_large",
	Middle: image.Rect(44, 50, 76, 180),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroMinMaxCloseLargePng))
	if err != nil { panic(err) }
	WindowRetroMinMaxCloseLarge.Image = img
}
