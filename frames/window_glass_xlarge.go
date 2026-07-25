package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_glass_xlarge.png
var WindowGlassXlargePng []byte

var WindowGlassXlarge = &Definition{
	Name: "window_glass_xlarge",
	Middle: image.Rect(18, 72, 270, 270),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowGlassXlargePng))
	if err != nil { panic(err) }
	WindowGlassXlarge.Image = img
}
