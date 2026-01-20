package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_glass_large.png
var WindowGlassLargePng []byte

var WindowGlassLarge = &Definition{
	Name: "window_glass_large",
	Middle: image.Rect(16, 36, 112, 112),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowGlassLargePng))
	if err != nil { panic(err) }
	WindowGlassLarge.Image = img
}
