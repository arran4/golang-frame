package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro_close_only_large.png
var WindowRetroCloseOnlyLargePng []byte

var WindowRetroCloseOnlyLarge = &Definition{
	Name: "window_retro_close_only_large",
	Middle: image.Rect(44, 50, 144, 180),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroCloseOnlyLargePng))
	if err != nil { panic(err) }
	WindowRetroCloseOnlyLarge.Image = img
}
