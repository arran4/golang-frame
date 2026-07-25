package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro_min_max_close_xlarge.png
var WindowRetroMinMaxCloseXlargePng []byte

var WindowRetroMinMaxCloseXlarge = &Definition{
	Name: "window_retro_min_max_close_xlarge",
	Middle: image.Rect(66, 75, 114, 270),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroMinMaxCloseXlargePng))
	if err != nil { panic(err) }
	WindowRetroMinMaxCloseXlarge.Image = img
}
