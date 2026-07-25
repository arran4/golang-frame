package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro_min_max_close.png
var WindowRetroMinMaxClosePng []byte

var WindowRetroMinMaxClose = &Definition{
	Name: "window_retro_min_max_close",
	Middle: image.Rect(22, 25, 38, 90),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroMinMaxClosePng))
	if err != nil { panic(err) }
	WindowRetroMinMaxClose.Image = img
}
