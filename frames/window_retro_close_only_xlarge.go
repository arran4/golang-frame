package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro_close_only_xlarge.png
var WindowRetroCloseOnlyXlargePng []byte

var WindowRetroCloseOnlyXlarge = &Definition{
	Name: "window_retro_close_only_xlarge",
	Middle: image.Rect(66, 75, 216, 270),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroCloseOnlyXlargePng))
	if err != nil { panic(err) }
	WindowRetroCloseOnlyXlarge.Image = img
}
