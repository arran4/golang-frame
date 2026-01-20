package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro_xlarge.png
var WindowRetroXlargePng []byte

var WindowRetroXlarge = &Definition{
	Name: "window_retro_xlarge",
	Middle: image.Rect(24, 60, 168, 168),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroXlargePng))
	if err != nil { panic(err) }
	WindowRetroXlarge.Image = img
}
