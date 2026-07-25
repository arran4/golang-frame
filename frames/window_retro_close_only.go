package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro_close_only.png
var WindowRetroCloseOnlyPng []byte

var WindowRetroCloseOnly = &Definition{
	Name: "window_retro_close_only",
	Middle: image.Rect(22, 25, 72, 90),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroCloseOnlyPng))
	if err != nil { panic(err) }
	WindowRetroCloseOnly.Image = img
}
