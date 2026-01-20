package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_future.png
var WindowFuturePng []byte

var WindowFuture = &Definition{
	Name: "window_future",
	Middle: image.Rect(12, 24, 84, 84),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowFuturePng))
	if err != nil { panic(err) }
	WindowFuture.Image = img
}
