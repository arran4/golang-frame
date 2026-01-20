package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_future_xlarge.png
var WindowFutureXlargePng []byte

var WindowFutureXlarge = &Definition{
	Name: "window_future_xlarge",
	Middle: image.Rect(36, 72, 252, 252),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowFutureXlargePng))
	if err != nil { panic(err) }
	WindowFutureXlarge.Image = img
}
