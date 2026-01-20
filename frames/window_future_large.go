package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_future_large.png
var WindowFutureLargePng []byte

var WindowFutureLarge = &Definition{
	Name: "window_future_large",
	Middle: image.Rect(24, 48, 104, 104),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowFutureLargePng))
	if err != nil { panic(err) }
	WindowFutureLarge.Image = img
}
