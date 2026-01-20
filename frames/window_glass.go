package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_glass.png
var WindowGlassPng []byte

var WindowGlass = &Definition{
	Name: "window_glass",
	Middle: image.Rect(8, 18, 56, 56),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowGlassPng))
	if err != nil { panic(err) }
	WindowGlass.Image = img
}
